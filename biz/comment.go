package biz

import (
	"context"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"
	"moredoc/util/captcha"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

type CommentAPIService struct {
	pb.UnimplementedCommentAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewCommentAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *CommentAPIService) {
	return &CommentAPIService{dbModel: dbModel, logger: logger.Named("CommentAPIService")}
}

func (s *CommentAPIService) checkLogin(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCLogin(s.dbModel, ctx)
}

func (s *CommentAPIService) checkPermission(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCPermission(s.dbModel, ctx)
}

// 发表评论
func (s *CommentAPIService) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*emptypb.Empty, error) {
	userClaims, err := checkGRPCLogin(s.dbModel, ctx)
	if err != nil {
		return nil, err
	}

	// 评论验证码错误
	cfg := s.dbModel.GetConfigOfSecurity(model.ConfigSecurityEnableCaptchaComment)
	if cfg.EnableCaptchaComment {
		if req.CaptchaId == "" || req.Captcha == "" {
			return nil, status.Errorf(codes.InvalidArgument, "请输入验证码")
		}
		if !captcha.VerifyCaptcha(req.CaptchaId, req.Captcha, true) {
			return nil, status.Errorf(codes.InvalidArgument, "验证码错误")
		}
	}

	if _, err = s.dbModel.CanIAccessComment(userClaims.UserId); err != nil {
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}

	comment := &model.Comment{}
	err = util.CopyStruct(req, comment)
	if err != nil {
		s.logger.Error("CreateDocument", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "发布评论失败："+err.Error())
	}

	if comment.DocumentId <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "文档id不能为空")
	}

	if comment.Content == "" {
		return nil, status.Errorf(codes.InvalidArgument, "评论内容不能为空")
	}

	comment.IP = util.GetGRPCRemoteIP(ctx)
	comment.Status = int8(s.dbModel.GetDefaultCommentStatus(userClaims.UserId))
	comment.UserId = userClaims.UserId
	if req.Type == 1 {
		err = s.dbModel.CreateArticleComment(comment)
	} else {
		err = s.dbModel.CreateDocumentComment(comment)
	}
	if err != nil {
		s.logger.Error("CreateComment", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "发布评论失败："+err.Error())
	}

	return &emptypb.Empty{}, nil
}

// 更新评论，仅限管理员
func (s *CommentAPIService) UpdateComment(ctx context.Context, req *pb.Comment) (*emptypb.Empty, error) {
	userClaims, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("UpdateComment", zap.Any("user", userClaims), zap.Any("req", req))

	// 只允许更新评论状态和内容
	updateFields := []string{"content", "status"}
	comment := &model.Comment{}
	util.CopyStruct(req, comment)
	err = s.dbModel.UpdateComment(comment, updateFields...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "更新评论失败")
	}
	return &emptypb.Empty{}, nil
}

// 登录的用户，可删除自己的评论
// 管理员，可删除任意评论
func (s *CommentAPIService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*emptypb.Empty, error) {
	userClaims, err := s.checkPermission(ctx)
	if userClaims == nil {
		// 未登录用户
		return nil, err
	}

	var userIds []int64
	if !userClaims.HaveAccess { // 非管理员，限定只能删除自己的评论
		userIds = append(userIds, userClaims.UserId)
	}

	err = s.dbModel.DeleteComment(req.Id, userIds...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *CommentAPIService) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.Comment, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}
	s.logger.Debug("GetComment", zap.Any("req", req))
	comment, err := s.dbModel.GetComment(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取评论失败:"+err.Error())
	}
	user, _ := s.dbModel.GetUser(comment.UserId, model.UserPublicFields...)

	pbComment := &pb.Comment{}
	util.CopyStruct(comment, pbComment)
	util.CopyStruct(user, pbComment.User)

	return pbComment, nil
}

func (s *CommentAPIService) ListComment(ctx context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	userClaims, _ := s.checkPermission(ctx)
	isLogin := userClaims != nil
	haveAccess := userClaims != nil && userClaims.HaveAccess
	opt := &model.OptionGetCommentList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
		QueryIn:   make(map[string][]interface{}),
		QueryLike: make(map[string][]interface{}),
		Sort:      []string{req.Order},
	}

	// default status
	opt.QueryIn["status"] = []interface{}{model.CommentStatusApproved}

	if req.DocumentId > 0 {
		opt.QueryIn["document_id"] = []interface{}{req.DocumentId}
		opt.Page = 1
		opt.Size = 1000000
	}

	if len(req.Type) > 0 {
		opt.QueryIn["type"] = util.Slice2Interface(req.Type)
	}

	if len(req.ParentId) > 0 {
		opt.QueryIn["parent_id"] = util.Slice2Interface(req.ParentId)
	}

	if haveAccess && req.Wd != "" {
		opt.QueryLike["content"] = []interface{}{req.Wd}
	}

	if (isLogin && req.UserId == userClaims.UserId) || haveAccess {
		delete(opt.QueryIn, "status")
		if len(req.Status) > 0 {
			opt.QueryIn["status"] = util.Slice2Interface(req.Status)
		}
	}

	if req.UserId > 0 {
		opt.QueryIn["user_id"] = []interface{}{req.UserId}
	}

	comments, total, err := s.dbModel.GetCommentList(opt)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, status.Errorf(codes.Internal, "获取评论列表失败")
	}

	resp := &pb.ListCommentReply{
		Total: total,
	}
	util.CopyStruct(comments, &resp.Comment)

	var (
		userIds                 []int64
		documentIds             []int64
		articleIds              []int64
		userIdMapCommentIdx     = make(map[int64][]int)
		documentIdMapCommentIdx = make(map[int64][]int)
		articleIdMapCommentIdx  = make(map[int64][]int)
	)
	for idx, comment := range comments {
		userIds = append(userIds, comment.UserId)
		if comment.Type == 1 {
			articleIds = append(articleIds, comment.DocumentId)
			articleIdMapCommentIdx[comment.DocumentId] = append(articleIdMapCommentIdx[comment.DocumentId], idx)
		} else {
			documentIds = append(documentIds, comment.DocumentId)
			documentIdMapCommentIdx[comment.DocumentId] = append(documentIdMapCommentIdx[comment.DocumentId], idx)
		}
		userIdMapCommentIdx[comment.UserId] = append(userIdMapCommentIdx[comment.UserId], idx)
	}

	if len(userIds) > 0 {
		users, _, _ := s.dbModel.GetUserList(&model.OptionGetUserList{
			SelectFields: model.UserPublicFields,
			WithCount:    false,
			QueryIn:      map[string][]interface{}{"id": util.Slice2Interface(userIds)},
		})

		for _, user := range users {
			indexes := userIdMapCommentIdx[user.Id]
			for _, idx := range indexes {
				util.CopyStruct(user, &resp.Comment[idx].User)
			}
		}
	}

	if req.WithDocumentTitle {
		if len(documentIds) > 0 {
			documents, _, _ := s.dbModel.GetDocumentList(&model.OptionGetDocumentList{
				SelectFields: []string{"id", "title", "uuid"},
				WithCount:    false,
				QueryIn:      map[string][]interface{}{"id": util.Slice2Interface(documentIds)},
			})

			for _, document := range documents {
				indexes := documentIdMapCommentIdx[document.Id]
				for _, idx := range indexes {
					resp.Comment[idx].DocumentTitle = document.Title
					resp.Comment[idx].DocumentUuid = document.UUID
				}
			}
		}

		if len(articleIds) > 0 {
			articles, _, _ := s.dbModel.GetArticleList(&model.OptionGetArticleList{
				SelectFields: []string{"id", "title", "identifier"},
				WithCount:    false,
				QueryIn:      map[string][]interface{}{"id": util.Slice2Interface(articleIds)},
			})

			for _, article := range articles {
				indexes := articleIdMapCommentIdx[article.Id]
				for _, idx := range indexes {
					resp.Comment[idx].DocumentTitle = article.Title
					resp.Comment[idx].DocumentUuid = article.Identifier
				}
			}
		}

	}

	return resp, nil
}

func (s *CommentAPIService) CheckComment(ctx context.Context, req *pb.CheckCommentRequest) (*emptypb.Empty, error) {
	_, err := s.checkPermission(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.UpdateCommentStatus(req.Id, req.Status)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

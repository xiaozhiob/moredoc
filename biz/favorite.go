package biz

import (
	"context"

	pb "moredoc/api/v1"
	"moredoc/middleware/auth"
	"moredoc/model"
	"moredoc/util"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FavoriteAPIService struct {
	pb.UnimplementedFavoriteAPIServer
	dbModel *model.DBModel
	logger  *zap.Logger
}

func NewFavoriteAPIService(dbModel *model.DBModel, logger *zap.Logger) (service *FavoriteAPIService) {
	return &FavoriteAPIService{dbModel: dbModel, logger: logger.Named("FavoriteAPIService")}
}

func (s *FavoriteAPIService) checkLogin(ctx context.Context) (*auth.UserClaims, error) {
	return checkGRPCLogin(s.dbModel, ctx)
}

func (s *FavoriteAPIService) CreateFavorite(ctx context.Context, req *pb.Favorite) (*pb.Favorite, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	yes, _ := s.dbModel.CanIAccessFavorite(userClaims.UserId)
	if !yes {
		return nil, status.Errorf(codes.PermissionDenied, "您已经被禁止使用收藏功能")
	}

	favorite := &model.Favorite{
		UserId:     userClaims.UserId,
		DocumentId: req.DocumentId,
		Type:       req.Type,
		IP:         util.GetGRPCRemoteIP(ctx),
	}

	exsit, _ := s.dbModel.GetUserFavorite(favorite.UserId, favorite.DocumentId, req.Type)
	if exsit.Id > 0 {
		return nil, status.Errorf(codes.AlreadyExists, "您已经收藏过了")
	}
	if req.Type == 1 {
		err = s.dbModel.CreateArticleFavorite(favorite)
	} else {
		err = s.dbModel.CreateDocumentFavorite(favorite)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "收藏失败:"+err.Error())
	}

	pbFavorite := &pb.Favorite{}
	util.CopyStruct(favorite, pbFavorite)

	return pbFavorite, nil
}

func (s *FavoriteAPIService) DeleteFavorite(ctx context.Context, req *pb.DeleteFavoriteRequest) (*emptypb.Empty, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	err = s.dbModel.DeleteFavorite(userClaims.UserId, req.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "删除失败:"+err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (s *FavoriteAPIService) GetFavorite(ctx context.Context, req *pb.GetFavoriteRequest) (*pb.Favorite, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	favorite, _ := s.dbModel.GetUserFavorite(userClaims.UserId, req.DocumentId, req.Type)

	pbFavorite := &pb.Favorite{}
	if favorite.Id > 0 {
		util.CopyStruct(&favorite, pbFavorite)
	}

	return pbFavorite, nil
}

// 获取用户自身的收藏列表(这里指的是文档的收藏)
func (s *FavoriteAPIService) ListFavorite(ctx context.Context, req *pb.ListFavoriteRequest) (*pb.ListFavoriteReply, error) {
	userClaims, err := s.checkLogin(ctx)
	if err != nil {
		return nil, err
	}

	opt := &model.OptionGetFavoriteList{
		Page:      int(req.Page),
		Size:      int(req.Size_),
		WithCount: true,
		QueryIn: map[string][]interface{}{
			"user_id": {userClaims.UserId},
		},
	}

	var (
		favorites []*pb.Favorite
		total     int64
	)
	if req.Type == 1 {
		favorites, total, err = s.dbModel.GetArticleFavoriteList(opt)
	} else {
		favorites, total, err = s.dbModel.GetDocumentFavoriteList(opt)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "获取失败:"+err.Error())
	}
	resp := &pb.ListFavoriteReply{
		Total:    total,
		Favorite: favorites,
	}
	return resp, nil
}

package model

func getPermissions() (permissions []Permission) {
	permissions = []Permission{
		{Title: "查看用户列表", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/ListUser"},
		{Title: "创建用户组", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/CreateGroup"},
		{Title: "查看友链列表", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/ListFriendlink"},
		{Title: "创建友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/CreateFriendlink"},
		{Title: "更新友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/UpdateFriendlink"},
		{Title: "删除友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/DeleteFriendlink"},
		{Title: "查看附件列表", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/ListAttachment"},
		{Title: "删除附件", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/DeleteAttachment"},
		{Title: "更新附件信息", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/UpdateAttachment"},
		{Title: "获取轮播图列表", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/ListBanner"},
		{Title: "上传轮播图", Description: "", Method: "POST", Path: "/api/v1/upload/banner"},
		{Title: "创建轮播图", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/CreateBanner"},
		{Title: "更新轮播图", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/UpdateBanner"},
		{Title: "删除轮播图", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/DeleteBanner"},
		{Title: "根据ID查询附件", Description: "", Method: "GRPC", Path: "/api.v1.AttachmentAPI/GetAttachment"},
		{Title: "根据ID查询轮播图", Description: "", Method: "GRPC", Path: "/api.v1.BannerAPI/GetBanner"},
		{Title: "根据ID查询友链", Description: "", Method: "GRPC", Path: "/api.v1.FriendlinkAPI/GetFriendlink"},
		{Title: "删除用户组", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/DeleteGroup"},
		{Title: "更新用户组", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/UpdateGroup"},
		{Title: "获取用户组列表", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/ListGroup"},
		{Title: "获取权限列表", Description: "", Method: "GRPC", Path: "/api.v1.PermissionAPI/ListPermission"},
		{Title: "根据ID查询权限", Description: "", Method: "GRPC", Path: "/api.v1.PermissionAPI/GetPermission"},
		{Title: "更新权限信息", Description: "", Method: "GRPC", Path: "/api.v1.PermissionAPI/UpdatePermission"},
		{Title: "获取用户组授权", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/GetGroupPermission"},
		{Title: "给用户组授权", Description: "", Method: "GRPC", Path: "/api.v1.GroupAPI/UpdateGroupPermission"},
		{Title: "获取系统配置列表", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/ListConfig"},
		{Title: "更新系统配置", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/UpdateConfig"},
		{Title: "新增用户", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/AddUser"},
		{Title: "设置用户密码与分组", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/SetUser"},
		{Title: "删除用户", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/DeleteUser"},
		{Title: "更新用户资料", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/UpdateUserProfile"},
		{Title: "根据ID查询用户", Description: "", Method: "GRPC", Path: "/api.v1.UserAPI/GetUser"},
		{Title: "查询文档分类列表", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/ListCategory"},
		{Title: "根据ID查询文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/GetCategory"},
		{Title: "新增文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/CreateCategory"},
		{Title: "更新文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/UpdateCategory"},
		{Title: "删除文档分类", Description: "", Method: "GRPC", Path: "/api.v1.CategoryAPI/DeleteCategory"},
		{Title: "获取文档列表", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/ListDocument"},
		{Title: "删除文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/DeleteDocument"},
		{Title: "根据ID查询文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/GetDocument"},
		{Title: "更新文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/UpdateDocument"},
		{Title: "获取回收站文档列表", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/ListRecycleDocument"},
		{Title: "恢复回收站文档", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/RecoverRecycleDocument"},
		{Title: "清空回收站文档", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/ClearRecycleDocument"},
		{Title: "删除回收站文档", Description: "", Method: "GRPC", Path: "/api.v1.RecycleAPI/DeleteRecycleDocument"},
		{Title: "查询文章列表", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/ListArticle"},
		{Title: "更新文章/单页", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/UpdateArticle"},
		{Title: "删除文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/DeleteArticle"},
		{Title: "后台创建文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/CreateArticle"},
		{Title: "上传文章图片和音视频", Description: "", Method: "POST", Path: "/api/v1/upload/article"},
		{Title: "上传文档分类封面", Description: "", Method: "POST", Path: "/api/v1/upload/category"},
		{Title: "上传配置图片文件", Description: "", Method: "POST", Path: "/api/v1/upload/config"},
		{Title: "获取评论列表", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/ListComment"},
		{Title: "获取单个评论", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/GetComment"},
		{Title: "批量审核评论", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/CheckComment"},
		{Title: "删除评论", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/DeleteComment"},
		{Title: "推荐文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/SetDocumentRecommend"},
		{Title: "查询举报列表", Description: "", Method: "GRPC", Path: "/api.v1.ReportAPI/ListReport"},
		{Title: "处理举报内容", Description: "", Method: "GRPC", Path: "/api.v1.ReportAPI/UpdateReport"},
		{Title: "删除举报内容", Description: "", Method: "GRPC", Path: "/api.v1.ReportAPI/DeleteReport"},
		{Title: "查看系统信息", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/GetStats"},
		{Title: "更新站点地图sitemap", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/UpdateSitemap"},
		{Title: "获取环境依赖检测", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/GetEnvs"},
		{Title: "更新头像", Description: "", Method: "POST", Path: "/api/v1/upload/avatar"},
		{Title: "一键重转失败文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/SetDocumentReconvert"},
		{Title: "批量修改文档分类", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/SetDocumentsCategory"},
		{Title: "编辑导航", Description: "", Method: "GRPC", Path: "/api.v1.NavigationAPI/UpdateNavigation"},
		{Title: "新增导航", Description: "", Method: "GRPC", Path: "/api.v1.NavigationAPI/CreateNavigation"},
		{Title: "删除导航", Description: "", Method: "GRPC", Path: "/api.v1.NavigationAPI/DeleteNavigation"},
		{Title: "获取用户处罚列表", Description: "", Method: "GRPC", Path: "/api.v1.PunishmentAPI/ListPunishment"},
		{Title: "获取单个处罚信息", Description: "", Method: "GRPC", Path: "/api.v1.PunishmentAPI/GetPunishment"},
		{Title: "取消用户处罚", Description: "", Method: "GRPC", Path: "/api.v1.PunishmentAPI/CancelPunishment"},
		{Title: "更新用户处罚", Description: "", Method: "GRPC", Path: "/api.v1.PunishmentAPI/UpdatePunishment"},
		{Title: "添加用户处罚", Description: "", Method: "GRPC", Path: "/api.v1.PunishmentAPI/CreatePunishment"},
		{Title: "获取广告列表", Description: "", Method: "GRPC", Path: "/api.v1.AdvertisementAPI/ListAdvertisement"},
		{Title: "添加广告", Description: "", Method: "GRPC", Path: "/api.v1.AdvertisementAPI/CreateAdvertisement"},
		{Title: "获取广告内容", Description: "", Method: "GRPC", Path: "/api.v1.AdvertisementAPI/GetAdvertisement"},
		{Title: "更新广告", Description: "", Method: "GRPC", Path: "/api.v1.AdvertisementAPI/UpdateAdvertisement"},
		{Title: "删除广告", Description: "", Method: "GRPC", Path: "/api.v1.AdvertisementAPI/DeleteAdvertisement"},
		{Title: "获取搜索记录列表", Description: "", Method: "GRPC", Path: "/api.v1.SearchRecordAPI/ListSearchRecord"},
		{Title: "删除搜索记录", Description: "", Method: "GRPC", Path: "/api.v1.SearchRecordAPI/DeleteSearchRecord"},
		{Title: "更新评论", Description: "", Method: "GRPC", Path: "/api.v1.CommentAPI/UpdateComment"},
		{Title: "文档审核", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/CheckDocument"},
		{Title: "下载待审文档", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/DownloadDocumentToBeReviewed"},
		{Title: "语言管理", Description: "", Method: "GRPC", Path: "/api.v1.LanguageAPI/ListLanguage"},
		{Title: "更新语言状态", Description: "", Method: "GRPC", Path: "/api.v1.LanguageAPI/UpdateLanguageStatus"},
		{Title: "删除语言", Description: "", Method: "GRPC", Path: "/api.v1.LanguageAPI/DeleteLanguage"},
		{Title: "更新语言", Description: "", Method: "GRPC", Path: "/api.v1.LanguageAPI/UpdateLanguage"},
		{Title: "批量设置文档语言", Description: "", Method: "GRPC", Path: "/api.v1.DocumentAPI/SetDocumentsLanguage"},
		{Title: "查看服务器硬件信息", Description: "", Method: "GRPC", Path: "/api.v1.ConfigAPI/GetDeviceInfo"},
		{Title: "查询回收站文章列表", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/ListRecycleArticle"},
		{Title: "文章批量分类", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/SetArticlesCategory"},
		{Title: "从回收站恢复文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/RestoreRecycleArticle"},
		{Title: "删除回收站文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/DeleteRecycleArticle"},
		{Title: "清空回收站文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/EmptyRecycleArticle"},
		{Title: "推荐文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/RecommendArticles"},
		{Title: "批量审批文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/CheckArticles"},
		{Title: "查阅所有状态文章", Description: "", Method: "GRPC", Path: "/api.v1.ArticleAPI/GetArticle"},
	}
	return
}

func getLangs() (langs []Language) {
	langs = []Language{
		{Language: "中文（简体）", Code: "zh-CN", Enable: true},
		{Language: "中文（繁体）", Code: "zh-TW", Enable: true},
		{Language: "英语", Code: "en", Enable: true},
		{Language: "法语", Code: "fr", Enable: true},
		{Language: "韩语", Code: "ko", Enable: true},
		{Language: "德语", Code: "de", Enable: true},
		{Language: "日语", Code: "ja", Enable: true},
		{Language: "俄语", Code: "ru", Enable: true},
		{Language: "阿尔巴尼亚语", Code: "sq", Enable: false},
		{Language: "南非荷兰语", Code: "af", Enable: false},
		{Language: "阿姆哈拉语", Code: "am", Enable: false},
		{Language: "阿拉伯语", Code: "ar", Enable: false},
		{Language: "亚美尼亚文", Code: "hy", Enable: false},
		{Language: "阿萨姆语", Code: "as", Enable: false},
		{Language: "艾马拉语", Code: "ay", Enable: false},
		{Language: "阿塞拜疆语", Code: "az", Enable: false},
		{Language: "班巴拉语", Code: "bm", Enable: false},
		{Language: "巴斯克语", Code: "eu", Enable: false},
		{Language: "白俄罗斯语", Code: "be", Enable: false},
		{Language: "孟加拉文", Code: "bn", Enable: false},
		{Language: "博杰普尔语", Code: "bho", Enable: false},
		{Language: "波斯尼亚语", Code: "bs", Enable: false},
		{Language: "保加利亚语", Code: "bg", Enable: false},
		{Language: "加泰罗尼亚语", Code: "ca", Enable: false},
		{Language: "宿务语", Code: "ceb", Enable: false},
		{Language: "科西嘉语", Code: "co", Enable: false},
		{Language: "克罗地亚语", Code: "hr", Enable: false},
		{Language: "捷克语", Code: "cs", Enable: false},
		{Language: "丹麦语", Code: "da", Enable: false},
		{Language: "迪维希语", Code: "dv", Enable: false},
		{Language: "多格来语", Code: "doi", Enable: false},
		{Language: "荷兰语", Code: "nl", Enable: false},
		{Language: "世界语", Code: "eo", Enable: false},
		{Language: "爱沙尼亚语", Code: "et", Enable: false},
		{Language: "埃维语", Code: "ee", Enable: false},
		{Language: "菲律宾语（塔加拉语）", Code: "fil", Enable: false},
		{Language: "芬兰语", Code: "fi", Enable: false},
		{Language: "弗里斯兰语", Code: "fy", Enable: false},
		{Language: "加利西亚语", Code: "gl", Enable: false},
		{Language: "格鲁吉亚语", Code: "ka", Enable: false},
		{Language: "希腊文", Code: "el", Enable: false},
		{Language: "瓜拉尼人", Code: "gn", Enable: false},
		{Language: "古吉拉特文", Code: "gu", Enable: false},
		{Language: "海地克里奥尔语", Code: "ht", Enable: false},
		{Language: "豪萨语", Code: "ha", Enable: false},
		{Language: "夏威夷语", Code: "haw", Enable: false},
		{Language: "希伯来语", Code: "he", Enable: false},
		{Language: "印地语", Code: "hi", Enable: false},
		{Language: "苗语", Code: "hmn", Enable: false},
		{Language: "匈牙利语", Code: "hu", Enable: false},
		{Language: "冰岛语", Code: "is", Enable: false},
		{Language: "伊博语", Code: "ig", Enable: false},
		{Language: "伊洛卡诺语", Code: "ilo", Enable: false},
		{Language: "印度尼西亚语", Code: "id", Enable: false},
		{Language: "爱尔兰语", Code: "ga", Enable: false},
		{Language: "意大利语", Code: "it", Enable: false},
		{Language: "爪哇语", Code: "jv", Enable: false},
		{Language: "卡纳达文", Code: "kn", Enable: false},
		{Language: "哈萨克语", Code: "kk", Enable: false},
		{Language: "高棉语", Code: "km", Enable: false},
		{Language: "卢旺达语", Code: "rw", Enable: false},
		{Language: "贡根语", Code: "gom", Enable: false},
		{Language: "克里奥尔语", Code: "kri", Enable: false},
		{Language: "库尔德语", Code: "ku", Enable: false},
		{Language: "库尔德语（索拉尼）", Code: "ckb", Enable: false},
		{Language: "吉尔吉斯语", Code: "ky", Enable: false},
		{Language: "老挝语", Code: "lo", Enable: false},
		{Language: "拉丁文", Code: "la", Enable: false},
		{Language: "拉脱维亚语", Code: "lv", Enable: false},
		{Language: "林格拉语", Code: "ln", Enable: false},
		{Language: "立陶宛语", Code: "lt", Enable: false},
		{Language: "卢干达语", Code: "lg", Enable: false},
		{Language: "卢森堡语", Code: "lb", Enable: false},
		{Language: "马其顿语", Code: "mk", Enable: false},
		{Language: "迈蒂利语", Code: "mai", Enable: false},
		{Language: "马尔加什语", Code: "mg", Enable: false},
		{Language: "马来语", Code: "ms", Enable: false},
		{Language: "马拉雅拉姆文", Code: "ml", Enable: false},
		{Language: "马耳他语", Code: "mt", Enable: false},
		{Language: "毛利语", Code: "mi", Enable: false},
		{Language: "马拉地语", Code: "mr", Enable: false},
		{Language: "梅泰语（曼尼普尔语）", Code: "mni-Mtei", Enable: false},
		{Language: "米佐语", Code: "lus", Enable: false},
		{Language: "蒙古文", Code: "mn", Enable: false},
		{Language: "缅甸语", Code: "my", Enable: false},
		{Language: "尼泊尔语", Code: "ne", Enable: false},
		{Language: "挪威语", Code: "no", Enable: false},
		{Language: "尼杨扎语（齐切瓦语）", Code: "ny", Enable: false},
		{Language: "奥里亚语（奥里亚）", Code: "or", Enable: false},
		{Language: "奥罗莫语", Code: "om", Enable: false},
		{Language: "普什图语", Code: "ps", Enable: false},
		{Language: "波斯语", Code: "fa", Enable: false},
		{Language: "波兰语", Code: "pl", Enable: false},
		{Language: "葡萄牙语（葡萄牙、巴西）", Code: "pt", Enable: false},
		{Language: "旁遮普语", Code: "pa", Enable: false},
		{Language: "克丘亚语", Code: "qu", Enable: false},
		{Language: "罗马尼亚语", Code: "ro", Enable: false},
		{Language: "萨摩亚语", Code: "sm", Enable: false},
		{Language: "梵语", Code: "sa", Enable: false},
		{Language: "苏格兰盖尔语", Code: "gd", Enable: false},
		{Language: "塞佩蒂语", Code: "nso", Enable: false},
		{Language: "塞尔维亚语", Code: "sr", Enable: false},
		{Language: "塞索托语", Code: "st", Enable: false},
		{Language: "修纳语", Code: "sn", Enable: false},
		{Language: "信德语", Code: "sd", Enable: false},
		{Language: "僧伽罗语", Code: "si", Enable: false},
		{Language: "斯洛伐克语", Code: "sk", Enable: false},
		{Language: "斯洛文尼亚语", Code: "sl", Enable: false},
		{Language: "索马里语", Code: "so", Enable: false},
		{Language: "西班牙语", Code: "es", Enable: false},
		{Language: "巽他语", Code: "su", Enable: false},
		{Language: "斯瓦希里语", Code: "sw", Enable: false},
		{Language: "瑞典语", Code: "sv", Enable: false},
		{Language: "塔加路语（菲律宾语）", Code: "tl", Enable: false},
		{Language: "塔吉克语", Code: "tg", Enable: false},
		{Language: "泰米尔语", Code: "ta", Enable: false},
		{Language: "鞑靼语", Code: "tt", Enable: false},
		{Language: "泰卢固语", Code: "te", Enable: false},
		{Language: "泰语", Code: "th", Enable: false},
		{Language: "蒂格尼亚语", Code: "ti", Enable: false},
		{Language: "宗加语", Code: "ts", Enable: false},
		{Language: "土耳其语", Code: "tr", Enable: false},
		{Language: "土库曼语", Code: "tk", Enable: false},
		{Language: "契维语（阿坎语）", Code: "ak", Enable: false},
		{Language: "乌克兰语", Code: "uk", Enable: false},
		{Language: "乌尔都语", Code: "ur", Enable: false},
		{Language: "维吾尔语", Code: "ug", Enable: false},
		{Language: "乌兹别克语", Code: "uz", Enable: false},
		{Language: "越南语", Code: "vi", Enable: false},
		{Language: "威尔士语", Code: "cy", Enable: false},
		{Language: "班图语", Code: "xh", Enable: false},
		{Language: "意第绪语", Code: "yi", Enable: false},
		{Language: "约鲁巴语", Code: "yo", Enable: false},
		{Language: "祖鲁语", Code: "zu", Enable: false},
	}
	return
}

package extra1

import (
	"context"
	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListLogic {
	return &FavoriteListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListLogic) FavoriteList(req *types.Douyin_favorite_list_request) (resp *types.Douyin_favorite_list_response, err error) {
	//userId, err := token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	//if err != nil {
	//	return &types.Douyin_favorite_list_response{
	//		Status_code: 1,
	//		Status_msg:  "token失效",
	//	}, nil
	//}
	//if userId != req.User_id {
	//	return &types.Douyin_favorite_list_response{
	//		Status_code: 2,
	//		Status_msg:  "用户失效",
	//	}, nil
	//}

	res, err := l.svcCtx.FavoriteModel.FindFavoriteVideos(l.ctx, req.User_id)

	if err != nil {
		return &types.Douyin_favorite_list_response{
			Status_code: 3,
			Status_msg:  "查询失败",
		}, nil
	}
	favorites := make([]types.Video, len(res))
	for i := range res {
		favorites[i] = types.Video{
			Id: res[i].VideoId,
			Author: types.User{
				Id:             res[i].UserId,
				Name:           res[i].UserName,
				Follow_count:   res[i].FollowCount,
				Follower_count: res[i].FollowerCount,
				Is_follow:      res[i].IsFollow == 0,
			},
			Play_url:       res[i].PlayUrl,
			Cover_url:      res[i].CoverUrl,
			Favorite_count: res[i].FavoriteCount,
			Comment_count:  res[i].CommentCount,
			Is_favorite:    true,
			Title:          res[i].Title,
		}
	}
	return &types.Douyin_favorite_list_response{
		Status_code: 0,
		Video_list:  favorites,
	}, nil
}

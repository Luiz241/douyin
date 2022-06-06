package basic

import (
	"context"
	"douyin/common/token"
	"time"

	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.Douyin_feed_request) (resp *types.Douyin_feed_response, err error) {
	userId, err := token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		userId = -1
	}

	res, err := l.svcCtx.VideoModel.FindVideosBefore(l.ctx, userId, time.Unix(req.Latest_time/1000, 0))
	if err != nil {
		return &types.Douyin_feed_response{
			Status_code: 3,
			Status_msg:  err.Error() + "视频数据库查询失败",
		}, nil
	}

	videos := make([]types.Video, len(res))

	for i := range res {
		videos[i] = types.Video{
			Id: res[i].VideoId,
			Author: types.User{
				Id:             res[i].UserId,
				Name:           res[i].UserName,
				Follow_count:   res[i].FollowCount,
				Follower_count: res[i].FollowerCount,
				Is_follow:      res[i].IsFollow == 0 && userId != -1,
			},
			Play_url: res[i].PlayUrl,
			//Cover_url:      res[i].CoverUrl,
			Cover_url:      "http://10.0.2.2:8888/douyin/static/spec2.jpg",
			Favorite_count: res[i].FavoriteCount,
			Comment_count:  res[i].CommentCount,
			Is_favorite:    res[i].IsFavorite == 0 && userId != -1,
			Title:          res[i].Title,
		}
	}

	return &types.Douyin_feed_response{
		Status_code: 0,
		Video_list:  videos,
	}, nil
}

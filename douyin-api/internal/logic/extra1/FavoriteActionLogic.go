package extra1

import (
	"context"
	"douyin/common/token"
	"douyin/douyin-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.Douyin_favorite_action_request) (resp *types.Douyin_favorite_action_response, err error) {
	userId, err := token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return &types.Douyin_favorite_action_response{
			Status_code: 1,
			Status_msg:  "token失效",
		}, nil
	}
	//if userId != req.User_id {
	//	return &types.Douyin_favorite_action_response{
	//		Status_code: 2,
	//		Status_msg:  "用户失效",
	//	}, nil
	//}
	if req.Action_type == 1 {
		err := l.svcCtx.FavoriteModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			_, err := l.svcCtx.FavoriteModel.TransInsert(ctx, session, &model.Favorite{
				UserId:  userId,
				VideoId: req.Video_id,
			})
			if err != nil {
				return err
			}
			err = l.svcCtx.VideoModel.TransUpdateFavoriteCounts(ctx, session, req.Video_id, true)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return &types.Douyin_favorite_action_response{
				Status_code: 3,
				Status_msg:  "点赞失败",
			}, nil
		}
	} else if req.Action_type == 2 {
		err := l.svcCtx.FavoriteModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			err := l.svcCtx.FavoriteModel.TransDelete(ctx, session, userId, req.Video_id)
			if err != nil {
				return err
			}
			err = l.svcCtx.VideoModel.TransUpdateFavoriteCounts(ctx, session, req.Video_id, false)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return &types.Douyin_favorite_action_response{
				Status_code: 4,
				Status_msg:  "取消点赞失败",
			}, nil
		}
	}
	return &types.Douyin_favorite_action_response{
		Status_code: 0,
	}, nil
}

package extra2

import (
	"context"
	"douyin/common/token"
	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"
	"douyin/douyin-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.Douyin_relation_action_request) (resp *types.Douyin_relation_action_response, err error) {
	userId, err := token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return &types.Douyin_relation_action_response{
			Status_code: 1,
			Status_msg:  "token失效",
		}, nil
	}

	if req.Action_type == 1 {
		err := l.svcCtx.RelationModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			_, err := l.svcCtx.RelationModel.TransInsert(ctx, session, &model.Relation{
				FolloweeId: req.To_user_id,
				FollowerId: userId,
			})
			if err != nil {
				return err
			}
			err = l.svcCtx.UserModel.TransUpdateFollower(ctx, session, req.To_user_id, true) //给被关注者改变粉丝
			if err != nil {
				return err
			}
			err = l.svcCtx.UserModel.TransUpdateFollow(ctx, session, userId, true) //给关注者改变关注数
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return &types.Douyin_relation_action_response{
				Status_code: 4,
				Status_msg:  "关注失败",
			}, nil
		}
	} else if req.Action_type == 2 {
		err := l.svcCtx.RelationModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			err := l.svcCtx.RelationModel.TransDelete(ctx, session, req.To_user_id, userId)
			if err != nil {
				return err
			}
			err = l.svcCtx.UserModel.TransUpdateFollower(ctx, session, req.To_user_id, false) //给被关注者改变粉丝
			if err != nil {
				return err
			}
			err = l.svcCtx.UserModel.TransUpdateFollow(ctx, session, userId, false) //给关注者改变关注数
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return &types.Douyin_relation_action_response{
				Status_code: 5,
				Status_msg:  "取关失败",
			}, nil
		}
	} else {
		return &types.Douyin_relation_action_response{
			Status_code: 3,
			Status_msg:  "错误操作",
		}, nil
	}
	return &types.Douyin_relation_action_response{
		Status_code: 0,
	}, nil
}

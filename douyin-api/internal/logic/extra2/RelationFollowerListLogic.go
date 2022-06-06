package extra2

import (
	"context"
	"douyin/common/token"

	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(req *types.Douyin_relation_follower_list_request) (resp *types.Douyin_relation_follower_list_response, err error) {
	_, err = token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return &types.Douyin_relation_follower_list_response{
			Status_code: 1,
			Status_msg:  "token失效",
		}, nil
	}
	//if userId != req.User_id {
	//	return &types.Douyin_relation_follower_list_response{
	//		Status_code: 2,
	//		Status_msg:  "用户失效",
	//	}, nil
	//}
	res, err := l.svcCtx.RelationModel.FindAllFollower(l.ctx, req.User_id)
	if err != nil {
		return &types.Douyin_relation_follower_list_response{
			Status_code: 3,
			Status_msg:  "查询失败",
		}, nil
	}
	users := make([]types.User, len(res))
	for i := range users {
		users[i] = types.User{
			Id:             res[i].UserId,
			Name:           res[i].UserName,
			Follow_count:   res[i].FollowCount,
			Follower_count: res[i].FollowerCount,
			Is_follow:      res[i].IsFollow == 0,
		}
	}
	return &types.Douyin_relation_follower_list_response{
		Status_code: 0,
		User_list:   users,
	}, nil
}

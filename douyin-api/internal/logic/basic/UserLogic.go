package basic

import (
	"context"
	"douyin/common/token"

	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.Douyin_user_request) (resp *types.Douyin_user_response, err error) {
	id, err := token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return &types.Douyin_user_response{
			Status_code: 1,
			Status_msg:  err.Error(),
		}, nil
	}

	if id != req.User_id {
		return &types.Douyin_user_response{
			Status_code: 2,
			Status_msg:  "用户已更改，请重新登录",
		}, nil
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.User_id)
	if err != nil {
		return &types.Douyin_user_response{
			Status_code: 3,
			Status_msg:  "用户查询失败",
		}, nil
	}

	respUser := types.User{
		Id:             user.UserId,
		Name:           user.UserName,
		Follow_count:   user.FollowCount,
		Follower_count: user.FollowerCount,
		Is_follow:      false,
	}

	return &types.Douyin_user_response{
		Status_code: 0,
		User:        respUser,
	}, nil
}

package basic

import (
	"context"
	"douyin/common/crypt"
	"douyin/common/token"
	"time"

	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.Douyin_user_login_request) (resp *types.Douyin_user_login_response, err error) {
	user, err := l.svcCtx.UserModel.FindOneByNameAndPassword(l.ctx, req.Username, crypt.Md5ByString(req.Password))
	if err != nil {
		return &types.Douyin_user_login_response{
			Status_code: 1,
			Status_msg:  "用户名密码不匹配，请重试",
		}, nil
	}

	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	myToken, err := token.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, expire, user.UserId)
	if err != nil {
		return &types.Douyin_user_login_response{
			Status_code: 2,
			Status_msg:  "token生成失败",
		}, nil
	}

	return &types.Douyin_user_login_response{
		Status_code: 0,
		User_id:     user.UserId,
		Token:       myToken,
	}, nil
}

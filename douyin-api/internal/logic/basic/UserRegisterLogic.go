package basic

import (
	"context"
	"douyin/common/token"
	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"
	"douyin/douyin-api/model"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.Douyin_user_register_request) (resp *types.Douyin_user_register_response, err error) {
	_, err = l.svcCtx.UserModel.FindOneByName(l.ctx, req.Username)
	// 用户已存在
	if err == nil {
		return &types.Douyin_user_register_response{
			Status_code: 1,
			Status_msg:  "用户已存在",
		}, nil
	}

	user := &model.User{
		UserName: req.Username,
		Password: req.Password,
	}

	res, err := l.svcCtx.UserModel.Insert(l.ctx, user)

	if err != nil {
		return &types.Douyin_user_register_response{
			Status_code: 2,
			Status_msg:  "用户创建失败",
		}, nil
	}
	id, _ := res.LastInsertId()
	now := time.Now().Unix()
	expire := l.svcCtx.Config.Auth.AccessExpire
	myToken, err := token.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, expire, id)

	if err != nil {
		return &types.Douyin_user_register_response{
			Status_code: 3,
			Status_msg:  "token生成失败",
		}, nil
	}
	return &types.Douyin_user_register_response{
		Status_code: 0,
		User_id:     id,
		Token:       myToken,
	}, nil
}

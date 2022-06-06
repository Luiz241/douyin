package svc

import (
	"douyin/douyin-api/internal/config"
	"douyin/douyin-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	VideoModel    model.VideoModel
	CommentModel  model.CommentModel
	FavoriteModel model.FavoriteModel
	RelationModel model.RelationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
		VideoModel:    model.NewVideoModel(sqlx.NewMysql(c.DB.DataSource)),
		CommentModel:  model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource)),
		FavoriteModel: model.NewFavoriteModel(sqlx.NewMysql(c.DB.DataSource)),
		RelationModel: model.NewRelationModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}

// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"path"

	basic "douyin/douyin-api/internal/handler/basic"
	extra1 "douyin/douyin-api/internal/handler/extra1"
	extra2 "douyin/douyin-api/internal/handler/extra2"
	"douyin/douyin-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/feed",
				Handler: basic.FeedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: basic.UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: basic.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/user",
				Handler: basic.UserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/publish/action",
				Handler: basic.PublishActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/publish/list",
				Handler: basic.PublishListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/favorite/action",
				Handler: extra1.FavoriteActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/favorite/list",
				Handler: extra1.FavoriteListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/action",
				Handler: extra1.CommentActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/comment/list",
				Handler: extra1.CommentListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/relation/action",
				Handler: extra2.RelationActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/follow/list",
				Handler: extra2.RelationFollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/relation/follower/list",
				Handler: extra2.RelationFollowerListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method: http.MethodGet,
				Path:   "/static/:file",
				Handler: func(svcCtx *svc.ServiceContext) http.HandlerFunc {
					return func(w http.ResponseWriter, r *http.Request) {
						type Request struct {
							File string `path:"file"`
						}
						var req Request
						if err := httpx.Parse(r, &req); err != nil {
							httpx.Error(w, err)
							return
						}
						http.ServeFile(w, r, path.Join(serverCtx.Config.StaticPath, req.File))
						//body, err := ioutil.ReadFile(path.Join(serverCtx.Config.StaticPath, req.File))
						//if err != nil {
						//	httpx.Error(w, err)
						//	return
						//}
						//
						//w.Write(body)
					}
				}(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)
}

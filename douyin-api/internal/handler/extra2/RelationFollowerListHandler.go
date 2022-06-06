package extra2

import (
	"net/http"

	"douyin/douyin-api/internal/logic/extra2"
	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFollowerListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Douyin_relation_follower_list_request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := extra2.NewRelationFollowerListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFollowerList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

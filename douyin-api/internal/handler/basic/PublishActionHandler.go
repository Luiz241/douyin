package basic

import (
	"net/http"

	"douyin/douyin-api/internal/logic/basic"
	"douyin/douyin-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req types.Douyin_publish_action_request
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.Error(w, err)
		//	return
		//}

		l := basic.NewPublishActionLogic(r, r.Context(), svcCtx)
		resp, err := l.PublishAction()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}

package extra1

import (
	"context"
	"douyin/common/token"
	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentListLogic) CommentList(req *types.Douyin_comment_list_request) (resp *types.Douyin_comment_list_response, err error) {
	userId, err := token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return &types.Douyin_comment_list_response{
			Status_code: 1,
			Status_msg:  "token失效",
		}, nil
	}
	//if userId != req.User_id {
	//	return &types.Douyin_comment_list_response{
	//		Status_code: 2,
	//		Status_msg:  "用户失效",
	//	}, nil
	//}

	res, err := l.svcCtx.CommentModel.FindCommentsByVideoId(l.ctx, userId, req.Video_id)

	if err != nil {
		return &types.Douyin_comment_list_response{
			Status_code: 3,
			Status_msg:  err.Error() + "查询失败",
		}, nil
	}
	comments := make([]types.Comment, len(res))
	for i := range res {
		comments[i] = types.Comment{
			Id: res[i].CommentId,
			User: types.User{
				Id:             res[i].UserId,
				Name:           res[i].Username,
				Follow_count:   res[i].FollowCount,
				Follower_count: res[i].FollowerCount,
				Is_follow:      res[i].IsFollow == 0,
			},
			Content:     res[i].Comment,
			Create_date: res[i].CreateTime.Format("01-02"),
		}
	}
	return &types.Douyin_comment_list_response{
		Status_code:  0,
		Comment_list: comments,
	}, nil
}

package extra1

import (
	"context"
	"douyin/common/token"
	"douyin/douyin-api/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentActionLogic {
	return &CommentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentActionLogic) CommentAction(req *types.Douyin_comment_action_request) (resp *types.Douyin_comment_action_response, err error) {
	userId, err := token.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return &types.Douyin_comment_action_response{
			Status_code: 1,
			Status_msg:  "token失效",
		}, nil
	}
	//if userId != req.User_id {
	//	return &types.Douyin_comment_action_response{
	//		Status_code: 2,
	//		Status_msg:  "用户失效",
	//	}, nil
	//}

	if req.Action_type == 1 {
		var commentId int64
		err = l.svcCtx.CommentModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			res, err := l.svcCtx.CommentModel.TransInsert(ctx, session, &model.Comment{
				UserId:  userId,
				Comment: req.Comment_text,
				VideoId: req.Video_id,
			})
			if err != nil {
				return err
			}
			err = l.svcCtx.VideoModel.TransUpdateCommentCounts(ctx, session, req.Video_id, true)
			if err != nil {
				return err
			}
			commentId, _ = res.LastInsertId()
			return nil
		})

		comment, err := l.svcCtx.CommentModel.FindCommentAndUser(l.ctx, userId, commentId)
		if err != nil {
			return &types.Douyin_comment_action_response{
				Status_code: 3,
				Status_msg:  "评论创建失效",
			}, nil
		}
		return &types.Douyin_comment_action_response{
			Comment: types.Comment{
				Id: comment.CommentId,
				User: types.User{
					Id:             comment.UserId,
					Name:           comment.Username,
					Follow_count:   comment.FollowCount,
					Follower_count: comment.FollowerCount,
					Is_follow:      comment.IsFollow == 0,
				},
				Content:     comment.Comment,
				Create_date: comment.CreateTime.Format("01-02"),
			},
		}, nil
	} else if req.Action_type == 2 {
		err = l.svcCtx.CommentModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
			err := l.svcCtx.CommentModel.TransDelete(l.ctx, session, req.Comment_id)
			if err != nil {
				return err
			}
			err = l.svcCtx.VideoModel.TransUpdateCommentCounts(l.ctx, session, req.Video_id, false)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return &types.Douyin_comment_action_response{
				Status_code: 4,
				Status_msg:  "评论删除失效",
			}, nil
		}
	}
	return
}

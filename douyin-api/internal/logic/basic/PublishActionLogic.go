package basic

import (
	"context"
	"douyin/common/token"
	"douyin/douyin-api/model"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"

	"douyin/douyin-api/internal/svc"
	"douyin/douyin-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20

type PublishActionLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		r:      r,
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishActionLogic) PublishAction() (resp *types.Douyin_publish_action_response, err error) {
	myToken := l.r.PostFormValue("token")

	userId, err := token.ParseToken(myToken, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return &types.Douyin_publish_action_response{
			Status_code: 1,
			Status_msg:  "token失效",
		}, nil
	}
	title := l.r.PostFormValue("title")
	l.r.ParseMultipartForm(maxFileSize)
	file, _, err := l.r.FormFile("data")
	if err != nil {
		return &types.Douyin_publish_action_response{
			Status_code: 2,
			Status_msg:  "文件传输失败",
		}, nil
	}
	fileName := strconv.FormatInt(userId, 10) + "_" + strconv.FormatInt(time.Now().Unix(), 10) + ".mp4"
	tempFile, err := os.Create(path.Join(l.svcCtx.Config.StaticPath, fileName))
	if err != nil {
		return &types.Douyin_publish_action_response{
			Status_code: 3,
			Status_msg:  "文件保存失败",
		}, nil
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)
	cmd := exec.Command("ffmpeg", "-i", path.Join(l.svcCtx.Config.StaticPath, fileName), "-vframes", "1",
		"-f", "image2", path.Join(l.svcCtx.Config.StaticPath, fileName[:len(fileName)-4]+".jpg"))
	if err = cmd.Start(); err != nil {
		logx.Error("封面提取失败")
		return &types.Douyin_publish_action_response{
			Status_code: 5,
			Status_msg:  "封面提取失败",
		}, nil
	}
	_, err = l.svcCtx.VideoModel.Insert(l.ctx, &model.Video{
		UserId:   userId,
		Title:    title,
		PlayUrl:  l.svcCtx.Config.PlayPath + fileName,
		CoverUrl: l.svcCtx.Config.PlayPath + fileName[:len(fileName)-4] + ".jpg",
	})
	if err != nil {
		return &types.Douyin_publish_action_response{
			Status_code: 4,
			Status_msg:  "数据库保存失败",
		}, nil
	}
	return &types.Douyin_publish_action_response{
		Status_code: 0,
	}, nil
}

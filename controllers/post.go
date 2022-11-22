package controllers

import (
	"strconv"
	"web_app/logic"
	"web_app/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数校验
	p := new(models.Post)
	err := c.ShouldBindJSON(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	id, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeneedLogin)
		return
	}
	p.AuthorID = id
	//2.创建帖子
	err = logic.CreatePost(p)
	if err != nil {
		zap.L().Error("logic.CreatePost(p) field", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

//获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	//1.获取参数（从url中获取id）
	id := c.Param("id")
	pid, err2 := strconv.ParseInt(id, 10, 64)
	if err2 != nil {
		zap.L().Error("get post detail with invalid param1", zap.Error(nil))
		ResponseError(c, CodeInvalidParam)
		return
	}
	if len(id) == 0 {
		zap.L().Error("get post detail with invalid param2", zap.Error(nil))
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetPostDetail(pid)
	if err != nil {
		zap.L().Error("get post detail with invalid param3", zap.Error(nil))
		ResponseError(c, CodeInvalidParam)
		return
	}
	ResponseSuccess(c, data)
}
func GetPostListHandler(c *gin.Context) {
	offset, limit := GetPageInfo(c)
	data, err := logic.GetPostList(limit, offset)
	if err != nil {
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}

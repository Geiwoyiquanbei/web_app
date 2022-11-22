package controllers

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserID = "uerid"

var ERRorUserNotLogin = errors.New("用户未登录")

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserID)
	if !ok {
		err = ERRorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ERRorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int64, int64) {
	//获取分页参数
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	var (
		limit  int64
		offset int64
	)
	offset, _ = strconv.ParseInt(offsetStr, 10, 64)
	limit, _ = strconv.ParseInt(limitStr, 10, 64)
	return offset, limit
}

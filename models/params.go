package models

type ParamSignUp struct {
	UserName   string `json:"username"  binding:"required"`
	Password   string `json:"password"  binding:"required"`
	RePassword string `json:"repassword"  binding:"required,eqfield=Password"`
}
type ParamLogin struct {
	UserName string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int    `json:"direction,string" binding:"oneof=1 0 -1 "` //赞成票是1，反对票是-1 取消投票就是0
}

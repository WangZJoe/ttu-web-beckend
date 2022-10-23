package apiv1

import "github.com/gogf/gf/v2/frame/g"

type UserLoginReq struct {
	g.Meta   `path:"/user/login" tags:"User" method:"post" summary:"用户登录"`
	Username string `json:"username" v:"required" title:"用户名"`
	Password string `json:"password" v:"required" title:"密码"`
	Token    string `json:"token" v:"required" title:"token身份码"`
}

type UserLoginRes struct {
	Code     string `json:"code" title:"状态code" dc:"成功是1，失败是0"`
	Status   string `json:"status" title:"登录情况"`
	UserGrid string
}

type UserChangePasswordReq struct {
	g.Meta      `path:"/user/changePassword" tags:"User" method:"post" summary:"用户修改密码"`
	Username    string `json:"username" v:"required" title:"用户名"`
	OldPassword string `json:"oldPassword" v:"required" title:"旧密码"`
	NewPassword string `json:"newPassword" v:"required" title:"新密码"`
}

type UserChangePasswordRes struct {
	Code   string `json:"code" title:"状态code" dc:"成功是1，失败是0"`
	Status string `json:"status" title:"修改情况"`
}

package utility

import "ttu-backend/internal/utility/structs"

var bizCode structs.BizCode

var (
	BizCodeSystemError             = bizCode.NewBizCode(9999, "系统发生错误，请联系在线客服", "")
	BizCodeOK                      = bizCode.NewBizCode(1001, "OK", "")
	BizCodeUserNameIsExisted       = bizCode.NewBizCode(1002, "用户名已存在", "")
	BizCodeUserNameIsNotExisted    = bizCode.NewBizCode(1003, "用户名不存在", "")
	BizCodeUserLoginPasswordFailed = bizCode.NewBizCode(1004, "密码错误", "")
	BizCodeNewTokenFailed          = bizCode.NewBizCode(1005, "Token 生成失败，请联系在线客服", "")
	BizCodeVerifyTokenFailed       = bizCode.NewBizCode(1006, "Token 验证失败，请联系在线客服", "")
	BizCodeTokenIsInvalid          = bizCode.NewBizCode(1006, "Token 失效，请重新登录", "")
)

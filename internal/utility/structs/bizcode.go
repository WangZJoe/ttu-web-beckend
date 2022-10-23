package structs

import (
	"fmt"
)

// BizCode 业务状态码
type BizCode struct {
	code    int
	message string
	detail  interface{}
}

// NewBizCode 创建业务状态码
func (c *BizCode) NewBizCode(code int, message string, detail interface{}) BizCode {
	return BizCode{
		code:    code,
		message: message,
		detail:  detail,
	}
}

// Code 获取业务状态码的 code 字段值
func (c BizCode) Code() int {
	return c.code
}

// Message 获取业务状态码的 message 字段值
func (c BizCode) Message() string {
	return c.message
}

// Detail 获取业务状态码的 detail 字段值
func (c BizCode) Detail() interface{} {
	return c.detail
}

// String 格式化输出业务状态码
func (c BizCode) String() string {
	if c.detail != nil {
		return fmt.Sprintf(`%d:%s %v`, c.code, c.message, c.detail)
	}
	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}

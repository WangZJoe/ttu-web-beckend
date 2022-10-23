package utility

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"ttu-backend/internal/consts"
	"ttu-backend/internal/utility/structs"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

type (
	sMiddleware struct{}
)

var (
	insMiddleware = sMiddleware{}
)

func Middleware() *sMiddleware {
	return &insMiddleware
}

// Ctx 加载 Context 中间件
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	Context().Init(r, &structs.Context{})

	r.Middleware.Next()
}

// HandlerResponse 请求返回数据处理中间件
func (s *sMiddleware) HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		msg  string
		ctx  = r.Context()
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
		msg = http.StatusText(r.Response.Status)
		switch r.Response.Status {
		case http.StatusNotFound:
			code = gcode.CodeNotFound
		case http.StatusForbidden:
			code = gcode.CodeNotAuthorized
		default:
			code = gcode.CodeUnknown
		}
	} else if bizCode, ok := code.(structs.BizCode); ok {
		code = bizCode
		msg = bizCode.Message()
	} else {
		code = BizCodeOK
		msg = BizCodeOK.Message()
	}
	writeJson(r, ctx, code.Code(), msg, res)
}

// Auth 用户认证中间件
func (s *sMiddleware) Auth(r *ghttp.Request) {
	var (
		err error
		ctx = r.Context()
	)

	authorization := r.GetHeader(consts.JwtHeaderKey)

	jwtIsInvalid, err := Jwt().ParseJwt(ctx, authorization)
	if !jwtIsInvalid {
		code := gerror.Code(err)
		writeJson(r, ctx, code.Code(), code.Message(), nil)
		return
	}

	r.Middleware.Next()
}

func writeJson(r *ghttp.Request, ctx context.Context, code int, msg string, res interface{}) {
	internalErr := r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code,
		Message: msg,
		Data:    res,
	})
	if internalErr != nil {
		// 系统错误（Json 写入 Response 失败），需要打印日志
		g.Log().Error(ctx, internalErr)
	}
}

// JsonOutputsForLogger is for JSON marshaling in sequence.
type JsonOutputsForLogger struct {
	Time       string `json:"time"`       // Logging time, which is the time that logging triggers.
	Level      string `json:"level"`      // Using level, like LEVEL_INFO, LEVEL_ERRO, etc.
	CallerPath string `json:"callerPath"` // The source file path and its line number that calls logging.
	Content    string `json:"content"`    // Content is the main logging content that passed by you.
}

// LoggingJsonHandler is a example handler for logging JSON format content.
var LoggingJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	jsonForLogger := JsonOutputsForLogger{
		Time:       in.TimeFormat,
		Level:      in.LevelFormat,
		CallerPath: in.CallerPath,
		Content:    in.Content,
	}
	jsonBytes, err := json.Marshal(jsonForLogger)
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
		return
	}
	in.Buffer.Write(jsonBytes)
	in.Buffer.WriteString("\n")
	in.Next()
}

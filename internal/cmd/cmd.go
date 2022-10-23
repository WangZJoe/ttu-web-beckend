package cmd

import (
	"context"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gsession"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/protocol/goai"
	"time"
	"ttu-backend/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"ttu-backend/internal/handler"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",

		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetSessionMaxAge(time.Minute)
			s.SetSessionStorage(gsession.NewStorageMemory())
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.ALL("/set", func(r *ghttp.Request) {
					r.Session.MustSet("time", gtime.Timestamp())
					r.Response.Write("ok")
				})
				group.ALL("/get", func(r *ghttp.Request) {
					r.Response.Write(r.Session.Data())
				})
				group.ALL("/del", func(r *ghttp.Request) {
					_ = r.Session.RemoveAll()
					r.Response.Write("ok")
				})
			})
			s.Group("/api", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					handler.Base,
					handler.User,
					//handler.Data,
				)
			})

			enhanceOpenAPIDoc(s)
			//g.SetDebug(true)
			s.SetIndexFolder(true)
			if gfile.Exists("dist") {
				s.AddStaticPath("/", "/dist")
				s.AddSearchPath("/dist")
			}
			s.EnableHTTPS("/https/server.crt", "/https/server.key")
			//s.SetPort(8199)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "",
			URL:  "https://",
		},
	}

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		{Name: consts.OpenAPITagNameUser},
		//{Name: consts.OpenAPITagNameChat},
	}
}

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

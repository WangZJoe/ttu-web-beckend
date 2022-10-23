package utility

import (
	"context"

	"ttu-backend/internal/consts"
	"ttu-backend/internal/utility/structs"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sContext struct{}
)

var (
	insContext = sContext{}
)

func Context() *sContext {
	return &insContext
}

// Init 初始化 Context
func (s *sContext) Init(r *ghttp.Request, customCtx *structs.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

// Get 获取 Context
func (s *sContext) Get(ctx context.Context) *structs.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*structs.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 设置 User Context
func (s *sContext) SetUser(ctx context.Context, ctxUser *structs.ContextUser) {
	s.Get(ctx).User = ctxUser
}

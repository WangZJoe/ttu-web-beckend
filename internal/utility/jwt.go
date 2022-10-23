package utility

import (
	"context"
	"encoding/json"
	"github.com/cristalhq/jwt/v4"
	"strings"
	"time"

	"ttu-backend/internal/consts"
	"ttu-backend/internal/utility/structs"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sJwt struct{}

var insJwt = sJwt{}

func Jwt() *sJwt {
	return &insJwt
}

/*
   iss (issuer)：签发人
   exp (expiration time)：过期时间
   sub (subject)：主题
   aud (audience)：受众
   nbf (Not Before)：生效时间
   iat (Issued At)：签发时间
   jti (JWT ID)：编号
*/

// NewJwt 生成 JWT
func (s *sJwt) NewJwt(ctx context.Context, userGuid, username string) (string, error) {
	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(consts.JwtPrivateKey))
	if err != nil {
		// 系统错误（生成 Token 失败），需要打印日志
		g.Log().Error(ctx, err)
		return "", gerror.NewCode(BizCodeNewTokenFailed)
	}

	expiresAtTime := jwt.NewNumericDate(time.Now().Add(consts.JwtExpTime))
	claims := &jwt.RegisteredClaims{
		Issuer:    consts.JwtIssuerName,
		ExpiresAt: expiresAtTime,
	}
	myClaims := &structs.JwtClaims{
		UserGuid:         userGuid,
		Username:         username,
		RegisteredClaims: claims,
	}

	token, err := jwt.NewBuilder(signer).Build(myClaims)
	if err != nil {
		// 系统错误（生成 Token 失败），需要打印日志
		g.Log().Error(ctx, err)
		return "", gerror.NewCode(BizCodeNewTokenFailed)
	}

	return token.String(), nil
}

// ParseJwt 解析验证 JWT
func (s *sJwt) ParseJwt(ctx context.Context, token string) (bool, error) {
	if !strings.HasPrefix(token, consts.JwtHeaderPrefix) {
		// 业务错误（Token 格式错误），不需要打印日志
		return false, gerror.NewCode(BizCodeVerifyTokenFailed)
	}

	token = strings.TrimPrefix(token, consts.JwtHeaderPrefix)

	verifierHS, err := jwt.NewVerifierHS(jwt.HS256, []byte(consts.JwtPrivateKey))
	if err != nil {
		// 系统错误（生成 Token 验证器失败），需要打印日志
		g.Log().Error(ctx, err)
		return false, gerror.NewCode(BizCodeVerifyTokenFailed)
	}

	parsedToken, err := jwt.Parse([]byte(token), verifierHS)
	if err != nil {
		// 系统错误（Token 验证失败），需要打印日志
		g.Log().Error(ctx, err)
		return false, gerror.NewCode(BizCodeVerifyTokenFailed)
	}

	var newClaims structs.JwtClaims

	err = json.Unmarshal(parsedToken.Claims(), &newClaims)
	if err != nil {
		// 系统错误（Token 解析失败），需要打印日志
		g.Log().Error(ctx, err)
		return false, gerror.NewCode(BizCodeVerifyTokenFailed)
	}

	if !newClaims.IsValidExpiresAt(time.Now()) {
		// 业务错误（Token 失效），不需要打印日志
		return false, gerror.NewCode(BizCodeTokenIsInvalid)
	}

	Context().SetUser(ctx, &structs.ContextUser{
		UserGuid: newClaims.UserGuid,
		Username: newClaims.Username,
	})

	return true, nil
}

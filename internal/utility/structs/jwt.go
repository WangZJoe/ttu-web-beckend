package structs

import "github.com/cristalhq/jwt/v4"

type JwtClaims struct {
	UserGuid string
	Username string
	*jwt.RegisteredClaims
}

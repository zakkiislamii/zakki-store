package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY_Admin = []byte("erfiw9ergher9guheqr9ubfgghrtjtyjtyhjfgyhjhau9frbgn9effsegsgserrugh9erugher9gher9ghaegur9hgrfnvr")

type JWTAdmin struct {
	Username string
	jwt.RegisteredClaims
}

package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEYY = []byte("erfiw9ergher9guheqr9ubfghau9frbgn9effsegsgserrugh9erugher9gher9ghaegur9hgrfnvr")

type JWTClaimm struct {
	Username string
	jwt.RegisteredClaims
}

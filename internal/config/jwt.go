package config

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JWTConfig struct {
	Issuer           string
	Secret           string
	Expiration       uint
	SecretRefresh    string
	ExpirationRefreh uint
}

type JWTClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewJWTManager(viper *viper.Viper) *JWTConfig {
	issuer := viper.GetString("JWT_ISSUER")
	secret := viper.GetString("JWT_SECRET")
	exp := viper.GetUint("JWT_EXPIRATION")
	secretRefresh := viper.GetString("JWT_SECRET_REFRESH")
	expRefresh := viper.GetUint("JWT_EXPIRATION_REFRESH")

	return &JWTConfig{
		Issuer:           issuer,
		Secret:           secret,
		Expiration:       exp,
		SecretRefresh:    secretRefresh,
		ExpirationRefreh: expRefresh,
	}
}

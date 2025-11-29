package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rafaeldepontes/auth-go/internal/errorhandler"
)

type JwtBuilder struct {
	secretKey string
}

func NewJwtBuilder(secretKey string) *JwtBuilder {
	return &JwtBuilder{secretKey}
}

func (builder JwtBuilder) GenerateToken(id uint, email string, duration time.Duration) (string, *UserClaims, error) {
	var userClaims *UserClaims
	userClaims, err := NewUserClaims(id, email, duration)
	if err != nil {
		return "", nil, err
	}

	var tokenJwt *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodES256, userClaims)
	token, err := tokenJwt.SignedString(builder.secretKey)
	if err != nil {
		return "", nil, err
	}

	return token, userClaims, nil
}

func (builder JwtBuilder) VerifyToken(token string) (*UserClaims, error) {
	var userClaims *UserClaims
	var tokenJwt *jwt.Token
	tokenJwt, err := jwt.ParseWithClaims(token, userClaims, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, errorhandler.ErrorInvalidTokenSigningMethod
		}
		return []byte(builder.secretKey), nil
	})

	if err != nil {
		return nil, errorhandler.ErrorParsingToken
	}

	userClaims, ok := tokenJwt.Claims.(*UserClaims)
	if !ok {
		return nil, errorhandler.ErrorInvalidTokenClaim
	}

	return userClaims, nil
}

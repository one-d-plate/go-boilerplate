package pkg

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/one-d-plate/go-boilerplate.git/src/consts"
)

func ParseJwtToken(token string, pub *rsa.PublicKey, claims jwt.Claims) (jwt.MapClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(consts.ErrInvalidSigningMethod)
		}
		return pub, nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New(consts.ErrInvalidSigningMethod)
	}

	return parsedToken.Claims.(jwt.MapClaims), nil
}

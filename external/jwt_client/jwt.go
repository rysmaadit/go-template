package jwt_client

import (
	"fmt"

	"github.com/rysmaadit/go-template/contract"
	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/rysmaadit/go-template/common/errors"
)

type jwtClient struct{}

func New() *jwtClient {
	return &jwtClient{}
}

type JWTClientInterface interface {
	GenerateTokenStringWithClaims(jwtClaims contract.JWTMapClaim, secret string) (string, error)
	ParseTokenWithClaims(tokenString string, claims jwt.MapClaims, secret string) error
}

func (j *jwtClient) GenerateTokenStringWithClaims(jwtClaims contract.JWTMapClaim, secret string) (string, error) {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	token, err := at.SignedString([]byte(secret))

	if err != nil {
		errMsg := fmt.Sprintf("error signed JWT credentials: %v", err)
		return "", errors.New(errMsg)
	}

	return token, nil
}

func (j *jwtClient) ParseTokenWithClaims(tokenString string, claims jwt.MapClaims, secret string) error {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		log.Error("error verified token", err)
		return errors.NewUnauthorizedError("error claim token")
	}

	if !token.Valid {
		log.Error("error invalid token", err)
		return errors.NewUnauthorizedError("invalid token")
	}

	return nil
}

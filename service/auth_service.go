package service

import (
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rysmaadit/go-template/common/errors"
	"github.com/rysmaadit/go-template/config"
	"github.com/rysmaadit/go-template/contract"
	"github.com/rysmaadit/go-template/external/jwt_client"
	log "github.com/sirupsen/logrus"
)

type authService struct {
	appConfig *config.Config
	jwtClient jwt_client.JWTClientInterface
}

type AuthServiceInterface interface {
	GetToken() (*contract.GetTokenResponseContract, error)
	VerifyToken(req *contract.ValidateTokenRequestContract) (*contract.JWTMapClaim, error)
}

func NewAuthService(appConfig *config.Config, jwtClient jwt_client.JWTClientInterface) *authService {
	return &authService{
		appConfig: appConfig,
		jwtClient: jwtClient,
	}
}

func (s *authService) GetToken() (*contract.GetTokenResponseContract, error) {
	atClaims := contract.JWTMapClaim{
		Authorized: true,
		RequestID:  uuid.New().String(),
	}

	token, err := s.jwtClient.GenerateTokenStringWithClaims(atClaims, s.appConfig.JWTSecret)

	if err != nil {
		errMsg := fmt.Sprintf("error signed JWT credentials: %v", err)
		log.Errorf(errMsg)
		return nil, errors.NewInternalError(err, errMsg)
	}

	return &contract.GetTokenResponseContract{Token: token}, err
}

func (s *authService) VerifyToken(req *contract.ValidateTokenRequestContract) (*contract.JWTMapClaim, error) {
	claims := jwt.MapClaims{}

	err := s.jwtClient.ParseTokenWithClaims(req.Token, claims, s.appConfig.JWTSecret)

	if err != nil {
		log.Errorln(err)
		return nil, errors.NewUnauthorizedError("invalid parse token with claims")
	}

	authorized := fmt.Sprintf("%v", claims["authorized"])
	requestID := fmt.Sprintf("%v", claims["requestID"])

	if authorized == "" || requestID == "" {
		return nil, errors.NewUnauthorizedError("invalid payload")
	}

	ok, err := strconv.ParseBool(authorized)

	if err != nil || !ok {
		log.Errorln(err)
		return nil, errors.NewUnauthorizedError("invalid payload")
	}

	resp := &contract.JWTMapClaim{
		Authorized:     claims["authorized"].(bool),
		RequestID:      claims["requestID"].(string),
		StandardClaims: jwt.StandardClaims{},
	}

	return resp, nil
}

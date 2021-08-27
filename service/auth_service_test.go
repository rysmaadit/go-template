package service

import (
	"reflect"
	"testing"

	"github.com/rysmaadit/go-template/common/errors"

	"github.com/rysmaadit/go-template/config"
	"github.com/rysmaadit/go-template/contract"
	"github.com/rysmaadit/go-template/external/jwt_client"
	"github.com/rysmaadit/go-template/external/mocks"
	mockVal "github.com/stretchr/testify/mock"
)

func Test_authService_GetToken(t *testing.T) {
	const (
		dummyTokenSecret  = "JWT_DUMMY_SECRET"
		generatedJWTToken = "token"
	)

	type fields struct {
		appConfig *config.Config
		jwtClient jwt_client.JWTClientInterface
	}
	tests := []struct {
		name    string
		fields  fields
		mockJWT func(mock *mocks.JWTClientInterface)
		want    *contract.GetTokenResponseContract
		wantErr bool
	}{
		{
			name: "given nothing will generate JWT token",
			mockJWT: func(mock *mocks.JWTClientInterface) {
				mock.On("GenerateTokenStringWithClaims", mockVal.AnythingOfType("contract.JWTMapClaim"), dummyTokenSecret).
					Return(generatedJWTToken, nil)
			},
			want:    &contract.GetTokenResponseContract{Token: generatedJWTToken},
			wantErr: false,
		},
		{
			name: "given nothing, error generate JWT token, should return error",
			mockJWT: func(mock *mocks.JWTClientInterface) {
				mock.On("GenerateTokenStringWithClaims", mockVal.AnythingOfType("contract.JWTMapClaim"), dummyTokenSecret).
					Return("", errors.New("err"))
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockJWT := new(mocks.JWTClientInterface)
			appConfig := &config.Config{
				JWTSecret: dummyTokenSecret,
			}

			s := &authService{
				appConfig: appConfig,
				jwtClient: mockJWT,
			}

			tt.mockJWT(mockJWT)

			got, err := s.GetToken()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authService_VerifyToken(t *testing.T) {
	const (
		dummyTokenSecret = "JWT_DUMMY_SECRET"
		sampleToken      = "token"
	)

	validateTokenContract := &contract.ValidateTokenRequestContract{Token: sampleToken}

	type fields struct {
		appConfig *config.Config
		jwtClient jwt_client.JWTClientInterface
	}
	type args struct {
		req *contract.ValidateTokenRequestContract
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mockJWT func(mock *mocks.JWTClientInterface)
		want    *contract.JWTMapClaim
		wantErr bool
	}{
		{
			name: "given invalid token should return error",
			args: args{req: validateTokenContract},
			mockJWT: func(mock *mocks.JWTClientInterface) {
				mock.On("ParseTokenWithClaims", sampleToken, mockVal.AnythingOfType("jwt.MapClaims"), dummyTokenSecret).
					Return(nil)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockJWT := new(mocks.JWTClientInterface)
			appConfig := &config.Config{
				JWTSecret: dummyTokenSecret,
			}

			s := &authService{
				appConfig: appConfig,
				jwtClient: mockJWT,
			}

			tt.mockJWT(mockJWT)
			got, err := s.VerifyToken(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VerifyToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

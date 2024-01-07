package service

import (
	"nursing_api/internal/domain/auth"
	"nursing_api/internal/domain/auth/dto"
	authUseCase "nursing_api/internal/domain/auth/usecase"
	userDto "nursing_api/internal/domain/user/dto"
	userUseCase "nursing_api/internal/domain/user/usecase"
	"nursing_api/pkg/jwt"
)

type authService struct {
	userUseCase userUseCase.UserUseCase
	jwtClient   *jwt.JwtClient
}

func NewAuthService(
	userUseCase userUseCase.UserUseCase,
	jwtClient *jwt.JwtClient,
) authUseCase.AuthUseCase {
	return &authService{
		userUseCase: userUseCase,
		jwtClient:   jwtClient,
	}
}

func (a authService) SignIn(req *dto.SignInRequest) *dto.SignInResponse {
	//TODO implement me
	panic("implement me")
}

func (a authService) SignUp(req *dto.SignUpRequest) *dto.SignUpResponse {
	registerDto := &userDto.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	response := a.userUseCase.RegisterUser(registerDto)
	if !response.Success {
		return dto.FailSignUp(response.Message, response.Error)
	}

	tokenRequest := &jwt.TokenRequest{
		ID:          response.User.ID.String(),
		Credentials: []string{},
	}
	jwtToken, err := a.jwtClient.Generator.GenerateNewTokens(tokenRequest)
	if err != nil {
		return dto.FailSignUp("토큰 생성에 실패하였습니다.", err)
	}

	token := &auth.Token{
		AccessToken:  jwtToken.AccessToken,
		RefreshToken: jwtToken.RefreshToken,
	}
	return dto.OkSignUp(token)
}

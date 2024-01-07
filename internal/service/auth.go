package service

import (
	"nursing_api/internal/domain/auth"
	"nursing_api/internal/domain/auth/dto"
	"nursing_api/internal/domain/auth/repository"
	authUseCase "nursing_api/internal/domain/auth/usecase"
	userDto "nursing_api/internal/domain/user/dto"
	userUseCase "nursing_api/internal/domain/user/usecase"
	"nursing_api/pkg/jwt"
)

type authService struct {
	userUseCase    userUseCase.UserUseCase
	authRepository repository.AuthRepository
	jwtClient      *jwt.JwtClient
}

func NewAuthService(
	userUseCase userUseCase.UserUseCase,
	authRepository repository.AuthRepository,
	jwtClient *jwt.JwtClient,
) authUseCase.AuthUseCase {
	return &authService{
		userUseCase:    userUseCase,
		authRepository: authRepository,
		jwtClient:      jwtClient,
	}
}

func (a authService) SignIn(req *dto.SignInRequest) *dto.SignInResponse {
	loginDto := &userDto.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	// 로그인 처리
	response := a.userUseCase.VerifyUser(loginDto)
	if !response.Success {
		return dto.FailSignIn(response.Message, response.Error)
	}
	userId := response.User.ID

	// 토큰 생성
	tokenRequest := &jwt.TokenRequest{
		ID:          userId.String(),
		Credentials: []string{},
	}
	jwtToken, err := a.jwtClient.Generator.GenerateNewTokens(tokenRequest)
	if err != nil {
		return dto.FailSignIn("토큰 생성에 실패하였습니다.", err)
	}

	// 토큰 저장
	newToken := &auth.Token{
		AccessToken:         jwtToken.AccessToken,
		AccessTokenExpires:  jwtToken.AccessTokenExpires,
		RefreshToken:        jwtToken.RefreshToken,
		RefreshTokenExpires: jwtToken.RefreshTokenExpires,
	}
	a.authRepository.DeleteToken(userId)
	savedToken, err := a.authRepository.SaveToken(userId, newToken)
	if err != nil {
		return dto.FailSignIn("토큰 저장에 실패하였습니다.", err)
	}

	return dto.OkSignIn(savedToken)
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

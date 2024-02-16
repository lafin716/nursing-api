package auth

import (
	"errors"
	"github.com/google/uuid"
	"nursing_api/internal/domain/user"
	"nursing_api/pkg/jwt"
)

type authService struct {
	userUseCase    user.UserUseCase
	authRepository AuthRepository
	jwtClient      *jwt.JwtClient
}

func NewService(
	userUseCase user.UserUseCase,
	authRepository AuthRepository,
	jwtClient *jwt.JwtClient,
) AuthUseCase {
	return &authService{
		userUseCase:    userUseCase,
		authRepository: authRepository,
		jwtClient:      jwtClient,
	}
}

func (a authService) SignIn(req *SignInRequest) *SignInResponse {
	loginDto := &user.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	// 로그인 처리
	response := a.userUseCase.VerifyUser(loginDto)
	if !response.Success {
		return FailSignIn(response.Message, response.Error)
	}
	userId := response.User.ID

	// 토큰 생성
	tokenRequest := &jwt.TokenRequest{
		ID:          userId.String(),
		Credentials: []string{},
	}
	jwtToken, err := a.jwtClient.Generator.GenerateNewTokens(tokenRequest)
	if err != nil {
		return FailSignIn("토큰 생성에 실패하였습니다.", err)
	}

	// 토큰 저장
	newToken := &Token{
		AccessToken:         jwtToken.AccessToken,
		AccessTokenExpires:  jwtToken.AccessTokenExpires,
		RefreshToken:        jwtToken.RefreshToken,
		RefreshTokenExpires: jwtToken.RefreshTokenExpires,
		AutoLogin:           req.AutoLogin,
	}
	a.authRepository.DeleteToken(userId)
	savedToken, err := a.authRepository.SaveToken(userId, newToken)
	if err != nil {
		return FailSignIn("토큰 저장에 실패하였습니다.", err)
	}

	return OkSignIn(savedToken)
}

func (a authService) SignUp(req *SignUpRequest) *SignUpResponse {
	registerDto := &user.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	response := a.userUseCase.RegisterUser(registerDto)
	if !response.Success {
		return FailSignUp(response.Message, response.Error)
	}

	tokenRequest := &jwt.TokenRequest{
		ID:          response.User.ID.String(),
		Credentials: []string{},
	}
	jwtToken, err := a.jwtClient.Generator.GenerateNewTokens(tokenRequest)
	if err != nil {
		return FailSignUp("토큰 생성에 실패하였습니다.", err)
	}

	token := &Token{
		AccessToken:  jwtToken.AccessToken,
		RefreshToken: jwtToken.RefreshToken,
	}
	return OkSignUp(token)
}

func (a authService) SignOut(userId uuid.UUID) error {
	result, err := a.authRepository.DeleteToken(userId)
	if err != nil {
		return err
	}

	if !result {
		return errors.New("이미 로그아웃 되었습니다.")
	}

	return nil
}

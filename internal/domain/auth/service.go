package auth

import (
	"errors"
	"github.com/google/uuid"
	"nursing_api/internal/common/dto"
	"nursing_api/internal/common/response"
	"nursing_api/internal/domain/user"
	"nursing_api/pkg/jwt"
)

type authService struct {
	userUseCase    user.UseCase
	authRepository Repository
	jwtClient      *jwt.JwtClient
}

type UseCase interface {
	SignIn(req *SignInRequest) dto.BaseResponse[Token]
	SignUp(req *SignUpRequest) dto.BaseResponse[Token]
	SignOut(userId uuid.UUID) error
	RefreshToken(req *RefreshTokenRequest) dto.BaseResponse[Token]
}

func NewService(
	userUseCase user.UseCase,
	authRepository Repository,
	jwtClient *jwt.JwtClient,
) UseCase {
	return &authService{
		userUseCase:    userUseCase,
		authRepository: authRepository,
		jwtClient:      jwtClient,
	}
}

func (a authService) SignIn(req *SignInRequest) dto.BaseResponse[Token] {
	loginDto := &user.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	}

	// 로그인 처리
	resp := a.userUseCase.VerifyUser(loginDto)
	if !resp.IsSuccess() {
		return dto.Fail[Token](response.ResultCode(resp.GetCode()), *resp.GetError())
	}
	userId := resp.GetData().ID

	// 토큰 생성
	tokenRequest := &jwt.TokenRequest{
		ID:          userId.String(),
		Credentials: []string{},
	}
	jwtToken, err := a.jwtClient.Generator.GenerateNewTokens(tokenRequest)
	if err != nil {
		return dto.Fail[Token](response.CODE_FAIL_CREATE_TOKEN, err)
	}

	// 토큰 저장
	newToken := &Token{
		AccessToken:         jwtToken.AccessToken,
		AccessTokenExpires:  jwtToken.AccessTokenExpires,
		RefreshToken:        jwtToken.RefreshToken,
		RefreshTokenExpires: jwtToken.RefreshTokenExpires,
	}
	a.authRepository.DeleteToken(userId)
	savedToken, err := a.authRepository.SaveToken(userId, newToken)
	if err != nil {
		return dto.Fail[Token](response.CODE_FAIL_SAVE_TOKEN, err)
	}

	return dto.Ok[Token](response.CODE_SUCCESS, savedToken)
}

func (a authService) SignUp(req *SignUpRequest) dto.BaseResponse[Token] {
	registerDto := &user.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	resp := a.userUseCase.RegisterUser(registerDto)
	if !resp.IsSuccess() {
		return dto.Fail[Token](response.ResultCode(resp.GetCode()), *resp.GetError())
	}

	foundUser := resp.GetData()
	tokenRequest := &jwt.TokenRequest{
		ID:          foundUser.ID.String(),
		Credentials: []string{},
	}
	jwtToken, err := a.jwtClient.Generator.GenerateNewTokens(tokenRequest)
	if err != nil {
		return dto.Fail[Token](response.CODE_FAIL_CREATE_TOKEN, err)
	}

	token := &Token{
		AccessToken:  jwtToken.AccessToken,
		RefreshToken: jwtToken.RefreshToken,
	}
	return dto.Ok[Token](response.CODE_SUCCESS, token)
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

func (a authService) RefreshToken(req *RefreshTokenRequest) dto.BaseResponse[Token] {
	userId := req.AccessToken.UserID
	// 토큰 생성
	tokenRequest := &jwt.TokenRequest{
		ID:          userId.String(),
		Credentials: []string{},
	}
	jwtToken, err := a.jwtClient.Generator.GenerateNewTokens(tokenRequest)
	if err != nil {
		return dto.Fail[Token](response.CODE_FAIL_CREATE_TOKEN, err)
	}

	// 토큰 저장
	newToken := &Token{
		AccessToken:         jwtToken.AccessToken,
		AccessTokenExpires:  jwtToken.AccessTokenExpires,
		RefreshToken:        jwtToken.RefreshToken,
		RefreshTokenExpires: jwtToken.RefreshTokenExpires,
	}
	a.authRepository.DeleteToken(userId)
	savedToken, err := a.authRepository.SaveToken(userId, newToken)
	if err != nil {
		return dto.Fail[Token](response.CODE_FAIL_SAVE_TOKEN, err)
	}

	return dto.Ok[Token](response.CODE_SUCCESS, savedToken)
}

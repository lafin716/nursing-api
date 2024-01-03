package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"log"
	"nursing_api/app/models"
	"nursing_api/app/utils"
	"nursing_api/pkg/auth"
	"time"
)

// 회원가입
func UserSignUp(c *fiber.Ctx) error {
	signUp := &models.SignUp{}

	// 파라미터 파싱
	if err := parseParams(c, signUp); err != nil {
		return nil
	}

	// 유효성 검사
	if err := validate(c, signUp); err != nil {
		return nil
	}

	// 유저 객체 생성
	user := &models.User{
		ID:           uuid.New(),
		Name:         signUp.Name,
		Email:        signUp.Email,
		PasswordHash: utils.GeneratePassword(signUp.Password),
		UserStatus:   0,
		SignInType:   "email",
		FailCount:    0,
		CreatedAt:    time.Now(),
	}

	// 비밀번호는 가림
	user.PasswordHash = ""

	return utils.ResponseEntity{
		Code: utils.CODE_SUCCESS,
		Data: user,
	}.ResponseOk(c)
}

// 로그인
func UserSignIn(c *fiber.Ctx) error {
	signIn := &models.SignIn{}

	// 파라미터 파싱
	if err := parseParams(c, signIn); err != nil {
		return nil
	}

	// 유효성 검사
	if err := validate(c, signIn); err != nil {
		return nil
	}

	// TODO DB에서 유저정보 조회, 현재는 임시 유저객체로 조회 mock
	user := &models.User{
		ID:           uuid.New(),
		Name:         "임시유저",
		Email:        signIn.Email,
		PasswordHash: utils.GeneratePassword(signIn.Password),
		FailCount:    0,
		LastLoggedIn: time.Now(),
		CreatedAt:    time.Now(),
	}

	// 비밀번호 검증
	comparePassword := utils.ComparePassword(user.PasswordHash, signIn.Password)
	if !comparePassword {
		return utils.ResponseEntity{
			Code: utils.CODE_INVALID_SIGN_IN,
		}.ResponseBadRequest(c)
	}

	// 토큰 생성
	tokens, err := auth.GenerateNewTokens(user.ID.String(), []string{})
	if err != nil {
		return utils.ResponseEntity{
			Code:          utils.CODE_ERROR_WITH_MSG,
			MessageParams: []string{"토큰 생성 실패"},
		}.ResponseError(c)
	}

	return utils.ResponseEntity{
		Code: utils.CODE_SUCCESS,
		Data: tokens,
	}.ResponseOk(c)
}

// 로그아웃
func UserSignOut(c *fiber.Ctx) error {
	claims, err := parseJwt(c)
	if err != nil {
		return nil
	}

	userId := claims.UserID.String()
	log.Printf("로그아웃 : %s", userId)

	return utils.ResponseEntity{Code: utils.CODE_SUCCESS}.ResponseOk(c)
}

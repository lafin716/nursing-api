package controllers

import (
	"github.com/google/uuid"
	"log"
	"nursing_api/app/container"
	"nursing_api/app/models"
	"nursing_api/app/repository"
	"nursing_api/app/usecase"
	"nursing_api/app/utils"
	"nursing_api/pkg/auth"
	"nursing_api/pkg/database"
	"time"
)

// 회원가입
func UserSignUp(container *container.FiberContainer) error {
	c := container.Ctx
	db := container.DBClient
	service := getService(db)
	signUp := &models.SignUp{}

	// 파라미터 파싱
	if err := parseParams(c, signUp); err != nil {
		return nil
	}

	// 유효성 검사
	if err := validate(c, signUp); err != nil {
		return nil
	}

	// 기존 이메일 가입 여부 확인
	emailCount := service.CountUserByEmail(signUp.Email)
	if emailCount < 0 {
		return utils.ResponseEntity{
			Code:          utils.CODE_ERROR_WITH_MSG,
			MessageParams: []string{"이메일 정보를 조회할 수 없습니다."},
		}.ResponseError(c)
	}

	// 이메일 카운트가 0 이상인 경우 이미 가입된 이메일 응답
	if emailCount > 0 {
		return utils.ResponseEntity{
			Code: utils.CODE_DUPLICATE_EMAIL,
		}.ResponseOk(c)
	}

	// 유저 객체 생성
	newUser := &models.User{
		ID:           uuid.New(),
		Name:         signUp.Name,
		Email:        signUp.Email,
		PasswordHash: utils.GeneratePassword(signUp.Password),
		UserType:     "EMAIL",
		UserStatus:   models.USER_ACTIVE,
		FailCount:    0,
		CreatedAt:    time.Now(),
	}

	savedUser := service.SaveUser(newUser)
	savedUser.PasswordHash = ""

	return utils.ResponseEntity{
		Code: utils.CODE_SUCCESS,
		Data: savedUser,
	}.ResponseOk(c)
}

// 로그인
func UserSignIn(container *container.FiberContainer) error {
	c := container.Ctx
	service := getService(container.DBClient)
	signIn := &models.SignIn{}

	// 파라미터 파싱
	if err := parseParams(c, signIn); err != nil {
		return nil
	}

	// 유효성 검사
	if err := validate(c, signIn); err != nil {
		return nil
	}

	foundUser := service.GetUserByEmail(signIn.Email)
	if foundUser == nil {
		return utils.ResponseEntity{
			Code: utils.CODE_USER_NOTFOUND,
		}.ResponseOk(c)
	}

	// 비밀번호 검증
	comparePassword := utils.ComparePassword(foundUser.PasswordHash, signIn.Password)
	if !comparePassword {
		return utils.ResponseEntity{
			Code: utils.CODE_INVALID_SIGN_IN,
		}.ResponseBadRequest(c)
	}

	// 토큰 생성
	tokenRequest := &auth.TokenRequest{
		ID:          foundUser.ID.String(),
		Credentials: []string{},
	}
	tokens, err := container.JwtClient.Generator.GenerateNewTokens(tokenRequest)
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
func UserSignOut(container *container.FiberContainer) error {
	c := container.Ctx
	claims, err := parseJwt(container)
	if err != nil {
		return nil
	}

	userId := claims.UserID.String()
	log.Printf("로그아웃 : %s", userId)

	return utils.ResponseEntity{Code: utils.CODE_SUCCESS}.ResponseOk(c)
}

func getService(db *database.DatabaseClient) usecase.UserUseCase {
	userRepository := repository.NewUserRepository(db.Client, db.Ctx)
	gateway := usecase.NewUserGateway(userRepository)
	return usecase.NewUserService(gateway)
}

package api

import (
	"fmt"
	"github.com/google/wire"
	"log"
	"nursing_api/pkg/validates"
)

var Set = wire.NewSet(
	NewMainHttpApi,
	NewUserHttpApi,
	NewAuthHttpApi,
	NewMedicineHttpApi,
	NewErrorHttpApi,
	NewPrescriptionApi,
	NewTakeHistoryHttpApi,
	NewPlanHttpApi,
)

func validateParameter(data interface{}) map[string]string {
	ruleGroup := validates.NewRuleGroup(
		validates.NewRule("password", passwordValidator),
		validates.NewRule("keyword", keywordValidator),
	)

	validator, err := validates.NewValidatorWithCustomValidators(ruleGroup)
	if err != nil {
		log.Printf("유효성 검사 모듈 생성 중 오류가 발생하였습니다: %v", err)
		return nil
	}

	errs := validator.Struct(data)
	if errs != nil {
		validateErrors := validates.ValidateErrors(errs)
		return resolveErrorMessage(validateErrors)
	}

	return nil
}

func resolveErrorMessage(errs map[string]validates.ValidateError) map[string]string {
	message := map[string]string{}
	for k, v := range errs {
		message[k] = resolveMessage(v)
	}

	return message
}

func resolveMessage(err validates.ValidateError) string {
	message := "유효하지 않은 파라미터입니다."
	switch err.Tag {
	case "required":
		return fmt.Sprintf("%s 필드는 필수입니다.", err.Field)
	case "min":
		return fmt.Sprintf("%s 필드는 최소 %s 자 이상 입력하세요.", err.Field, err.Param)
	case "max":
		return fmt.Sprintf("%s 필드는 최대 %s 자 이하로 입력하세요.", err.Field, err.Param)
	case "password":
		return fmt.Sprintf("비밀번호는 8~20자 이내로 입력하세요.")
	case "keyword":
		return fmt.Sprintf("검색어는 최소 2자 이상 입력하세요.")
	}

	return message
}

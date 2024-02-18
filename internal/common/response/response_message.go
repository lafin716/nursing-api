package response

import "fmt"

const (
	CODE_SUCCESS = 200

	CODE_UNAUTHORIZED      = 401
	CODE_INVALID_JWT       = 402
	CODE_INVALID_SIGN_IN   = 403
	CODE_NOTFOUND          = 404
	CODE_NO_DATA           = 405
	CODE_INVALID_PARAM     = 422
	CODE_INVALID_PARAM_MSG = 423

	CODE_ERROR                        = 500
	CODE_ERROR_WITH_MSG               = 501
	CODE_FAIL_SIGN_UP                 = 510
	CODE_FAIL_SIGN_OUT                = 511
	CODE_FAIL_LEAVE                   = 512
	CODE_FAIL_ADD_PRESCRIPTION        = 520
	CODE_FAIL_UPDATE_PRESCRIPTION     = 521
	CODE_FAIL_DELETE_PRESCRIPTION     = 522
	CODE_FAIL_ADDITEM_PRESCRIPTION    = 523
	CODE_FAIL_UPDATEITEM_PRESCRIPTION = 524
	CODE_FAIL_DELETEITEM_PRESCRIPTION = 525

	CODE_FAIL_DELETE_PLAN = 550

	CODE_DUPLICATE_EMAIL = 600
	CODE_USER_NOTFOUND   = 601
)

var codeMessages = map[int]string{
	CODE_SUCCESS:                      "정상 요청 되었습니다.",
	CODE_NOTFOUND:                     "존재하지 않는 엔드포인트입니다.",
	CODE_UNAUTHORIZED:                 "권한이 없습니다.",
	CODE_INVALID_JWT:                  "인증 토큰이 유효하지 않습니다.",
	CODE_INVALID_SIGN_IN:              "로그인 정보가 일치하지 않습니다.",
	CODE_FAIL_SIGN_UP:                 "회원가입이 정상적으로 처리되지 않았습니다.",
	CODE_FAIL_SIGN_OUT:                "로그아웃 중 오류가 발생하였습니다.",
	CODE_INVALID_PARAM:                "유효하지 않은 파라미터입니다.",
	CODE_INVALID_PARAM_MSG:            "유효하지 않은 파라미터입니다 :: %s",
	CODE_NO_DATA:                      "데이터가 없습니다.",
	CODE_ERROR:                        "오류가 발생하였습니다",
	CODE_ERROR_WITH_MSG:               "오류가 발생하였습니다 :: %s",
	CODE_FAIL_LEAVE:                   "회원탈퇴 중 오류가 발생하였습니다.",
	CODE_FAIL_ADD_PRESCRIPTION:        "처방전 등록 중 오류가 발생하였습니다.",
	CODE_FAIL_UPDATE_PRESCRIPTION:     "처방전 업데이트 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETE_PRESCRIPTION:     "처방전 삭제 중 오류가 발생하였습니다.",
	CODE_FAIL_ADDITEM_PRESCRIPTION:    "처방전 의약품 등록 중 오류가 발생하였습니다.",
	CODE_FAIL_UPDATEITEM_PRESCRIPTION: "처방전 의약품 업데이트 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETEITEM_PRESCRIPTION: "처방전 의약품 삭제 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETE_PLAN:             "복용계획 삭제 중 오류가 발생하였습니다.",
	CODE_DUPLICATE_EMAIL:              "이미 가입된 이메일입니다.",
	CODE_USER_NOTFOUND:                "회원정보를 찾을 수 없습니다.",
}

func GetMessage(code int, param ...string) string {
	if codeMessages[code] == "" {
		return ""
	}

	if param == nil {
		return codeMessages[code]
	}

	return fmt.Sprintf(codeMessages[code], param)
}

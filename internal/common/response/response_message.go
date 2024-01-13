package response

import "fmt"

const (
	CODE_SUCCESS = 200

	CODE_UNAUTHORIZED    = 401
	CODE_INVALID_JWT     = 402
	CODE_INVALID_SIGN_IN = 403
	CODE_NOTFOUND        = 404
	CODE_NO_DATA         = 405
	CODE_INVALID_PARAM   = 422

	CODE_ERROR          = 500
	CODE_ERROR_WITH_MSG = 501
	CODE_FAIL_SIGN_UP   = 510
	CODE_FAIL_SIGN_OUT  = 511

	CODE_DUPLICATE_EMAIL = 600
	CODE_USER_NOTFOUND   = 601
)

var codeMessages = map[int]string{
	CODE_SUCCESS:         "정상 요청 되었습니다.",
	CODE_NOTFOUND:        "존재하지 않는 엔드포인트입니다.",
	CODE_UNAUTHORIZED:    "권한이 없습니다.",
	CODE_INVALID_JWT:     "인증 토큰이 유효하지 않습니다.",
	CODE_INVALID_SIGN_IN: "로그인 정보가 일치하지 않습니다.",
	CODE_FAIL_SIGN_UP:    "회원가입이 정상적으로 처리되지 않았습니다.",
	CODE_FAIL_SIGN_OUT:   "로그아웃 중 오류가 발생하였습니다.",
	CODE_INVALID_PARAM:   "유효하지 않은 파라미터입니다.",
	CODE_NO_DATA:         "데이터가 없습니다.",
	CODE_ERROR:           "오류가 발생하였습니다",
	CODE_ERROR_WITH_MSG:  "오류가 발생하였습니다 :: %s",
	CODE_DUPLICATE_EMAIL: "이미 가입된 이메일입니다.",
	CODE_USER_NOTFOUND:   "회원정보를 찾을 수 없습니다.",
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

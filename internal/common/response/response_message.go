package response

import "fmt"

type ResultCode int

const (
	// 공통
	CODE_SUCCESS           = ResultCode(200)
	CODE_UNAUTHORIZED      = ResultCode(401)
	CODE_NOTFOUND          = ResultCode(404)
	CODE_NO_DATA           = ResultCode(405)
	CODE_INVALID_PARAM     = ResultCode(422)
	CODE_INVALID_PARAM_MSG = ResultCode(423)
	CODE_ERROR             = ResultCode(500)
	CODE_ERROR_WITH_MSG    = ResultCode(501)
	CODE_ERROR_TRANSACTION = ResultCode(505)

	// 유효성 관련
	CODE_NOT_AVAILABLE_DATE = ResultCode(600)

	// 인증
	CODE_INVALID_JWT         = ResultCode(1000)
	CODE_EXPIRED_ACCESSTOKEN = ResultCode(1001)
	CODE_FAIL_SIGN_UP        = ResultCode(1020)
	CODE_FAIL_SIGN_OUT       = ResultCode(1030)
	CODE_FAIL_LEAVE          = ResultCode(1040)
	CODE_FAIL_REFRESH_TOKEN  = ResultCode(1050)
	CODE_FAIL_CREATE_TOKEN   = ResultCode(1060)
	CODE_FAIL_SAVE_TOKEN     = ResultCode(1070)

	// 회원
	CODE_DUPLICATE_EMAIL                  = ResultCode(1100)
	CODE_USER_NOTFOUND                    = ResultCode(1110)
	CODE_AVAILABLE_EMAIL                  = ResultCode(1120)
	CODE_OCCUR_ERROR_DURING_CHECK_EMAIL   = ResultCode(1130)
	CODE_OCCUR_ERROR_DURING_GET_USER      = ResultCode(1131)
	CODE_OCCUR_ERROR_DURING_SAVE_USER     = ResultCode(1132)
	CODE_NOT_MATCHING_SIGN_IN_INFORMATION = ResultCode(1140)

	// 의약품
	CODE_NOT_AVAILABLE_MEDICINE_KEY = ResultCode(2000)
	CODE_NOT_FOUND_MEDICINE         = ResultCode(2010)
	CODE_SEARCH_MEDICINE_ZERO       = ResultCode(2020)
	CODE_FAIL_SAVE_MEDICINE_DATA    = ResultCode(2030)
	CODE_FAIL_SAVE_MEDICINE_ZERO    = ResultCode(2040)

	// 처방전
	CODE_FAIL_ADD_PRESCRIPTION        = ResultCode(4000)
	CODE_FAIL_UPDATE_PRESCRIPTION     = ResultCode(4010)
	CODE_FAIL_DELETE_PRESCRIPTION     = ResultCode(4020)
	CODE_FAIL_ADDITEM_PRESCRIPTION    = ResultCode(4030)
	CODE_FAIL_UPDATEITEM_PRESCRIPTION = ResultCode(4040)
	CODE_FAIL_DELETEITEM_PRESCRIPTION = ResultCode(4050)

	// 복용계획
	CODE_NOT_FOUND_PLAN                   = ResultCode(5000)
	CODE_FAIL_ADD_PLAN                    = ResultCode(5001)
	CODE_NOT_FOUND_PLAN_ITEM              = ResultCode(5005)
	CODE_NOT_FOUND_PLAN_TIMEZONE          = ResultCode(5010)
	CODE_FAIL_CONNECT_TIMEZONE            = ResultCode(5020)
	CODE_FAIL_ADD_ITEM_PLAN               = ResultCode(5030)
	CODE_FAIL_DELETE_PLAN                 = ResultCode(5040)
	CODE_FAIL_DELETE_PLAN_TIMEZONE        = ResultCode(5050)
	CODE_FAIL_DELETE_ITEM_PLAN            = ResultCode(5060)
	CODE_NO_RESULT_PLAN_BY_DATE           = ResultCode(5070)
	CODE_FAIL_TAKE_PLAN                   = ResultCode(5080)
	CODE_TOO_MUCH_TAKE_AMOUNT             = ResultCode(5081)
	CODE_ZERO_TAKE_AMOUNT                 = ResultCode(5082)
	CODE_NOT_VALID_AMOUNT                 = ResultCode(5083)
	CODE_REMAIN_AMOUNT_CANNOT_BE_NEGATIVE = ResultCode(5084)
	CODE_FAIL_SAVE_PLAN_LOG               = ResultCode(5090)
	CODE_CANNOT_UPDATE_MEMO_BEFORE_TAKING = ResultCode(5100)
	CODE_FAIL_UPDATE_MEMO                 = ResultCode(5110)
	CODE_NO_RESULT_PLAN_BY_MONTH          = ResultCode(5120)
	CODE_FAIL_UPDATE_RECOVER_TAKE_AMOUNT  = ResultCode(5130)

	// 복용내역
	CODE_FAIL_GET_LIST_TAKE_HISTORY      = ResultCode(6000)
	CODE_FAIL_GET_ITEM_LIST_TAKE_HISTORY = ResultCode(6001)
	CODE_FAIL_ADD_TAKE_HISTORY_ITEM      = ResultCode(6002)
	CODE_FAIL_UPDATE_TAKEHISTORY         = ResultCode(6003)

	// 복용시간대
	CODE_FAIL_GET_LIST_TAKE_TIMEZONE          = ResultCode(7000)
	CODE_NOT_FOUND_TIMEZONE                   = ResultCode(7001)
	CODE_DUPLICATE_TIMEZONE                   = ResultCode(7002)
	CODE_ALREADY_EXIST_TIMEZONE_NAME          = ResultCode(7003)
	CODE_ALREADY_EXIST_TIMEZONE_HOUR_MINUTE   = ResultCode(7004)
	CODE_FAIL_DURING_CREATE_TIMEZONE          = ResultCode(7005)
	CODE_FAIL_DURING_UPDATE_TIMEZONE          = ResultCode(7006)
	CODE_FAIL_DURING_DELETE_TIMEZONE          = ResultCode(7007)
	CODE_FAIL_DURING_CHECK_DELETABLE_TIMEZONE = ResultCode(7008)
	CODE_CANNOT_DELETE_TIMEZONE               = ResultCode(7009)
)

var codeMessages = map[ResultCode]string{
	// 공통
	CODE_SUCCESS:           "정상 요청 되었습니다.",
	CODE_NOTFOUND:          "존재하지 않는 엔드포인트입니다.",
	CODE_UNAUTHORIZED:      "권한이 없습니다.",
	CODE_INVALID_PARAM:     "유효하지 않은 파라미터입니다.",
	CODE_INVALID_PARAM_MSG: "유효하지 않은 파라미터입니다 :: %s",
	CODE_NO_DATA:           "데이터가 없습니다.",
	CODE_ERROR:             "오류가 발생하였습니다",
	CODE_ERROR_WITH_MSG:    "오류가 발생하였습니다 :: %s",
	CODE_ERROR_TRANSACTION: "트랜잭션 오류가 발생하였습니다.",

	// 유효성 관련
	CODE_NOT_AVAILABLE_DATE: "날짜형식이 유효하지 않습니다.",

	// 인증
	CODE_INVALID_JWT:         "인증 토큰이 유효하지 않습니다.",
	CODE_EXPIRED_ACCESSTOKEN: "토큰이 만료되었습니다.",
	CODE_FAIL_SIGN_UP:        "회원가입이 정상적으로 처리되지 않았습니다.",
	CODE_FAIL_SIGN_OUT:       "로그아웃 중 오류가 발생하였습니다.",
	CODE_FAIL_REFRESH_TOKEN:  "토큰 갱신 중 오류가 발생하였습니다.",
	CODE_FAIL_LEAVE:          "회원탈퇴 중 오류가 발생하였습니다.",
	CODE_FAIL_CREATE_TOKEN:   "토큰 생성에 실패하였습니다.",
	CODE_FAIL_SAVE_TOKEN:     "토큰 저장 중 오류가 발생하였습니다.",

	// 회원
	CODE_DUPLICATE_EMAIL:                  "이미 가입된 이메일입니다.",
	CODE_USER_NOTFOUND:                    "회원정보를 찾을 수 없습니다.",
	CODE_AVAILABLE_EMAIL:                  "사용 가능한 이메일입니다.",
	CODE_OCCUR_ERROR_DURING_CHECK_EMAIL:   "이메일 중복 확인 중 오류가 발생하였습니다.",
	CODE_NOT_MATCHING_SIGN_IN_INFORMATION: "로그인 정보가 일치하지 않습니다.",
	CODE_OCCUR_ERROR_DURING_GET_USER:      "회원 정보 조회 중 오류가 발생하였습니다.",
	CODE_OCCUR_ERROR_DURING_SAVE_USER:     "회원 정보 저장 중 오류가 발생하였습니다.",

	// 의약품
	CODE_NOT_AVAILABLE_MEDICINE_KEY: "검색어가 유효하지 않습니다.",
	CODE_NOT_FOUND_MEDICINE:         "의약품 정보를 찾을 수 없습니다.",
	CODE_SEARCH_MEDICINE_ZERO:       "의약품 검색 결과: 0건",
	CODE_FAIL_SAVE_MEDICINE_DATA:    "의약품 정보 저장 중 오류가 발생하였습니다.",
	CODE_FAIL_SAVE_MEDICINE_ZERO:    "의약품 정보 저장: 0건",

	// 복용계획
	CODE_NOT_FOUND_PLAN:                   "복용계획을 찾을 수 없습니다.",
	CODE_FAIL_ADD_PLAN:                    "복용계획 등록 중 오류가 발생하였습니다.",
	CODE_NOT_FOUND_PLAN_ITEM:              "복용계획 항목을 찾을 수 없습니다.",
	CODE_NOT_FOUND_PLAN_TIMEZONE:          "복용계획 시간대를 찾을 수 없습니다.",
	CODE_FAIL_CONNECT_TIMEZONE:            "복용계획 시간대 연결 중 오류가 발생하였습니다.",
	CODE_FAIL_ADD_ITEM_PLAN:               "복용계획 항목 등록 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETE_PLAN:                 "복용계획 삭제 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETE_PLAN_TIMEZONE:        "복용계획 시간대 삭제 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETE_ITEM_PLAN:            "복용계획 항목 삭제 중 오류가 발생하였습니다.",
	CODE_NO_RESULT_PLAN_BY_DATE:           "해당날짜에 복용계획이 없습니다.",
	CODE_FAIL_TAKE_PLAN:                   "복용처리 중 오류가 발생하였습니다.",
	CODE_TOO_MUCH_TAKE_AMOUNT:             "복용량이 초과되었습니다.",
	CODE_FAIL_SAVE_PLAN_LOG:               "복용이력 저장 중 오류가 발생하였습니다.",
	CODE_CANNOT_UPDATE_MEMO_BEFORE_TAKING: "복용처리 전에는 메모를 수정할 수 없습니다.",
	CODE_FAIL_UPDATE_MEMO:                 "메모 수정 중 오류가 발생하였습니다.",
	CODE_NO_RESULT_PLAN_BY_MONTH:          "해당월에 복용계획이 없습니다.",
	CODE_FAIL_UPDATE_RECOVER_TAKE_AMOUNT:  "복용량 복구 중 오류가 발생하였습니다.",
	CODE_ZERO_TAKE_AMOUNT:                 "복용량은 0이 될 수 없습니다",
	CODE_NOT_VALID_AMOUNT:                 "복용량과 남은량의 합은 총량과 같아야합니다.",
	CODE_REMAIN_AMOUNT_CANNOT_BE_NEGATIVE: "남은량은 음수가 될 수 없습니다.",

	// 처방전
	CODE_FAIL_ADD_PRESCRIPTION:        "처방전 등록 중 오류가 발생하였습니다.",
	CODE_FAIL_UPDATE_PRESCRIPTION:     "처방전 업데이트 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETE_PRESCRIPTION:     "처방전 삭제 중 오류가 발생하였습니다.",
	CODE_FAIL_ADDITEM_PRESCRIPTION:    "처방전 의약품 등록 중 오류가 발생하였습니다.",
	CODE_FAIL_UPDATEITEM_PRESCRIPTION: "처방전 의약품 업데이트 중 오류가 발생하였습니다.",
	CODE_FAIL_DELETEITEM_PRESCRIPTION: "처방전 의약품 삭제 중 오류가 발생하였습니다.",

	// 복용내역
	CODE_FAIL_ADD_TAKE_HISTORY_ITEM:      "복용내역 등록 중 오류가 발생하였습니다.",
	CODE_FAIL_GET_LIST_TAKE_HISTORY:      "복용내역 목록 조회 중 오류가 발생하였습니다.",
	CODE_FAIL_GET_ITEM_LIST_TAKE_HISTORY: "복용내역 아이템 목록 조회 중 오류가 발생하였습니다.",
	CODE_FAIL_UPDATE_TAKEHISTORY:         "복용내역 수정 중 오류가 발생하였습니다.",

	// 복용시간대
	CODE_FAIL_GET_LIST_TAKE_TIMEZONE:          "복용시간대 목록 조회 중 오류가 발생하였습니다.",
	CODE_NOT_FOUND_TIMEZONE:                   "시간대를 찾을 수 없습니다.",
	CODE_DUPLICATE_TIMEZONE:                   "중복되는 시간대가 존재합니다.",
	CODE_ALREADY_EXIST_TIMEZONE_NAME:          "이미 존재하는 이름입니다.",
	CODE_ALREADY_EXIST_TIMEZONE_HOUR_MINUTE:   "이미 등록된 시간입니다.",
	CODE_FAIL_DURING_CREATE_TIMEZONE:          "시간대 생성 중 오류가 발생하였습니다.",
	CODE_FAIL_DURING_UPDATE_TIMEZONE:          "시간대 수정 중 오류가 발생하였습니다.",
	CODE_FAIL_DURING_DELETE_TIMEZONE:          "시간대 삭제 중 오류가 발생하였습니다.",
	CODE_FAIL_DURING_CHECK_DELETABLE_TIMEZONE: "삭제 가능한 시간대인지 확인 중 오류가 발생하였습니다.",
	CODE_CANNOT_DELETE_TIMEZONE:               "연결 된 처방전이 복용기간 중으로 삭제할 수 없는 시간대입니다.",
}

func GetMessage(code ResultCode, param ...string) string {
	if codeMessages[code] == "" {
		return ""
	}
	if param == nil {
		return codeMessages[code]
	}

	return fmt.Sprintf(codeMessages[code], param)
}

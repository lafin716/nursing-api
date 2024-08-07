package response

import (
	"fmt"
	"testing"
	"time"
)

func Test_Time(t *testing.T) {
	now := time.Now()
	nowUtc := now.UTC()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(nowUtc.Format("2006-01-02 15:04:05"))

}

func TestResponseMessages(t *testing.T) {

	print(CODE_SUCCESS, ":")
	println(codeMessages[CODE_SUCCESS])
	print(CODE_UNAUTHORIZED, ":")
	println(codeMessages[CODE_UNAUTHORIZED])
	print(CODE_NOTFOUND, ":")
	println(codeMessages[CODE_NOTFOUND])
	print(CODE_NO_DATA, ":")
	println(codeMessages[CODE_NO_DATA])
	print(CODE_INVALID_PARAM, ":")
	println(codeMessages[CODE_INVALID_PARAM])
	print(CODE_INVALID_PARAM_MSG, ":")
	println(codeMessages[CODE_INVALID_PARAM_MSG])
	print(CODE_ERROR, ":")
	println(codeMessages[CODE_ERROR])
	print(CODE_ERROR_WITH_MSG, ":")
	println(codeMessages[CODE_ERROR_WITH_MSG])
	print(CODE_ERROR_TRANSACTION, ":")
	println(codeMessages[CODE_ERROR_TRANSACTION])

	// 유효성 관련
	print(CODE_NOT_AVAILABLE_DATE, ":")
	println(codeMessages[CODE_NOT_AVAILABLE_DATE])

	// 인증
	print(CODE_INVALID_JWT, ":")
	println(codeMessages[CODE_INVALID_JWT])
	print(CODE_EXPIRED_ACCESSTOKEN, ":")
	println(codeMessages[CODE_EXPIRED_ACCESSTOKEN])
	print(CODE_FAIL_SIGN_UP, ":")
	println(codeMessages[CODE_FAIL_SIGN_UP])
	print(CODE_FAIL_SIGN_OUT, ":")
	println(codeMessages[CODE_FAIL_SIGN_OUT])
	print(CODE_FAIL_LEAVE, ":")
	println(codeMessages[CODE_FAIL_LEAVE])
	print(CODE_FAIL_REFRESH_TOKEN, ":")
	println(codeMessages[CODE_FAIL_REFRESH_TOKEN])
	print(CODE_FAIL_CREATE_TOKEN, ":")
	println(codeMessages[CODE_FAIL_CREATE_TOKEN])
	print(CODE_FAIL_SAVE_TOKEN, ":")
	println(codeMessages[CODE_FAIL_SAVE_TOKEN])

	// 회원
	print(CODE_DUPLICATE_EMAIL, ":")
	println(codeMessages[CODE_DUPLICATE_EMAIL])
	print(CODE_USER_NOTFOUND, ":")
	println(codeMessages[CODE_USER_NOTFOUND])
	print(CODE_AVAILABLE_EMAIL, ":")
	println(codeMessages[CODE_AVAILABLE_EMAIL])
	print(CODE_OCCUR_ERROR_DURING_CHECK_EMAIL, ":")
	println(codeMessages[CODE_OCCUR_ERROR_DURING_CHECK_EMAIL])
	print(CODE_OCCUR_ERROR_DURING_GET_USER, ":")
	println(codeMessages[CODE_OCCUR_ERROR_DURING_GET_USER])
	print(CODE_OCCUR_ERROR_DURING_SAVE_USER, ":")
	println(codeMessages[CODE_OCCUR_ERROR_DURING_SAVE_USER])
	print(CODE_NOT_MATCHING_SIGN_IN_INFORMATION, ":")
	println(codeMessages[CODE_NOT_MATCHING_SIGN_IN_INFORMATION])

	// 의약품
	print(CODE_NOT_AVAILABLE_MEDICINE_KEY, ":")
	println(codeMessages[CODE_NOT_AVAILABLE_MEDICINE_KEY])
	print(CODE_NOT_FOUND_MEDICINE, ":")
	println(codeMessages[CODE_NOT_FOUND_MEDICINE])
	print(CODE_SEARCH_MEDICINE_ZERO, ":")
	println(codeMessages[CODE_SEARCH_MEDICINE_ZERO])
	print(CODE_FAIL_SAVE_MEDICINE_DATA, ":")
	println(codeMessages[CODE_FAIL_SAVE_MEDICINE_DATA])
	print(CODE_FAIL_SAVE_MEDICINE_ZERO, ":")
	println(codeMessages[CODE_FAIL_SAVE_MEDICINE_ZERO])

	// 처방전
	print(CODE_FAIL_ADD_PRESCRIPTION, ":")
	println(codeMessages[CODE_FAIL_ADD_PRESCRIPTION])
	print(CODE_FAIL_UPDATE_PRESCRIPTION, ":")
	println(codeMessages[CODE_FAIL_UPDATE_PRESCRIPTION])
	print(CODE_FAIL_DELETE_PRESCRIPTION, ":")
	println(codeMessages[CODE_FAIL_DELETE_PRESCRIPTION])
	print(CODE_FAIL_ADDITEM_PRESCRIPTION, ":")
	println(codeMessages[CODE_FAIL_ADDITEM_PRESCRIPTION])
	print(CODE_FAIL_UPDATEITEM_PRESCRIPTION, ":")
	println(codeMessages[CODE_FAIL_UPDATEITEM_PRESCRIPTION])
	print(CODE_FAIL_DELETEITEM_PRESCRIPTION, ":")
	println(codeMessages[CODE_FAIL_DELETEITEM_PRESCRIPTION])

	// 복용계획
	print(CODE_NOT_FOUND_PLAN, ":")
	println(codeMessages[CODE_NOT_FOUND_PLAN])
	print(CODE_NOT_FOUND_PLAN_ITEM, ":")
	println(codeMessages[CODE_NOT_FOUND_PLAN_ITEM])
	print(CODE_NOT_FOUND_PLAN_TIMEZONE, ":")
	println(codeMessages[CODE_NOT_FOUND_PLAN_TIMEZONE])
	print(CODE_FAIL_CONNECT_TIMEZONE, ":")
	println(codeMessages[CODE_FAIL_CONNECT_TIMEZONE])
	print(CODE_FAIL_ADD_ITEM_PLAN, ":")
	println(codeMessages[CODE_FAIL_ADD_ITEM_PLAN])
	print(CODE_FAIL_DELETE_PLAN, ":")
	println(codeMessages[CODE_FAIL_DELETE_PLAN])
	print(CODE_FAIL_DELETE_PLAN_TIMEZONE, ":")
	println(codeMessages[CODE_FAIL_DELETE_PLAN_TIMEZONE])
	print(CODE_FAIL_DELETE_ITEM_PLAN, ":")
	println(codeMessages[CODE_FAIL_DELETE_ITEM_PLAN])
	print(CODE_NO_RESULT_PLAN_BY_DATE, ":")
	println(codeMessages[CODE_NO_RESULT_PLAN_BY_DATE])
	print(CODE_FAIL_TAKE_PLAN, ":")
	println(codeMessages[CODE_FAIL_TAKE_PLAN])
	print(CODE_FAIL_SAVE_PLAN_LOG, ":")
	println(codeMessages[CODE_FAIL_SAVE_PLAN_LOG])
	print(CODE_CANNOT_UPDATE_MEMO_BEFORE_TAKING, ":")
	println(codeMessages[CODE_CANNOT_UPDATE_MEMO_BEFORE_TAKING])
	print(CODE_FAIL_UPDATE_MEMO, ":")
	println(codeMessages[CODE_FAIL_UPDATE_MEMO])

	// 복용내역
	print(CODE_FAIL_GET_LIST_TAKE_HISTORY, ":")
	println(codeMessages[CODE_FAIL_GET_LIST_TAKE_HISTORY])
	print(CODE_FAIL_GET_ITEM_LIST_TAKE_HISTORY, ":")
	println(codeMessages[CODE_FAIL_GET_ITEM_LIST_TAKE_HISTORY])

	// 복용시간대
	print(CODE_FAIL_GET_LIST_TAKE_TIMEZONE, ":")
	println(codeMessages[CODE_FAIL_GET_LIST_TAKE_TIMEZONE])
	print(CODE_NOT_FOUND_TIMEZONE, ":")
	println(codeMessages[CODE_NOT_FOUND_TIMEZONE])
	print(CODE_DUPLICATE_TIMEZONE, ":")
	println(codeMessages[CODE_DUPLICATE_TIMEZONE])
	print(CODE_ALREADY_EXIST_TIMEZONE_NAME, ":")
	println(codeMessages[CODE_ALREADY_EXIST_TIMEZONE_NAME])
	print(CODE_ALREADY_EXIST_TIMEZONE_HOUR_MINUTE, ":")
	println(codeMessages[CODE_ALREADY_EXIST_TIMEZONE_HOUR_MINUTE])
	print(CODE_FAIL_DURING_CREATE_TIMEZONE, ":")
	println(codeMessages[CODE_FAIL_DURING_CREATE_TIMEZONE])
	print(CODE_FAIL_DURING_UPDATE_TIMEZONE, ":")
	println(codeMessages[CODE_FAIL_DURING_UPDATE_TIMEZONE])
	print(CODE_FAIL_DURING_DELETE_TIMEZONE, ":")
	println(codeMessages[CODE_FAIL_DURING_DELETE_TIMEZONE])
}

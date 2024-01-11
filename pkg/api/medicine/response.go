package medicine_api

type Response[T any] struct {
	Header ResponseHeader  `json:"header"`
	Body   ResponseBody[T] `json:"body"`
}

type ResponseHeader struct {
	ResultCode string `json:"resultCode"`
	ResultMsg  string `json:"resultMsg"`
}

type ResponseBody[T any] struct {
	PageNo     int `json:"pageNo"`
	TotalCount int `json:"totalCount"`
	NumOfRows  int `json:"numOfRows"`
	Items      []T `json:"items"`
}

package common

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) Response {
	return Response{Code: 0, Msg: "ok", Data: data}
}

func Error(msg string) Response {
	return Response{Code: -1, Msg: msg, Data: nil}
}

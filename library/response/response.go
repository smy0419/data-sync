package response

type Response struct {
	Code uint16      `json:"code"`
	Data interface{} `json:"data"` // omitempty：When data is nil, not serialized
	Msg  string      `json:"msg"`
}

type Empty struct {
}

func OK(data interface{}) Response {
	return Response{
		Data: data,
		Msg:  "success",
	}
}

func Error(code uint16) Response {
	msg, _ := errorInfo[code]
	return Response{
		Code: code,
		Msg:  msg,
		Data: Empty{},
	}
}

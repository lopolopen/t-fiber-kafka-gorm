package dto

type Resp struct {
	Data any    `json:"data,omitempty"`
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}

func Err(err error) Resp {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	return Resp{
		Code: -1,
		Msg:  msg,
	}
}

func ErrMsg(msg string) Resp {
	return Resp{
		Code: -1,
		Msg:  msg,
	}
}

func OK[T any](data T) Resp {
	return Resp{
		Data: data,
		Code: 0,
	}
}

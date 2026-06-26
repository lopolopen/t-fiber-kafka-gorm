package dto

import "github.com/lopolopen/pkg/errorx"

type Resp struct {
	Data     any               `json:"data,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Message  string            `json:"message,omitempty"`
	Error    string            `json:"error,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

func Err(err *errorx.Error) Resp {
	var cause string
	if err := err.Unwrap(); err != nil {
		cause = err.Error()
	}
	return Resp{
		Reason:   err.Reason,
		Message:  err.Message,
		Metadata: err.Metadata,
		Error:    cause,
	}
}

func OK(data any) Resp {
	return Resp{
		Data: data,
	}
}

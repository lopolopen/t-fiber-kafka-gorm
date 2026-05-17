package dto

import "github.com/lopolopen/pkg/errorx"

type Resp[T any] struct {
	Data     T                 `json:"data,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Message  string            `json:"message,omitempty"`
	Error    string            `json:"error,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

func Err(err *errorx.Error) Resp[any] {
	var cause string
	if err := err.Unwrap(); err != nil {
		cause = err.Error()
	}
	return Resp[any]{
		Reason:   err.Reason,
		Message:  err.Message,
		Metadata: err.Metadata,
		Error:    cause,
	}
}

func OK[T any](data T) Resp[T] {
	return Resp[T]{
		Data: data,
	}
}

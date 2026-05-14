package dto

import "github.com/lopolopen/pkg/errorx"

type Resp[T any] struct {
	Data    T      `json:"data,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

func (r *Resp[T]) OK() bool {
	return r.Reason == ""
}

func Err(err *errorx.Error) Resp[struct{}] {
	return Resp[struct{}]{
		Reason:  err.Reason,
		Message: err.Message,
	}
}

func OK[T any](data T) Resp[T] {
	return Resp[T]{
		Data: data,
	}
}

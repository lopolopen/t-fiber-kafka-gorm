package errx

import (
	"github.com/lopolopen/pkg/errorx"
	"github.com/lopolopen/t-fiber-kafka-gorm/pkg/schema"
)

func ArgumentIsNil(argName string) error {
	return ErrNilArgument.WithMetadata(map[string]string{
		"argumentName": argName,
	})
}

var ErrInvalidIdempotencyKey = errorx.BadRequest(schema.ErrorReason_INVALID_IDEMPOTENCY_KEY.String(), "invalid idempotency key")
var ErrInvalidRequestFields = errorx.BadRequest(schema.ErrorReason_INVALID_REQUEST_FIELDS.String(), "invalid request fields")

var ErrNilArgument = errorx.InternalServer(schema.ErrorReason_NIL_ARGUMENT.String(), "nil argument")

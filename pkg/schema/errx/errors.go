package errx

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lopolopen/pkg/errorx"
	"github.com/lopolopen/t-fiber-kafka-gorm/pkg/schema"
)

func FrameworkErr(err *fiber.Error) *errorx.Error {
	return errorx.New(err.Code, schema.ErrorReason_FRAMEWORK_ERROR.String(), err.Message)
}

func ArgumentIsNil(argName string) *errorx.Error {
	return ErrNilArgument.WithMetadata(map[string]string{
		"argumentName": argName,
	})
}

var ErrUnspecified = errorx.InternalServer(schema.ErrorReason_UNSPECIFIED.String(), "")

var ErrInvalidIdempotencyKey = errorx.BadRequest(schema.ErrorReason_INVALID_IDEMPOTENCY_KEY.String(), "invalid idempotency key")
var ErrInvalidRequestFields = errorx.BadRequest(schema.ErrorReason_INVALID_REQUEST_FIELDS.String(), "invalid request fields")

var ErrNilArgument = errorx.InternalServer(schema.ErrorReason_NIL_ARGUMENT.String(), "nil argument")

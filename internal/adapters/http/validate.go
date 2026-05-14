package http

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/lopolopen/t-fiber-kafka-gorm/pkg/schema/errx"
)

type From int32

const (
	FromBody From = iota
	FromQuery
	FromParams
)

var validate = validator.New()

type Failure struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func validateStruct(s interface{}) []*Failure {
	var failures []*Failure
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var f Failure
			f.Field = err.StructNamespace()
			f.Tag = err.Tag()
			f.Value = err.Param()
			failures = append(failures, &f)
		}
	}
	return failures
}

// Validate is a generic middleware where T is the expected struct type
// and source specifies the data origin (e.g., body or query).
func Validate[T any](source From) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := new(T)

		var err error
		switch source {
		case FromQuery:
			err = c.QueryParser(payload)
		case FromParams:
			err = c.ParamsParser(payload)
		case FromBody:
			err = c.BodyParser(payload)
		}

		if err != nil {
			return err
		}

		if failures := validateStruct(payload); len(failures) > 0 {
			meta := make(map[string]string)
			for _, f := range failures {
				meta[f.Field] = fmt.Sprintf("failed on the '%s' tag, value: '%s'", f.Tag, f.Value)
			}
			return errx.ErrInvalidRequestFields.WithMetadata(meta)
		}

		c.Locals("payload", payload)
		return c.Next()
	}
}

package x

import (
	"fmt"
	"log/slog"
	"reflect"
	"runtime"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func SLogWithin(logger *slog.Logger, component any) *slog.Logger {
	c := componentName(component)
	return logger.With(slog.String("component", c))
}

func componentName(component any) string {
	c, ok := component.(string)
	var t reflect.Type
	if !ok {
		v := reflect.ValueOf(component)
		switch v.Kind() {
		case reflect.Func:
			c = funcName(v)
		case reflect.Pointer:
			t = v.Elem().Type()
		default:
			t = v.Type()
		}
	}
	if c == "" {
		c = fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())
	}
	return c
}

func funcName(v reflect.Value) string {
	funcPtr := runtime.FuncForPC(v.Pointer())
	if funcPtr == nil {
		return "unknown"
	}
	return funcPtr.Name()
}

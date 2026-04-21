package errx

import "fmt"

func ErrParamIsNil(param string) error {
	return fmt.Errorf("param %s is nil", param)
}

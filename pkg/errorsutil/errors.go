package errorsutil

import (
	"errors"
	"fmt"
)

func NotDefinedError(varName string) error {
	return errors.New(fmt.Sprintf("%s must be defined", varName))
}

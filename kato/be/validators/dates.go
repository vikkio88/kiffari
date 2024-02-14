package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var DateInTheFuture validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	if date.IsZero() {
		return true
	}

	today := time.Now()
	return !today.After(date)
}

package validator

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	// TODO: エラーをいい感じにする
  if err := cv.Validator.Struct(i); err != nil {
    return err
  }
  return nil
}
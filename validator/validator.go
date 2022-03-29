package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
  if err := cv.Validator.Struct(i); err != nil {
    for _, err := range err.(validator.ValidationErrors) {
      switch err.ActualTag() {
        case "required":
          return fmt.Errorf("%s は必須です", err.Field())
        case "email", "datetime", "uuid4":
          return fmt.Errorf("%s の形式が正しくありません", err.Field())
        case "min":
          return fmt.Errorf("%s は %s 文字以上にしてください", err.Field(), err.Param())
        case "max":
          return fmt.Errorf("%s は %s 文字以下にしてください", err.Field(), err.Param())
        case "gte":
          return fmt.Errorf("%s は %s 以上にしてください", err.Field(), err.Param())
        case "lte":
          return fmt.Errorf("%s は %s 以下にしてください", err.Field(), err.Param())
        default:
          return fmt.Errorf("%s が正しくありません err: %w", err.Field(), err)
      }
    }
  }
  return nil
}
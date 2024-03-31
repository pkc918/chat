package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkc918/chat/response"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

/**
 * 验证器 验证结构体值是否符合格式
 * @return {map[string]string} 当有未通过的字段 返回对应字段名称以及错误提示信息
 */

func Validator(target any) error {
	err := validate.Struct(target)
	if err != nil {
		exception := response.NewException(err.Error(), http.StatusBadRequest, response.ParameterIsInvalidCode)
		return exception
	}
	return nil
}

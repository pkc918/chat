package response

import (
	"net/http"
)

type Exception struct {
	message      string
	errorMessage map[string]string
	status       int
	code         int
}

func (e *Exception) Error() string {
	return e.message
}

func (e *Exception) ErrorMessage() map[string]string {
	return e.errorMessage
}

func (e *Exception) SetErrorMessage(field string, message string) {
	e.errorMessage[field] = message
}

func (e *Exception) SetMessage(message string) {
	e.message = message
}

func (e *Exception) Code() int {
	return e.code
}

func (e *Exception) Status() int {
	return e.status
}

func NewException(message string, status int, code int) *Exception {
	return &Exception{message: message, errorMessage: make(map[string]string), status: status, code: code}
}

// 自定义状态码
const (
	//
	ParameterIsInvalidCode = 1000 + iota
	//
	UserNotFoundCode
	//
	UserAlreadyBlockCode
	//
	UserAlreadyExistsCode
	//
	UserRegisterFailedCode
	//
	PasswordIncorrectCode
	//
	RecordNotExitsCode
	//
	RecordQueryFaildCode
	//
	RecordCreateFaildCode
	//
	RecordDeleteFailedCode
	//
	RecordUpdateFailedCode
	//
	AccountAuthorizeFailedCode
	//
	ServerErrorCode
	//
	RouterNotFoundCodeCode
	//
	RouterAuthorityLimitCode
)

var ErrParameterIsInvalid = NewException("Errors.ParameterIsInvalid", http.StatusBadRequest, ParameterIsInvalidCode)

var ErrUserNotFound = NewException("Errors.UserNotFound", http.StatusNotFound, UserNotFoundCode)

var ErrUserAlreadyBlock = NewException("Errors.UserAlreadyBlock", http.StatusBadRequest, UserAlreadyBlockCode)

var ErrUserAlreadyExists = NewException("Errors.UserAlreadyExists", http.StatusBadRequest, UserAlreadyExistsCode)

var ErrUserRegisterFailed = NewException("Errors.UserRegisterFailed", http.StatusBadRequest, UserRegisterFailedCode)

var ErrPasswordIncorrect = NewException("Errors.PasswordIncorrect", http.StatusBadRequest, PasswordIncorrectCode)

var ErrRecordNotExits = NewException("Errors.RecordNotExits", http.StatusNotFound, RecordNotExitsCode)

var ErrRecordQueryFaild = NewException("Errors.RecordQueryFaild", http.StatusBadRequest, RecordQueryFaildCode)

var ErrRecordCreateFaild = NewException("Errors.RecordCreateFaild", http.StatusBadRequest, RecordCreateFaildCode)

var ErrInternalServer = NewException("Errors.ServerError", http.StatusInternalServerError, ServerErrorCode)

var ErrRecordDeleteFailed = NewException("Errors.RecordDeleteFaild", http.StatusBadRequest, RecordDeleteFailedCode)

var ErrRecordUpdateFailed = NewException("Errors.RecordUpdateFaild", http.StatusBadRequest, RecordUpdateFailedCode)

var ErrAccountAuthorizeFailed = NewException("Errors.AccountAuthorizeFailed", http.StatusNonAuthoritativeInfo, AccountAuthorizeFailedCode)

var ErrAccountPermissionLimit = NewException("Errors.AccountPermissionLimit", http.StatusNonAuthoritativeInfo, AccountAuthorizeFailedCode)

var ErrRouterNotFound = NewException("Errors.RouterNotFound", http.StatusNotFound, RouterNotFoundCodeCode)

var ErrRouterForbidden = NewException("Errors.RouterForbidden", http.StatusForbidden, RouterAuthorityLimitCode)

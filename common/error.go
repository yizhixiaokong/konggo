package common

import "fmt"

type WebError interface {
	Code() int
	Msg() string

	Error() string

	Equal(err error) bool
}

type ResError struct {
	code    int
	message string
}

func (e *ResError) Code() int {
	return e.code
}

func (e *ResError) Msg() string {
	return e.message
}

// 实现 err 接口
func (e *ResError) Error() string {
	return fmt.Sprintf("code: %d, msg: %v", e.Code(), e.Msg())
}

func (e *ResError) Equal(err error) bool {
	return Cause(err).Code() == e.Code()
}

func (e *ResError) AddMsg(s interface{}) *ResError {
	e.message = e.message + fmt.Sprint(s)
	return e
}

func NewError(code int, message string) *ResError {
	return &ResError{
		code:    code,
		message: message,
	}
}

func Cause(err error) *ResError {
	if err == nil {
		return NewError(0, "")
	}

	e, ok := err.(*ResError)
	if !ok {
		e = NewError(-1, err.Error())
	}

	return e
}

func ErrNeedLogin() *ResError    { return NewError(StatusNeedLogin, MsgNeedLogin) }
func ErrServer() *ResError       { return NewError(StatusServerError, MsgServerError) }
func ErrInvalidParam() *ResError { return NewError(StatusInvalidParam, MsgInvalidParam) }
func ErrNotExist() *ResError     { return NewError(StatusTargetNotExist, MsgTargetNotExist) }
func ErrIsExist() *ResError      { return NewError(StatusTargetIsExist, MsgTargetIsExist) }
func ErrUnrealized() *ResError   { return NewError(StatusUnrealized, MsgUnrealized) }
func ErrInvalidParams(err interface{}) *ResError {
	msg := MsgInvalidParam
	if err != nil {
		msg += fmt.Sprintf(": %v", err)
	}
	return NewError(StatusInvalidParam, msg)
}

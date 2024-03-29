// Code generated by protoc-gen-zmicro-errno. DO NOT EDIT.
// versions:
// - protoc-gen-zmicro-errno v0.1.0
// - protoc                  v3.19.0
// source: errno/errno.proto

package errno

import (
	fmt "fmt"
	errors "github.com/zmicro-team/zmicro/core/errors"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the errors package it is being compiled against.
const _ = errors.ErrorsProtoPackageIsVersion3

type Option interface {
	apply(*errors.Error)
}
type optFunc func(e *errors.Error)

func (o optFunc) apply(e *errors.Error) { o(e) }
func WithMessage(s string) Option {
	return optFunc(func(e *errors.Error) {
		if s != "" {
			e.Message = s
		}
	})
}
func WithMetadata(k string, v string) Option {
	return optFunc(func(e *errors.Error) {
		if k != "" && v != "" {
			e.Metadata[k] = v
		}
	})
}
func _apply(e *errors.Error, opts ...Option) {
	for _, opt := range opts {
		opt.apply(e)
	}
}

func IsInternalServer(err error) bool {
	e := errors.FromError(err)
	return e.Code == 500
}
func ErrInternalServer(detail string) *errors.Error {
	return errors.New(500, "服务器错误", detail)
}
func ErrInternalServerf(format string, args ...any) *errors.Error {
	return errors.New(500, "服务器错误", fmt.Sprintf(format, args...))
}
func ErrInternalServerw(opt ...Option) *errors.Error {
	e := &errors.Error{
		Code:    500,
		Message: "服务器错误",
		Detail:  ErrorReason_internal_server.String(),
	}
	_apply(e, opt...)
	return e
}
func IsBadRequest(err error) bool {
	e := errors.FromError(err)
	return e.Code == 400
}
func ErrBadRequest(detail string) *errors.Error {
	return errors.New(400, "请求参数错误", detail)
}
func ErrBadRequestf(format string, args ...any) *errors.Error {
	return errors.New(400, "请求参数错误", fmt.Sprintf(format, args...))
}
func ErrBadRequestw(opt ...Option) *errors.Error {
	e := &errors.Error{
		Code:    400,
		Message: "请求参数错误",
		Detail:  ErrorReason_bad_request.String(),
	}
	_apply(e, opt...)
	return e
}
func IsCustom(err error) bool {
	e := errors.FromError(err)
	return e.Detail == ErrorReason_custom.String() && e.Code == 1000
}
func ErrCustom(message ...string) *errors.Error {
	if len(message) > 0 {
		return ErrCustomw(WithMessage(message[0]))
	}
	return ErrCustomw()
}
func ErrCustomf(format string, args ...any) *errors.Error {
	return ErrCustomw(WithMessage(fmt.Sprintf(format, args...)))
}
func ErrCustomw(opt ...Option) *errors.Error {
	e := &errors.Error{
		Code:    1000,
		Message: "自定义错误",
		Detail:  ErrorReason_custom.String(),
	}
	_apply(e, opt...)
	return e
}
func IsInvalidToken(err error) bool {
	e := errors.FromError(err)
	return e.Detail == ErrorReason_invalid_token.String() && e.Code == 1001
}
func ErrInvalidToken(message ...string) *errors.Error {
	if len(message) > 0 {
		return ErrInvalidTokenw(WithMessage(message[0]))
	}
	return ErrInvalidTokenw()
}
func ErrInvalidTokenf(format string, args ...any) *errors.Error {
	return ErrInvalidTokenw(WithMessage(fmt.Sprintf(format, args...)))
}
func ErrInvalidTokenw(opt ...Option) *errors.Error {
	e := &errors.Error{
		Code:    1001,
		Message: "无效的令牌",
		Detail:  ErrorReason_invalid_token.String(),
	}
	_apply(e, opt...)
	return e
}
func IsTokenExpired(err error) bool {
	e := errors.FromError(err)
	return e.Detail == ErrorReason_token_expired.String() && e.Code == 1002
}
func ErrTokenExpired(message ...string) *errors.Error {
	if len(message) > 0 {
		return ErrTokenExpiredw(WithMessage(message[0]))
	}
	return ErrTokenExpiredw()
}
func ErrTokenExpiredf(format string, args ...any) *errors.Error {
	return ErrTokenExpiredw(WithMessage(fmt.Sprintf(format, args...)))
}
func ErrTokenExpiredw(opt ...Option) *errors.Error {
	e := &errors.Error{
		Code:    1002,
		Message: "令牌过期",
		Detail:  ErrorReason_token_expired.String(),
	}
	_apply(e, opt...)
	return e
}
func IsTokenRevoked(err error) bool {
	e := errors.FromError(err)
	return e.Detail == ErrorReason_token_revoked.String() && e.Code == 1003
}
func ErrTokenRevoked(message ...string) *errors.Error {
	if len(message) > 0 {
		return ErrTokenRevokedw(WithMessage(message[0]))
	}
	return ErrTokenRevokedw()
}
func ErrTokenRevokedf(format string, args ...any) *errors.Error {
	return ErrTokenRevokedw(WithMessage(fmt.Sprintf(format, args...)))
}
func ErrTokenRevokedw(opt ...Option) *errors.Error {
	e := &errors.Error{
		Code:    1003,
		Message: "令牌已被吊销",
		Detail:  ErrorReason_token_revoked.String(),
	}
	_apply(e, opt...)
	return e
}
func IsLoginConflict(err error) bool {
	e := errors.FromError(err)
	return e.Detail == ErrorReason_login_conflict.String() && e.Code == 1004
}
func ErrLoginConflict(message ...string) *errors.Error {
	if len(message) > 0 {
		return ErrLoginConflictw(WithMessage(message[0]))
	}
	return ErrLoginConflictw()
}
func ErrLoginConflictf(format string, args ...any) *errors.Error {
	return ErrLoginConflictw(WithMessage(fmt.Sprintf(format, args...)))
}
func ErrLoginConflictw(opt ...Option) *errors.Error {
	e := &errors.Error{
		Code:    1004,
		Message: "登录冲突",
		Detail:  ErrorReason_login_conflict.String(),
	}
	_apply(e, opt...)
	return e
}

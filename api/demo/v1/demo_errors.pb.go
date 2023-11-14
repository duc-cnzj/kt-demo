// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsDemoUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DEMO_UNSPECIFIED.String() && e.Code == 500
}

func ErrorDemoUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DEMO_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsDemoNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DEMO_NOT_FOUND.String() && e.Code == 404
}

func ErrorDemoNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_DEMO_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsDemoNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DEMO_NOT_ALLOWED.String() && e.Code == 403
}

func ErrorDemoNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_DEMO_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

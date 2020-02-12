package cm

import (
	"fmt"
	"reflect"
)

type SettingNotFoundError struct {
	Key string
}

func NewSettingNotFoundError(key string) *SettingNotFoundError {
	return &SettingNotFoundError{key}
}

func (s *SettingNotFoundError) Error() string {
	return fmt.Sprintf("setting '%s' not found", s.Key)
}

type InvalidTypeError struct {
	Key      string
	Expected reflect.Type
	Actual   reflect.Type
}

func NewInvalidTypeError(key string, expected, actual reflect.Type) *InvalidTypeError {
	return &InvalidTypeError{
		Key:      key,
		Expected: expected,
		Actual:   actual,
	}
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("invalid type for key '%s': expected type '%s', got '%s'", e.Key, e.Expected, e.Actual)
}

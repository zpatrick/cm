package cm

import (
	"fmt"
	"reflect"
)

type KeyNotDefinedError struct {
	Key string
}

func NewKeyNotDefinedError(key string) *KeyNotDefinedError {
	return &KeyNotDefinedError{key}
}

func (e *KeyNotDefinedError) Error() string {
	return fmt.Sprintf("key not defined (key: '%s')", e.Key)
}

type KeyNotFoundError struct {
	Key string
}

func NewKeyNotFoundError(key string) *KeyNotFoundError {
	return &KeyNotFoundError{key}
}

func (e *KeyNotFoundError) Error() string {
	return fmt.Sprintf("key not found (key: '%s')", e.Key)
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
	return fmt.Sprintf("invalid type for key '%s': expected %s, got %s", e.Key, e.Expected, e.Actual)
}

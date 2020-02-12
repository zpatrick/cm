package cm

import "fmt"

type SettingNotFoundError struct {
	Key string
}

func NewSettingNotFoundError(key string) *SettingNotFoundError {
	return &SettingNotFoundError{key}
}

func (s *SettingNotFoundError) Error() string {
	return fmt.Sprintf("setting '%s' not found", s.Key)
}

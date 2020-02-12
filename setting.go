package cm

import (
	"errors"
	"fmt"
)

type IntSetting struct {
	Key           string
	Default       int
	DefaultIsZero bool
	Validate      IntValidator
	Providers     []IntProvider
}

func (s *IntSetting) Value() (int, error) {
	var value *int
	for _, p := range s.Providers {
		v, err := p.Int(s.Key)
		if err != nil {
			if snf := new(SettingNotFoundError); errors.Is(err, snf) {
				continue
			}

			return 0, fmt.Errorf("%s: failed to lookup int: %w", s.Key, err)
		}

		value = &v
		break
	}

	if value == nil {
		if !s.DefaultIsZero && s.Default == 0 {
			return 0, NewSettingNotFoundError(s.Key)
		}

		value = &s.Default
	}

	if err := s.Validate(*value); err != nil {
		return 0, fmt.Errorf("%s: failed validation: %w", s.Key, err)
	}

	return *value, nil
}

type IntValidator func(int) error

func ValidateIntBetween(lower, upper int) IntValidator {
	return IntValidator(func(v int) error {
		switch {
		case v < lower:
			return fmt.Errorf("%d is below lower limit %d", v, lower)
		case v > upper:
			return fmt.Errorf("%d is above upper limit %d", v, upper)
		default:
			return nil
		}
	})
}

func ValidateIntInSet(values ...int) IntValidator {
	set := make(map[int]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return IntValidator(func(v int) error {
		if _, ok := set[v]; !ok {
			return fmt.Errorf("%d not in set %v", v, set)
		}

		return nil
	})
}

type StringSetting struct {
	Key           string
	Default       string
	DefaultIsZero bool
	Validate      StringValidator
	Providers     []StringProvider
}

func (s *StringSetting) Value() (string, error) {
	var value *string
	for _, p := range s.Providers {
		v, err := p.String(s.Key)
		if err != nil {
			if snf := new(SettingNotFoundError); errors.Is(err, snf) {
				continue
			}

			return "", fmt.Errorf("%s: failed to lookup int: %w", s.Key, err)
		}

		value = &v
		break
	}

	if value == nil {
		if !s.DefaultIsZero && s.Default == "" {
			return "", NewSettingNotFoundError(s.Key)
		}

		value = &s.Default
	}

	if err := s.Validate(*value); err != nil {
		return "", fmt.Errorf("%s: failed validation: %w", s.Key, err)
	}

	return *value, nil
}

type StringValidator func(string) error

func ValidateStringInSet(values ...string) StringValidator {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}

	return StringValidator(func(v string) error {
		if _, ok := set[v]; !ok {
			return fmt.Errorf("%s not in set %v", v, set)
		}

		return nil
	})
}

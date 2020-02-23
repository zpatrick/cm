package cm

import (
	"errors"
	"fmt"
)

type IntSchema struct {
	Key           string
	Default       int
	DefaultIsZero bool
	Validate      IntValidator
	Providers     []IntProvider
}

func (s *IntSchema) Value() (int, error) {
	var value *int
	for _, p := range s.Providers {
		v, err := p.Int(s.Key)
		if err != nil {
			if knf := new(KeyNotFoundError); errors.As(err, &knf) {
				continue
			}

			return 0, fmt.Errorf("failed to lookup int '%s': %w", s.Key, err)
		}

		value = &v
		break
	}

	if value == nil {
		if s.Default == 0 && !s.DefaultIsZero {
			return 0, NewKeyNotFoundError(s.Key)
		}

		value = &s.Default
	}

	if s.Validate != nil {
		if err := s.Validate(*value); err != nil {
			return 0, fmt.Errorf("failed validation for int '%s': %w", s.Key, err)
		}
	}

	return *value, nil
}

type StringSchema struct {
	Key           string
	Default       string
	DefaultIsZero bool
	Validate      StringValidator
	Providers     []StringProvider
}

func (s *StringSchema) Value() (string, error) {
	var value *string
	for _, p := range s.Providers {
		v, err := p.String(s.Key)
		if err != nil {
			if knf := new(KeyNotFoundError); errors.As(err, &knf) {
				continue
			}

			return "", fmt.Errorf("failed to lookup string '%s': %w", s.Key, err)
		}

		value = &v
		break
	}

	if value == nil {
		if s.Default == "" && !s.DefaultIsZero {
			return "", NewKeyNotFoundError(s.Key)
		}

		value = &s.Default
	}

	if s.Validate != nil {
		if err := s.Validate(*value); err != nil {
			return "", fmt.Errorf("failed validation for string '%s': %w", s.Key, err)
		}
	}

	return *value, nil
}

type Schema struct {
	IntSchemas    []*IntSchema
	StringSchemas []*StringSchema
}

func (s *Schema) Validate() error {
	// no duplicate keys
	// no key has 0 providers
	// no key is nil or has nil providers
	// no key is empty string
	// should we get real values and call validate on each? add param to this method?
	return nil
}

type Config struct {
	intSchemas    map[string]*IntSchema
	stringSchemas map[string]*StringSchema
}

func NewConfig(s Schema) *Config {
	intSchemas := make(map[string]*IntSchema, len(s.IntSchemas))
	for _, is := range s.IntSchemas {
		intSchemas[is.Key] = is
	}

	stringSchemas := make(map[string]*StringSchema, len(s.StringSchemas))
	for _, ss := range s.StringSchemas {
		stringSchemas[ss.Key] = ss
	}

	return &Config{
		intSchemas:    intSchemas,
		stringSchemas: stringSchemas,
	}
}

func (c *Config) Int(key string) (int, error) {
	schema, ok := c.intSchemas[key]
	if !ok {
		return 0, NewKeyNotDefinedError(key)
	}

	return schema.Value()
}

func (c *Config) MustInt(key string) int {
	v, err := c.Int(key)
	if err != nil {
		panic(err)
	}

	return v
}

func (c *Config) String(key string) (string, error) {
	schema, ok := c.stringSchemas[key]
	if !ok {
		return "", NewKeyNotDefinedError(key)
	}

	return schema.Value()
}

func (c *Config) MustString(key string) string {
	v, err := c.String(key)
	if err != nil {
		panic(err)
	}

	return v
}
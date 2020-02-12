package cm

import (
	"fmt"
	"reflect"
)

// this coudl be an expensive object to carry around everywhere.
// maybe lazy evaluation is the beter way to go.

/*
	app.config.Validate()
	|-> How will this work? config.setting.each.Load()?

	app.config.String("redis.port")
	|-> config.providers.Lookup("redis.port")
		|-> providers[0].EnvVar uses os.Getenv
		|-> providers[1].File uses FileProvider.ReloadPolicy: [ReloadAlways, ReloadNever, ReloadAfter, ReloadOn(...)]
			since this is first time it loads.
	-> gets from file provider. does string conversion.
*/

type Settings map[string]interface{}

func (s Settings) Int(key string) (int, error) {
	v, ok := s[key]
	if !ok {
		return 0, NewSettingNotFoundError(key)
	}

	iv, ok := v.(int)
	if !ok {
		return 0, NewInvalidTypeError(key, reflect.TypeOf(iv), reflect.TypeOf(v))
	}

	return iv, nil
}

func (s Settings) MustInt(key string) int {
	iv, err := s.Int(key)
	if err != nil {
		panic(err)
	}

	return iv
}

func (s Settings) String(key string) (string, error) {
	v, ok := s[key]
	if !ok {
		return "", NewSettingNotFoundError(key)
	}

	sv, ok := v.(string)
	if !ok {
		return "", NewInvalidTypeError(key, reflect.TypeOf(sv), reflect.TypeOf(v))
	}

	return sv, nil
}

func (s Settings) MustString(key string) string {
	sv, err := s.String(key)
	if err != nil {
		panic(err)
	}

	return sv
}

/*
	Current system seems a bit weird: Providers load settings and return them as keys.
	Settings load themselves from providers, convert them, and them set them in the config.
	Config just loads and assumes it is of the correct type? or throw type error?
*/

type Config struct {
	Providers []Provider
	Settings  []Setting
	values    map[string]interface{}
}

func (c *Config) Load() (Settings, error) {
	settings := Settings{}
	for _, s := range c.Settings {
		if err := s.Load(c, settings); err != nil {
			return Settings{}, fmt.Errorf("failed to load key '%s': %w", s.Key(), err)
		}
	}

	return settings
}

func (c *Config) Int(key string) (int, error) {
	sv, err := c.Get(key)
	if err != nil {
		return 0, err
	}

	iv, ok := v.(int)
	if !ok {
		return 0, NewInvalidTypeError(key, reflect.TypeOf(iv), reflect.TypeOf(v))
	}

	return iv, nil
}

func (c *Config) MustInt(key string) int {
	iv, err := c.Int(key)
	if err != nil {
		panic(err)
	}

	return iv
}

func (c *Config) Set(key string, value interface{}) {
	c.values[key] = value
}

func (c *Config) Get(key string) (value string, err error) {

	/*
		for _, p := range c.Providers {
			v, err := p.Lookup(key)
			if err != nil {
				var knf *SettingNotFoundError
				if errors.As(err, &knf) {
					continue
				}

				return "", fmt.Errorf("config: failed to lookup key '%s': %w", key, err)
			}

			return v, nil
		}

		return "", NewSettingNotFoundError(key)
	*/
}

/*


































import (
	"errors"
	"fmt"
	"reflect"
)

type Setting interface {
	Default() string
	Validate(key string, val interface{}) error
}

type StringValidator func(key, v string) error

type StringSetting struct {
	DefaultValue string
	Validators   []StringValidator
}

func (s StringSetting) Default() interface{} {
	return s.DefaultValue
}

func (s StringSetting) Validate(key string, v interface{}) error {
	sv, ok := v.(string)
	if !ok {
		return NewInvalidTypeError(key, reflect.TypeOf(sv), reflect.TypeOf(v))
	}

	for _, validate := range s.Validators {
		if err := validate(key, sv); err != nil {
			return err
		}
	}

	return nil
}

type IntValidator func(key string, v int) error

type IntSetting struct {
	DefaultValue int
	Validators   []IntValidator
}

func (i IntSetting) Default() interface{} {
	return i.DefaultValue
}

func (i IntSetting) Validate(key string, v interface{}) error {
	iv, ok := v.(int)
	if !ok {
		return NewInvalidTypeError(key, reflect.TypeOf(iv), reflect.TypeOf(v))
	}

	for _, validate := range i.Validators {
		if err := validate(key, iv); err != nil {
			return err
		}
	}

	return nil
}

type Loader interface {
	Load(key string) (interface{}, error)
}

type StaticLoader map[string]interface{}

func (s StaticLoader) Load(key string) (interface{}, error) {
	v, ok := s[key]
	if !ok {
		return nil, NewSettingNotFoundError(key)
	}

	return v, nil
}

type EnvironmentLoader map[string]string

func (e EnvironmentLoader) Load(key string) (interface{}, error) {
	return nil, nil
}

type Schema struct {
	Settings map[string]Setting
	Loaders  []Loader
}

type Config struct {
	settings map[string]interface{}
}

func Load(schema Schema) (*Config, error) {
	cfg := &Config{
		settings: map[string]interface{}{},
	}

	for key, setting := range schema.Settings {
		for _, loader := range schema.Loaders {
			v, err := loader.Load(key)
			if err != nil {
				var knf *SettingNotFoundError
				if errors.As(err, &knf) {
					continue
				}

				return nil, fmt.Errorf("failed to load key '%s': %w", key, err)
			}

			if err := setting.Validate(key, v); err != nil {
				return nil, fmt.Errorf("failed to validate key '%s': %w", key, err)
			}

			cfg.Set(key, v)
			break
		}

		if !cfg.Exists(key) {
			cfg.Set(key, setting.Default())
		}
	}

	return cfg, nil
}

func (c *Config) Set(key string, v interface{}) {
	c.settings[key] = v
}

func (c *Config) Int(key string) (int, error) {
	v, ok := c.settings[key]
	if !ok {
		return 0, NewSettingNotFoundError(key)
	}

	iv, ok := v.(int)
	if !ok {
		return 0, NewInvalidTypeError(key, reflect.TypeOf(iv), reflect.TypeOf(v))
	}

	return iv, nil
}

func (c *Config) MustInt(key string) int {
	iv, err := c.Int(key)
	if err != nil {
		panic(err)
	}

	return iv
}

func (c *Config) Exists(key string) bool {
	_, ok := c.settings[key]
	return ok
}

*/

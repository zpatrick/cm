package cm

import (
	"fmt"
	"strconv"
)

type Setting interface {
	Key() string
	Load(cfg *Config) error
}

type IntSetting struct {
	SettingKey string
}

func (i IntSetting) Key() string {
	return i.SettingKey
}

func (i IntSetting) Load(cfg *Config) error {
	v, err := cfg.Get(i.Key())
	if err != nil {
		return err
	}

	if _, err := strconv.ParseInt(v, 10, 64); err != nil {
		return fmt.Errorf("failed to convert '%s' to an int: %w", v, err)
	}

	cfg.Set(i.Key(), v)
	return nil
}

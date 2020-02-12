package cm

import (
	"os"
	"time"
)

type Provider interface {
	Load() error
	Lookup(key string) (value string, err error)
}

type StaticProvider map[string]string

func (s StaticProvider) Load() error {
	return nil
}

func (s StaticProvider) Lookup(key string) (string, error) {
	v, ok := s[key]
	if !ok {
		return "", NewSettingNotFoundError(key)
	}

	return v, nil
}

type EnvironmentProvider map[string]string

func (e EnvironmentProvider) Load() error {
	return nil
}

func (e EnvironmentProvider) Lookup(key string) (string, error) {
	ev, ok := e[key]
	if !ok {
		return "", NewSettingNotFoundError(key)
	}

	v := os.Getenv(ev)
	if v == "" {
		return "", NewSettingNotFoundError(key)
	}

	return v, nil
}

type DataParser func(data []byte) (map[string]string, error)

func YAMLParser() DataParser {
	return func(data []byte) (map[string]string, error) {
		return nil, nil
	}
}

func JSONParser() DataParser {
	return func(data []byte) (map[string]string, error) {
		return nil, nil
	}
}

func INIParser() DataParser {
	return func(data []byte) (map[string]string, error) {
		return nil, nil
	}
}

type ReloadPolicy func() bool

func ReloadAlways() ReloadPolicy {
	return func() bool { return true }
}

func ReloadNever() ReloadPolicy {
	return func() bool { return false }
}

func ReloadEvery(ticker *time.Ticker) ReloadPolicy {
	return func() bool {
		select {
		case <-ticker.C:
			return true
		default:
			return false
		}
	}
}

func ReloadOnFileChange(path string) ReloadPolicy {
	return func() bool {
		return false
	}
}

type FileProvider struct {
	Path         string
	Parser       DataParser
	ReloadPolicy ReloadPolicy
}

package cm

import (
	"fmt"
	"os"
	"strconv"
)

type EnvironmentProvider struct{}

func NewEnvironmentProvider() EnvironmentProvider {
	return EnvironmentProvider{}
}

func (e EnvironmentProvider) get(key, envvar string) (string, error) {
	v := os.Getenv(envvar)
	if v == "" {
		return "", fmt.Errorf("EnvironmentProvider: %w", NewSettingNotFoundError(key))
	}

	return v, nil
}

func (e EnvironmentProvider) String(envvar string) StringProvider {
	return StringProviderFunc(func(key string) (string, error) {
		v, err := e.get(key, envvar)
		if err != nil {
			return "", err
		}

		return v, nil
	})
}

func (e EnvironmentProvider) Int(envvar string) IntProvider {
	return IntProviderFunc(func(key string) (int, error) {
		v, err := e.get(key, envvar)
		if err != nil {
			return 0, err
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("EnvironmentProvider: %s: failed to convert '%v' to int: %w", key, v, err)
		}

		return i, nil
	})
}

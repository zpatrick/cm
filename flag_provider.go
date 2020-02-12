package cm

import (
	"flag"
	"fmt"
)

type FlagProvider struct {
	*flag.FlagSet
}

func NewFlagProvider(name string, errorHandling flag.ErrorHandling) *FlagProvider {
	return &FlagProvider{flag.NewFlagSet(name, errorHandling)}
}

func (f *FlagProvider) assertParsed() error {
	if !f.Parsed() {
		return fmt.Errorf("FlagProvider: flag set has not been parsed")
	}

	return nil
}

func (f *FlagProvider) Int(name string, value int, usage string) IntProvider {
	ptr := f.FlagSet.Int(name, value, usage)
	return IntProviderFunc(func(key string) (int, error) {
		if err := f.assertParsed(); err != nil {
			return 0, err
		}

		return *ptr, nil
	})
}

func (f *FlagProvider) String(name, value, usage string) StringProvider {
	ptr := f.FlagSet.String(name, value, usage)
	return StringProviderFunc(func(key string) (string, error) {
		if err := f.assertParsed(); err != nil {
			return "", err
		}

		return *ptr, nil
	})
}

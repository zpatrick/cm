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
		return fmt.Errorf("FlagSet.Parse has not been called!")
	}

	return nil
}

func (f *FlagProvider) isFlagPassed(name string) bool {
	var passed bool
	f.Visit(func(f *flag.Flag) {
		if f.Name == name {
			passed = true
		}
	})

	return passed
}

// TODO: need a way to allow bypassing if the flag doesn't get set.
// Setting it to a special value may not work. ZeroValue is tricky because then
// they can never set default. Could have param 'skipIf int`, but that it wierd for other types i think.
// or is it?
func (f *FlagProvider) Int(name string, value int, usage string, useDefault bool) IntProvider {
	ptr := f.FlagSet.Int(name, value, usage)
	return IntProviderFunc(func(key string) (int, error) {
		if err := f.assertParsed(); err != nil {
			return 0, fmt.Errorf("FlagProvider: %w", err)
		}

		if !useDefault && !f.isFlagPassed(name) {
			return 0, fmt.Errorf("FlagProvider: %w", NewKeyNotFoundError(key))
		}

		return *ptr, nil
	})
}

func (f *FlagProvider) String(name, value, usage string, useDefault bool) StringProvider {
	ptr := f.FlagSet.String(name, value, usage)
	return StringProviderFunc(func(key string) (string, error) {
		if err := f.assertParsed(); err != nil {
			return "", fmt.Errorf("FlagProvider: %w", err)
		}

		if !useDefault && !f.isFlagPassed(name) {
			return "", fmt.Errorf("FlagProvider: %w", NewKeyNotFoundError(key))
		}

		return *ptr, nil
	})
}

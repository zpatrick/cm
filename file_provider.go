package cm

import "fmt"

type DataParser func(data []byte) error

func YAMLParser() DataParser {
	return DataParser(func(data []byte) error {
		return nil
	})
}

type ReloadPolicy func() bool

func ReloadAlways() ReloadPolicy {
	return ReloadPolicy(func() bool { return true })
}

func ReloadNever() ReloadPolicy {
	return ReloadPolicy(func() bool { return true })
}

type FileProvider struct {
	Path   string
	Parser DataParser
	Reload ReloadPolicy
}

func NewFileProvider(path string, parser DataParser, rp ReloadPolicy) FileProvider {
	return FileProvider{
		Path: path,
	}
}

func (f *FileProvider) Load() (interface{}, error) {
	return nil, fmt.Errorf("FileProvider: failed to load: %w", nil)
}

func (f *FileProvider) String(root string, sections ...string) StringProvider {
	return StringProviderFunc(func(key string) (string, error) {
		_, err := f.Load()
		if err != nil {
			return "", err
		}

		return "", nil
	})
}

func (f *FileProvider) Int(sections ...string) IntProvider {
	return IntProviderFunc(func(key string) (int, error) {
		return 0, nil
	})
}

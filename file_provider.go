package cm

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type DataParser func(content []byte) (fields map[string]interface{}, err error)

func ParseYAML(content []byte) (map[string]interface{}, error) {
	var fields map[string]interface{}
	if err := yaml.Unmarshal(content, &fields); err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	return fields, nil
}

type ReloadPolicy func() bool

func ReloadAlways() bool { return true }

func ReloadNever() bool { return false }

type FileProvider struct {
	Path   string
	Parse  DataParser
	Reload ReloadPolicy
	fields map[string]interface{}
}

func NewFileProvider(path string, parse DataParser, reload ReloadPolicy) FileProvider {
	return FileProvider{
		Path:   path,
		Parse:  parse,
		Reload: reload,
	}
}

func (f *FileProvider) loadFields() (map[string]interface{}, error) {
	if f.fields == nil || f.Reload() {
		content, err := ioutil.ReadFile(f.Path)
		if err != nil {
			return nil, fmt.Errorf("FileProvider: failed to load %s: %w", f.Path, err)
		}

		fields, err := f.Parse(content)
		if err != nil {
			return nil, fmt.Errorf("FileProvider: failed to parse file contents: %w", err)
		}

		f.fields = fields
	}

	return f.fields, nil
}

func (f *FileProvider) loadField(key string, root string, sections ...string) (interface{}, error) {
	fields, err := f.loadFields()
	if err != nil {
		return nil, err
	}

	chain := append([]string{root}, sections...)
	for i, c := range chain {
		val, ok := fields[c]
		if !ok {
			return nil, NewSettingNotFoundError(key)
		}

		if i == len(chain)-1 {
			return val, nil
		}

		f, ok := val.(map[string]interface{})
		if !ok {
			return nil, NewInvalidTypeError(key)
		}

		fields = f
	}

	return nil, fmt.Errorf("how did we get here")
}

func (f *FileProvider) String(root string, sections ...string) StringProvider {
	return StringProviderFunc(func(key string) (string, error) {
		value, err := f.loadField(key, root, sections...)
		if err != nil {
			return "", err
		}

		if value == nil {
			return "", NewSettingNotFoundError(key)
		}

		s, ok := value.(string)
		if !ok {
			return "", NewInvalidTypeError(key)
		}

		return s, nil
	})
}

func (f *FileProvider) Int(sections ...string) IntProvider {
	return IntProviderFunc(func(key string) (int, error) {
		return 0, nil
	})
}

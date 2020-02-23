package cm

import (
	"fmt"
	"io/ioutil"
	"reflect"

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
			return nil, fmt.Errorf("failed to load %s: %w", f.Path, err)
		}

		fields, err := f.Parse(content)
		if err != nil {
			return nil, fmt.Errorf("failed to parse file contents: %w", err)
		}

		f.fields = fields
	}

	return f.fields, nil
}

func (f *FileProvider) loadField(key string, section string, subsections ...string) (interface{}, error) {
	fields, err := f.loadFields()
	if err != nil {
		return nil, err
	}

	chain := append([]string{section}, subsections...)
	for i := len(chain) - 1; i > 0; i-- {
		sectionKey := chain[0]
		chain = chain[1:]

		section, ok := fields[sectionKey]
		if !ok {
			return nil, NewKeyNotFoundError(key)
		}

		switch s := section.(type) {
		case map[string]interface{}:
			fields = s
		case map[interface{}]interface{}:
			subfields := make(map[string]interface{}, len(s))
			for subsectionKey, subsectionValue := range s {
				subsectionKeyStr, ok := subsectionKey.(string)
				if !ok {
					return nil, NewInvalidTypeError(key, reflect.TypeOf(subsectionKeyStr), reflect.TypeOf(subsectionKey))
				}

				subfields[subsectionKeyStr] = subsectionValue
			}

			fields = subfields
		default:
			return nil, NewInvalidTypeError(key, reflect.TypeOf(map[string]interface{}{}), reflect.TypeOf(section))
		}
	}

	val, ok := fields[chain[0]]
	if !ok {
		return nil, NewKeyNotFoundError(key)
	}

	return val, nil
}

func (f *FileProvider) String(section string, subsections ...string) StringProvider {
	return StringProviderFunc(func(key string) (string, error) {
		val, err := f.loadField(key, section, subsections...)
		if err != nil {
			return "", err
		}

		s, ok := val.(string)
		if !ok {
			return "", fmt.Errorf("FileProvider: %w", NewInvalidTypeError(key, reflect.TypeOf(s), reflect.TypeOf(val)))
		}

		return s, nil
	})
}

func (f *FileProvider) Int(section string, subsections ...string) IntProvider {
	return IntProviderFunc(func(key string) (int, error) {
		val, err := f.loadField(key, section, subsections...)
		if err != nil {
			return 0, err
		}

		i, ok := val.(int)
		if !ok {
			return 0, fmt.Errorf("FileProvider: %w", NewInvalidTypeError(key, reflect.TypeOf(i), reflect.TypeOf(val)))
		}

		return i, nil
	})
}

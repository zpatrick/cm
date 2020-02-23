package cm

import "testing"

type ValidateWrapperFunc func(schema *Schema) error

func TestValidate(t *testing.T) {
	validateWrapperFunc := func(schema *Schema) error {
		return schema.Validate()
	}

	testValidationSuite(t, validateWrapperFunc)
}

func TestProviderRunsValidations(t *testing.T) {
	validateWrapperFunc := func(schema *Schema) error {
		_, err := schema.Provider()
		return err
	}

	testValidationSuite(t, validateWrapperFunc)
}

func testValidationSuite(t *testing.T, validateWrapperFunc ValidateWrapperFunc) {
	cases := []struct {
		Name   string
		Schema *Schema
	}{
		{
			Name: "IntSchema with no providers",
			Schema: &Schema{
				IntSchemas: []*IntSchema{
					{Key: "key"},
				},
			},
		},
		{
			Name: "StringSchema with no providers",
			Schema: &Schema{
				StringSchemas: []*StringSchema{
					{Key: "key"},
				},
			},
		},
		{
			Name: "Two IntSchemas with the same key",
			Schema: &Schema{
				IntSchemas: []*IntSchema{
					{Key: "key"},
					{Key: "key"},
				},
			},
		},
		{
			Name: "Two StringSchemas with the same key",
			Schema: &Schema{
				StringSchemas: []*StringSchema{
					{Key: "key"},
					{Key: "key"},
				},
			},
		},
		{
			Name: "IntSchema and StringSchema with the same key",
			Schema: &Schema{
				IntSchemas: []*IntSchema{
					{Key: "key"},
				},
				StringSchemas: []*StringSchema{
					{Key: "key"},
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if err := validateWrapperFunc(c.Schema); err == nil {
				t.Fatal("error was nil!")
			}
		})
	}
}

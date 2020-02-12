package cm

/*
	What if I wanted to mix static and dynamic loading?
	For exmample, live file reload.


func TestSchema(t *testing.T) {
	schema := Schema{
		Loaders: []Loader{
			StaticLoader{
				"redis.host": "localhost",
				"redis.port": 5000,
			},
			FileLoader{
				Path:   "config.yaml",
				Parser: YAMLParser,
			},
			EnvironmentLoader{},
		},
		Settings: map[string]Setting{
			"redis.host": StringSetting{
				DefaultValue: "localhost",
				Validators:   []StringValidator{},
			},
			"redis.port": IntSetting{
				DefaultValue: 4000,
				Validators: []IntValidator{
					IntMustExist,
				},
			},
		},
	}

	cfg, err := Load(schema)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(cfg.MustInt("redis.port"))
}
*/

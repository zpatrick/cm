package cm

/*
	SettingSchema{
		IntSchema{
			Key: "redis.port",
			Validate: ValidateIntBetween(0, 100),
			Loaders: []IntLoader{
				environment.LoadInt("APP_REDIS_PORT"),
				yaml.LoadInt("redis", "port"),
			},
		},
	}


	env := cm.NewEnvVarProvider()
	yaml := cm.NewFileProvider("path", cm.FormatYAML, cm.ReloadNever)
	redis := cm.NewRedisProvider("endpoint")

	Settings{
		IntSettings: []IntSetting{
			"redis.port": {
				Default: 9090,
				Validate: ValidateIntBetween(0, 100),
				Providers: []IntProvider{
					env.Int("APP_REDIS_PORT"),
					yaml.Int("redis", "port"),
				},
			},
		},
		StringSettings: map[string]StringSetting{
			"redis.host": {
				Default: "localhost",
				Providers: []StringProvider{
					env.String("APP_REDIS_HOST"),
					yaml.String("redis", "host"),
				},
			},
		},
	}
*/

type Schema struct {
	IntSettings    map[string]IntSetting
	StringSettings map[string]StringSetting
}

type Config struct {
	ints map[string]IntSetting
}

func (c *Config) Validate() error {
	// same key isn't defined miltuple times.
	// all settings have at least 1 provider.
	return nil
}

func (c *Config) Int(key string) (int, error) {
	setting, ok := c.ints[key]
	if !ok {
		return 0, NewSettingNotFoundError(key)
	}

	return setting.Value()
}

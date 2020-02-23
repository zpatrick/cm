package cm

import (
	"flag"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	flag := NewFlagProvider("", flag.ContinueOnError)
	env := NewEnvironmentProvider()
	file := NewFileProvider("config.yaml", ParseYAML, ReloadNever)

	schema := Schema{
		IntSettings: map[string]*IntSetting{
			"port": {
				Default:  9090,
				Validate: ValidateIntBetween(0, 65535),
				Providers: []IntProvider{
					file.Int("port"),
					flag.Int("port", 9090, "the application port"),
					env.Int("APP_PORT"),
				},
			},
			"redis.port": {
				Default:  4000,
				Validate: ValidateIntBetween(0, 65535),
				Providers: []IntProvider{
					env.Int("APP_REDIS_PORT"),
					flag.Int("redis-port", 4000, "redis port"),
					file.Int("redis", "port"),
				},
			},
		},
		StringSettings: map[string]*StringSetting{
			"redis.host": {
				Default:  "localhost",
				Validate: func(v string) error { return nil },
				Providers: []StringProvider{
					env.String("APP_REDIS_HOST"),
					flag.String("redis-host", "localhost", "redis host"),
					file.String("redis", "host"),
				},
			},
		},
	}

	t.Log(schema)

	if err := flag.Parse(os.Args); err != nil {
		t.Fatal(err)
	}

	host, err := schema.StringSettings["redis.host"].Value()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Host: %v", host)
}

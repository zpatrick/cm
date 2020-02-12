package cm

import (
	"flag"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	env := NewEnvironmentProvider()
	flag := NewFlagProvider("", flag.ContinueOnError)

	schema := Schema{
		IntSettings: map[string]IntSetting{
			"port": {
				Default:  9090,
				Validate: ValidateIntBetween(0, 65535),
				Providers: []IntProvider{
					env.Int("APP_PORT"),
					flag.Int("port", 9090, "the application port"),
				},
			},
			"redis.port": {
				Default:  4000,
				Validate: ValidateIntBetween(0, 65535),
				Providers: []IntProvider{
					env.Int("APP_REDIS_PORT"),
					flag.Int("redis-port", 4000, "redis port"),
				},
			},
		},
		StringSettings: map[string]StringSetting{
			"redis.host": {
				Default:  "localhost",
				Validate: func(v string) error { return nil },
				Providers: []StringProvider{
					env.String("APP_REDIS_HOST"),
					flag.String("redis-host", "localhost", "redis host"),
				},
			},
		},
	}

	if err := flag.Parse(os.Args); err != nil {
		t.Fatal(err)
	}

	t.Log(schema)
}

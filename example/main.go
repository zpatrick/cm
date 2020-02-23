package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zpatrick/cm"
)

func main() {
	environment := cm.NewEnvironmentProvider()
	flagProvider := cm.NewFlagProvider("app", flag.ContinueOnError)
	fileProvider := cm.NewFileProvider("config.yaml", cm.ParseYAML, cm.ReloadNever)

	// TODO: required?
	schema := cm.Schema{
		IntSchemas: []*cm.IntSchema{
			{
				Key:      "redis.port",
				Default:  4000,
				Validate: cm.ValidateIntBetween(0, 65535),
				Providers: []cm.IntProvider{
					flagProvider.Int("redis-port", 4000, "redis port", false),
					fileProvider.Int("redis", "port"),
					environment.Int("APP_REDIS_PORT"),
				},
			},
		},
		StringSchemas: []*cm.StringSchema{
			{
				Key:     "redis.host",
				Default: "localhost",
				Providers: []cm.StringProvider{
					flagProvider.String("redis-host", "localhost", "redis host", false),
					fileProvider.String("redis", "host"),
					environment.String("APP_REDIS_HOST"),
				},
			},
		},
	}

	if err := flagProvider.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	if err := schema.Validate(); err != nil {
		log.Fatal(err)
	}

	cfg := cm.NewConfig(schema)
	redisPort, err := cfg.Int("redis.port")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Port:", redisPort)

	redisHost, err := cfg.String("redis.host")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Host:", redisHost)
}

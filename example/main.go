package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/zpatrick/cm"
)

func main() {
	envProvider := cm.NewEnvironmentProvider()
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
					envProvider.Int("APP_REDIS_PORT"),
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
					envProvider.String("APP_REDIS_HOST"),
				},
			},
		},
	}

	if err := flagProvider.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	p, err := schema.Provider()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Redis Host:", p.MustString("redis.host"))
	fmt.Println("Redis Port:", p.MustInt("redis.port"))
}

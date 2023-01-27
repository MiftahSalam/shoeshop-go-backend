package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"shoeshop-backend/src/infrastructure/logger"
	"shoeshop-backend/src/shared/database"
)

type Configuration struct {
	Application Application      `json:"application"`
	Logger      logger.LogOption `json:"logger"`
	Database    database.Option  `json:"database"`
}

type Application struct {
	Name              string `json:"name"`
	Hostname          string `json:"hostname"`
	HttpPort          int64  `json:"http_port"`
	GQLHttpPort       int64  `json:"gql_http_port"`
	EnableMonitoring  bool   `json:"enable_monitoring"`
	MonitoringAddress string `json:"monitoring_address"`
}

func New(target interface{}) error {
	var (
		filename = os.Getenv("CONFIG_FILE")
	)

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", target); err != nil {
			return err
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return err
	}

	if err := envconfig.Process("", target); err != nil {
		return err
	}

	return nil
}

func (c *Configuration) HttpPort() string {
	return fmt.Sprintf(":%v", c.Application.HttpPort)
}

func (c *Configuration) GQLHttpPort() string {
	return fmt.Sprintf(":%v", c.Application.GQLHttpPort)
}

func Setup() *Configuration {
	fmt.Println("Try Setup Configuration ... ")

	c := &Configuration{}
	err := New(c)
	if err != nil {
		panic(err)
	}

	c.Application.Hostname, _ = os.Hostname()

	b, _ := json.Marshal(c)
	fmt.Println(string(b))

	return c
}

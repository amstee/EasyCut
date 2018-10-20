package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
	"fmt"
)

type DataConfig struct {
	Port 	int			`json:"port"`
	Origins []string	`json:"origins"`
	Jwks 	string		`json:"jwks"`
	Env		string		`json:"env"`
}

func (c *DataConfig) Display() {
	fmt.Println("Running with config : ")
	fmt.Printf("Port: %d\n", c.Port)
	fmt.Println("Origins: ")
	for _, o := range c.Origins {
		fmt.Println(o)
	}
	fmt.Println("Jwks: " + c.Jwks)
	fmt.Println("Env: " + c.Env)
}

func (c *DataConfig) LoadEnv() (error) {
	if port := os.Getenv("AUTH_PORT"); port != "" {
		p, err := strconv.Atoi(port); if err != nil {
			return err
		}
		c.Port = p
	}
	if jwks := os.Getenv("AUTH_JWKS"); jwks != "" {
		c.Jwks = jwks
	}
	if env := os.Getenv("AUTH_ENV"); env != "" {
		c.Env = env
	}
	return nil
}

func (c *DataConfig) LoadConfig() (error) {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	viper.SetDefault("port", "8080")
	viper.SetDefault("jwks", "https://easy-cut.eu.auth0.com/.well-known/jwks.json")
	viper.SetDefault("env", "dev")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	return c.LoadEnv()
}

var Content = new(DataConfig)
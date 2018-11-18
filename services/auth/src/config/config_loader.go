package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
	"fmt"
)

type DataConfig struct {
	Name 	string 		`json:"name"`
	Version string 		`json:"version"`
	Port 	int			`json:"port"`
	Origins []string	`json:"origins"`
	Issuer string 		`json:"issuer"`
	Jwks 	string		`json:"jwks"`
	Perms	string		`json:"perms"`
	Env		string		`json:"env"`
	TPrefix string 		`json:"tprefix"`
}

func (c *DataConfig) Display() {
	fmt.Println("__________________CONFIGURATION__________________")
	fmt.Println("Service -------- " + c.Name + " Version : " + c.Version + " --------")
	fmt.Printf("\tPort =\t\t%d\n", c.Port)
	fmt.Println("\tJwks =\t\t" + c.Jwks)
	fmt.Printf("\tBearer prefix =\t%s\n", c.TPrefix)
	fmt.Println("\tEnv =\t\t " + c.Env)
	fmt.Println("CORS :")
	for i, o := range c.Origins {
		fmt.Printf("\tOrigin %d is %s\n", i, o)
	}
	fmt.Println("__________________CONFIGURATION__________________")
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
	viper.SetDefault("name", "auth")
	viper.SetDefault("port", "8080")
	viper.SetDefault("perms", "http://perms:8080")
	viper.SetDefault("issuer", "https://easy-cut.eu.auth0.com/")
	viper.SetDefault("tprefix", "auth0|")
	viper.SetDefault("version", "0.0.1")
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
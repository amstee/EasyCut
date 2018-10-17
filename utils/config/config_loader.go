package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
	"fmt"
)

type Permission struct {
	Route string			`json:"route"`
	Permissions []string	`json:"permissions"`
}

type DataConfig struct {
	Port 	int									`json:"port"`
	Origins []string							`json:"origins"`
	Security string								`json:"security"`
	Env		string								`json:"env"`
	Routes []Permission							`json:"routes"`
}


func (c *DataConfig) Display() {
	fmt.Println("Running with config : ")
	fmt.Printf("Port: %d\n", c.Port)
	fmt.Println("Origins: ")
	for _, o := range c.Origins {
		fmt.Println(o)
	}
	fmt.Println("PermMatcher : ")
	for _, perm := range c.Routes {
		fmt.Println("Route : " + perm.Route)
		for _, v := range perm.Permissions {
			fmt.Println(v)
		}
	}
	fmt.Println("Security url: " + c.Security)
	fmt.Println("Env: " + c.Env)
}

func (c *DataConfig) LoadEnv() (error) {
	if port := os.Getenv("AUTH_PORT"); port != "" {
		p, err := strconv.Atoi(port); if err != nil {
			return err
		}
		c.Port = p
	}
	if security := os.Getenv("AUTH_JWKS"); security != "" {
		c.Security = security
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
	viper.SetDefault("security", "http://auth")
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
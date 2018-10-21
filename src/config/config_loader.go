package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type Permission struct {
	Route string			`json:"route"`
	Permissions []string	`json:"permissions"`
}

type Secrets struct {
	ClientId string		`json:"client_id"`
	ClientSecret string	`json:"client_secret"`
}

type DataConfig struct {
	Port int			`json:"port"`
	Origins []string	`json:"origins"`
	Security string		`json:"security"`
	Env	string			`json:"env"`
	Api	string			`json:"api"`
	Domain string		`json:"domain"`
	Oauth string		`json:"oauth"`
	Sfile string		`json:"sfile"`
	ClientId string		`json:"client_id"`
	ClientSecret string	`json:"client_secret"`
	Routes []Permission	`json:"routes"`
}

func (c *DataConfig) GetApi() string {
	return c.Domain + c.Api
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
	fmt.Println("Domain " + c.Domain)
	fmt.Println("Oauth " + c.Oauth)
	fmt.Println("Api "+ c.Api)
}

func (c *DataConfig) LoadSecrets() (error) {
	fmt.Println("Trying to open secrets file : " + c.Sfile)
	jsonFile, err := os.Open(c.Sfile); if err != nil {
		fmt.Println("unable to open secrets file")
		return nil
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile); if err != nil {
		return err
	}
	var secret Secrets
	if err := json.Unmarshal(byteValue, &secret); err != nil {
		return err
	}
	if secret.ClientId != "" {
		c.ClientId = secret.ClientId
	}
	if secret.ClientSecret != "" {
		c.ClientSecret = secret.ClientSecret
	}
	return nil
}

func (c *DataConfig) LoadEnv() (error) {
	if port := os.Getenv("EASY_CUT_PORT"); port != "" {
		p, err := strconv.Atoi(port); if err != nil {
			return err
		}
		c.Port = p
	}
	if security := os.Getenv("AUTH_URL"); security != "" {
		c.Security = security
	}
	if env := os.Getenv("EASY_CUT_ENV"); env != "" {
		c.Env = env
	}
	if clientId := os.Getenv("API_CLIENT_ID"); clientId != "" {
		c.ClientId = clientId
	} else {
		fmt.Println("Warning : API_CLIENT_ID not set in environment")
	}
	if clientSecret := os.Getenv("API_CLIENT_SECRET"); clientSecret != "" {
		c.ClientSecret = clientSecret
	} else {
		fmt.Println("Warning : API_CLIENT_SECRET not set in environment")
	}
	return nil
}

func (c *DataConfig) LoadConfig() (error) {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	viper.SetDefault("port", "8080")
	viper.SetDefault("domain", "https://easy-cut.eu.auth0.com/")
	viper.SetDefault("sfile", "/run/secrets/auth0_api")
	viper.SetDefault("api", "api/v2")
	viper.SetDefault("security", "http://auth:8080")
	viper.SetDefault("oauth", "oauth/token")
	viper.SetDefault("env", "dev")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	if err := c.LoadSecrets(); err != nil {
		return err
	}
	return c.LoadEnv()
}

var Content = new(DataConfig)
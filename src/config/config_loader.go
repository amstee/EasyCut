package config

import (
	"github.com/spf13/viper"
	"os"
	"strconv"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/pkg/errors"
)

type Permission struct {
	Match bool				`json:"match"`
	Route string			`json:"route"`
	Permissions []string	`json:"permissions"`
}

type Secrets struct {
	ClientId string		`json:"client_id"`
	ClientSecret string	`json:"client_secret"`
}

type Service struct {
	Name string 		`json:"name"`
	Url string 			`json:"url"`
}

type OAuth struct {
	Use bool 			`json:"use"`
	Audience string 	`json:"audience"`
	Domain string 		`json:"domain"`
	Extension string 	`json:"extension"`
}

type Api struct {
	Use bool 			`json:"use"`
	Domain string 		`json:"domain"`
	Extension string 	`json:"extension"`
	Tprefix string 		`json:"tprefix"`
}

type DataConfig struct {
	Auth0 bool 			`json:"auth0"`
	Version string 		`json:"version"`
	Name string 		`json:"name"`
	Port int			`json:"port"`
	Env	string			`json:"env"`
	Sfile string		`json:"sfile"`
	ClientId string		`json:"client_id"`
	ClientSecret string	`json:"client_secret"`
	Origins []string	`json:"origins"`
	Oauth OAuth  		`json:"oauth"`
	Api Api 			`json:"api"`
	Services []Service 	`json:"services"`
	Routes []Permission	`json:"routes"`
}

func GetOauth() string {
	return Content.GetOauth()
}

func (c *DataConfig) GetOauth() string {
	return c.Oauth.Domain + c.Oauth.Extension
}

func GetApi() string {
	return Content.GetApi()
}

func (c *DataConfig) GetApi() string {
	return c.Api.Domain + c.Api.Extension
}

func Display() {
	Content.Display()
}

func (c *DataConfig) Display() {
	e := false
	fmt.Println("__________________CONFIGURATION__________________")
	fmt.Println("Service -------- " + c.Name + " Version : " + c.Version + " --------")
	fmt.Printf("\tPort =\t\t%d\n", c.Port)
	fmt.Printf("\tEnvironment =\t%s\n", c.Env)
	if c.Oauth.Use {
		fmt.Println("Oauth configuration :")
		fmt.Printf("\tOauth domain =\t%s\n", c.Oauth.Domain)
		fmt.Printf("\tOauth extension =\t%s\n", c.Oauth.Extension)
	}
	if c.Api.Use {
		fmt.Println("Api configuration :")
		fmt.Printf("\tApi domain =\t%s\n", c.Api.Domain)
		fmt.Printf("\tApi extension =\t%s\n", c.Api.Extension)
		fmt.Printf("\tApi Bearer prefix =\t%s\n", c.Api.Tprefix)
	}
	fmt.Println("CORS :")
	for i, o := range c.Origins {
		fmt.Printf("\tOrigin %d is %s\n", i, o)
	}
	fmt.Println("Services :")
	for _, s := range c.Services {
		fmt.Printf("\tService : %s\t--> %s\n", s.Name, s.Url)
	}
	fmt.Println("Routes :")
	for _, r := range c.Routes {
		fmt.Printf("\tRoute : %s\n", r.Route)
		e = false
		for _, p := range r.Permissions {
			fmt.Printf("\t\tRequired role --> %s\n", p)
			e = true
		}
		if !e {
			fmt.Println("\t\tNo permissions needed")
		}
	}
	fmt.Println("__________________CONFIGURATION__________________")
}

func GetServiceURL(name string) string {
	for _, s := range Content.Services {
		if s.Name == name {
			return s.Url
		}
	}
	return ""
}

func LoadSecrets() (error) {
	return Content.LoadSecrets()
}

func (c *DataConfig) LoadSecrets() (error) {
	jsonFile, err := os.Open(c.Sfile); if err != nil {
		fmt.Println("## WARNING ## secrets file not found, make sure env variables are set")
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

func LoadEnv(OauthEnvPrefix string) (error) {
	return Content.LoadEnv(OauthEnvPrefix)
}

func (c *DataConfig) LoadEnv(OauthEnvPrefix string) (error) {
	if port := os.Getenv("EASY_CUT_PORT"); port != "" {
		p, err := strconv.Atoi(port); if err != nil {
			return err
		}
		c.Port = p
	}
	if security := os.Getenv("AUTH_URL"); security != "" {
		c.Services = append(c.Services, Service{Name: "security", Url: security})
	}
	if user := os.Getenv("USER_URL"); user != "" {
		c.Services = append(c.Services, Service{Name: "user", Url: user})
	}
	if env := os.Getenv("EASY_CUT_ENV"); env != "" {
		c.Env = env
	}
	if c.Oauth.Use {
		if clientId := os.Getenv(OauthEnvPrefix + "_ID"); clientId != "" {
			c.ClientId = clientId
		} else {
			fmt.Println("Warning : Oauth client ID not set in environment")
		}
		if clientSecret := os.Getenv(OauthEnvPrefix + "_SECRET"); clientSecret != "" {
			c.ClientSecret = clientSecret
		} else {
			fmt.Println("Warning : Oauth client SECRET not set in environment")
		}
	}
	return nil
}

func Load(OauthEnvPrefix string) (error) {
	return Content.LoadConfig(OauthEnvPrefix)
}

func (c *DataConfig) LoadConfig(OauthEnvPrefix string) (error) {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	viper.SetDefault("port", "8080")
	viper.SetDefault("domain", "https://easy-cut.eu.auth0.com/")
	viper.SetDefault("sfile", "/run/secrets/auth0_api")
	viper.SetDefault("api", "api/v2")
	viper.SetDefault("oauth", "oauth/token")
	viper.SetDefault("version", "0.0.1")
	viper.SetDefault("auth0", false)
	viper.SetDefault("env", "dev")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(c); err != nil {
		return err
	}
	if c.Name == "" {
		return errors.New("Service name not specified")
	}
	if c.Auth0 {
		if err := c.LoadSecrets(); err != nil {
			return err
		}
	}
	return c.LoadEnv(OauthEnvPrefix)
}

var Content = new(DataConfig)
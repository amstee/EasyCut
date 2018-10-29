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
	MatchUser bool			`json:"match_user"`
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

type DataConfig struct {
	Name string 		`json:"name"`
	Port int			`json:"port"`
	TPrefix string 		`json:"tprefix"`
	Env	string			`json:"env"`
	Api	string			`json:"api"`
	Domain string		`json:"domain"`
	Oauth string		`json:"oauth"`
	Sfile string		`json:"sfile"`
	ClientId string		`json:"client_id"`
	ClientSecret string	`json:"client_secret"`
	Origins []string	`json:"origins"`
	Services []Service 	`json:"services"`
	Routes []Permission	`json:"routes"`
}

func GetApi() string {
	return Content.GetApi()
}

func (c *DataConfig) GetApi() string {
	return c.Domain + c.Api
}

func Display() {
	Content.Display()
}

func (c *DataConfig) Display() {
	e := false
	fmt.Println("_________CONFIGURATION_________")
	fmt.Println("Service -------- " + c.Name + " --------")
	fmt.Printf("\tPort =\t\t%d\n", c.Port)
	fmt.Printf("\tEnvironment =\t%s\n", c.Env)
	fmt.Println("Auth0 configuration :")
	fmt.Printf("\tApi url =\t%s\n", c.Api)
	fmt.Printf("\tBearer prefix =\t%s\n", c.TPrefix)
	fmt.Printf("\tOauth url =\t%s\n", c.Oauth)
	fmt.Printf("\tDomain url =\t%s\n", c.Domain)
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
	fmt.Println("_________CONFIGURATION_________")
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

func LoadEnv() (error) {
	return Content.LoadEnv()
}

func (c *DataConfig) LoadEnv() (error) {
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

func Load() (error) {
	return Content.LoadConfig()
}

func (c *DataConfig) LoadConfig() (error) {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	viper.SetDefault("port", "8080")
	viper.SetDefault("domain", "https://easy-cut.eu.auth0.com/")
	viper.SetDefault("sfile", "/run/secrets/auth0_api")
	viper.SetDefault("api", "api/v2")
	viper.SetDefault("oauth", "oauth/token")
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
	if err := c.LoadSecrets(); err != nil {
		return err
	}
	return c.LoadEnv()
}

var Content = new(DataConfig)
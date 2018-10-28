package common

import (
	"github.com/urfave/negroni"
	"fmt"
	"github.com/amstee/easy-cut/src/config"
	"strconv"
)

func Run(service *negroni.Negroni) {
	fmt.Printf("Starting service %s on port %d\n", config.Content.Name, config.Content.Port)
	service.Run(":" + strconv.Itoa(config.Content.Port))
}

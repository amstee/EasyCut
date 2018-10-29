package common

import (
	"github.com/urfave/negroni"
	"github.com/amstee/easy-cut/src/config"
	"strconv"
)

func Run(service *negroni.Negroni) {
	service.Run(":" + strconv.Itoa(config.Content.Port))
}

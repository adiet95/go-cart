package config

import (
	"os"

	"github.com/adiet95/go-cart/src/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start apllication",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = os.Getenv("PORT")
		gin.SetMode(gin.ReleaseMode)
		mainRoute.Run(addrs)
		return nil
	} else {
		return err
	}
}

package routers

import (
	"errors"

	"github.com/adiet95/go-cart/src/database"
	"github.com/adiet95/go-cart/src/modules"
	"github.com/gin-gonic/gin"
)

func New() (*gin.Engine, error) {
	mainRoute := gin.Default()

	db, err := database.New()
	if err != nil {
		return nil, errors.New("failed init database")
	}

	modules.New(mainRoute, db)

	return mainRoute, nil
}

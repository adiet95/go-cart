package modules

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(rt *gin.Engine, db *gorm.DB) {
	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route := rt.Group("/cart")
	{
		route.GET("", ctrl.GetCart)
		route.GET("/name", ctrl.GetName)
		route.GET("/qty", ctrl.GetQty)
		route.POST("", ctrl.AddCart)
		route.DELETE("", ctrl.DeleteCart)

	}

}

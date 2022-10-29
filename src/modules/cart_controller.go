package modules

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/adiet95/go-cart/src/database"
	"github.com/adiet95/go-cart/src/interfaces"
	"github.com/adiet95/go-cart/src/libs"
	"github.com/gin-gonic/gin"
)

type cart_ctrl struct {
	svc interfaces.CartService
}

func NewCtrl(reps interfaces.CartService) *cart_ctrl {
	return &cart_ctrl{svc: reps}
}

func (r *cart_ctrl) GetCart(c *gin.Context) {
	r.svc.GetCart().Send(c)
}

func (r *cart_ctrl) AddCart(c *gin.Context) {
	var data database.Cart
	err := json.NewDecoder(c.Request.Body).Decode(&data)
	if err != nil {
		libs.New(err.Error(), 400, true)
		c.Abort()
	}

	r.svc.AddCart(&data).Send(c)
}

func (r *cart_ctrl) DeleteCart(c *gin.Context) {
	val := c.Request.URL.Query().Get("code")
	r.svc.DeleteCart(val).Send(c)
}

func (r *cart_ctrl) GetName(c *gin.Context) {
	val := c.Request.URL.Query().Get("name")
	v := strings.ToLower(val)
	r.svc.GetName(v).Send(c)
}

func (r *cart_ctrl) GetQty(c *gin.Context) {
	val := c.Request.URL.Query().Get("qty")
	v, _ := strconv.Atoi(val)

	r.svc.GetQty(v).Send(c)
}

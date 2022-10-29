package interfaces

import (
	"github.com/adiet95/go-cart/src/database"
	"github.com/adiet95/go-cart/src/libs"
)

type CartRepo interface {
	AddCart(data *database.Cart) (*database.Cart, error)
	GetCart() (*database.Carts, error)
	DeleteCart(code string) error
	FindCart(code string) (*database.Cart, error)
	UpdateCart(data *database.Cart, code string) error
	GetByName(name string) (*database.Carts, error)
	GetByQty(qty int) (*database.Carts, error)
}

type CartService interface {
	AddCart(data *database.Cart) *libs.Response
	DeleteCart(code string) *libs.Response
	GetCart() *libs.Response
	GetName(name string) *libs.Response
	GetQty(qty int) *libs.Response
}

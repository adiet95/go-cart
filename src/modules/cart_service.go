package modules

import (
	"github.com/adiet95/go-cart/src/database"
	"github.com/adiet95/go-cart/src/interfaces"
	"github.com/adiet95/go-cart/src/libs"
)

type cart_service struct {
	cart_repo interfaces.CartRepo
}

func NewService(reps interfaces.CartRepo) *cart_service {
	return &cart_service{reps}
}

func (r *cart_service) AddCart(data *database.Cart) *libs.Response {
	res, err := r.cart_repo.FindCart(data.ProductCode)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	if res.ProductCode == data.ProductCode {
		data.Quantity = data.Quantity + res.Quantity
		err1 := r.cart_repo.UpdateCart(data, data.ProductCode)
		if err1 != nil {
			return libs.New(err.Error(), 400, true)
		}
		return libs.New("product update successful", 202, false)
	}
	re, err := r.cart_repo.AddCart(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(re, 202, false)
}

func (r *cart_service) DeleteCart(code string) *libs.Response {
	res := r.cart_repo.DeleteCart(code)
	if res != nil {
		return libs.New(res.Error(), 400, true)
	}
	return libs.New("product delete successful", 204, false)
}

func (r *cart_service) GetCart() *libs.Response {
	res, err := r.cart_repo.GetCart()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (r *cart_service) GetName(name string) *libs.Response {
	res, err := r.cart_repo.GetByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

func (r *cart_service) GetQty(qty int) *libs.Response {
	res, err := r.cart_repo.GetByQty(qty)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(res, 200, false)
}

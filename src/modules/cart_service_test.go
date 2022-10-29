package modules

import (
	"testing"

	"github.com/adiet95/go-cart/src/database"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestGetCart(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMocks = database.Carts{
		{ProductName: "Mobil Langkah", Quantity: 1},
		{ProductName: "Motor", Quantity: 3},
	}

	repo.mock.On("GetCart").Return(&dataMocks, nil)

	data := service.GetCart()
	res := data.Data.(*database.Carts)

	for i, v := range *res {
		assert.Equal(t, dataMocks[i].ProductName, v.ProductName)
		assert.Equal(t, dataMocks[i].Quantity, v.Quantity)
	}
}

func TestAddCart(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)
	var dataMock = database.Cart{ProductName: "Motor", Quantity: 3, ProductCode: "a-01"}

	repo.mock.On("FindCart", "a-01").Return(&dataMock, nil)
	repo.mock.On("UpdateCart", &dataMock, "a-01").Return(&dataMock, nil)

	data := service.AddCart(&dataMock)
	res := data.IsError
	var expect bool = false

	assert.Equal(t, expect, res)
}

func TestDelete(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = database.Cart{ProductName: "Motor", Quantity: 3, CartId: "12345"}
	repo.mock.On("DeleteUser", "12345").Return(&dataMock, nil)

	data := service.DeleteCart("12345")
	res := data

	assert.Equal(t, 204, res.Code)
}

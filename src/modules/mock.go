package modules

import (
	"github.com/adiet95/go-cart/src/database"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) AddCart(data *database.Cart) (*database.Cart, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*database.Cart), nil
}

func (m *RepoMock) GetCart() (*database.Carts, error) {
	args := m.mock.Called()
	return args.Get(0).(*database.Carts), nil
}

func (m *RepoMock) DeleteCart(code string) error {
	return nil
}

func (m *RepoMock) FindCart(code string) (*database.Cart, error) {
	args := m.mock.Called(code)
	return args.Get(0).(*database.Cart), nil
}

func (m *RepoMock) UpdateCart(data *database.Cart, code string) error {
	return nil
}

func (m *RepoMock) GetByName(name string) (*database.Carts, error) {
	args := m.mock.Called(name)
	return args.Get(0).(*database.Carts), nil
}

func (m *RepoMock) GetByQty(qty int) (*database.Carts, error) {
	args := m.mock.Called(qty)
	return args.Get(0).(*database.Carts), nil
}

package modules

import (
	"errors"

	"github.com/adiet95/go-cart/src/database"
	"gorm.io/gorm"
)

type cart_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *cart_repo {
	return &cart_repo{db}
}

func (r *cart_repo) AddCart(data *database.Cart) (*database.Cart, error) {
	var datas database.Carts
	res := r.db.Where("cart_id = ?", data.CartId).Find(&datas)

	if res.RowsAffected != 0 {
		return nil, errors.New("cart already registered")
	}

	result := r.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failled to obtain data")
	}
	return data, nil
}

func (r *cart_repo) GetCart() (*database.Carts, error) {
	var datas database.Carts
	result := r.db.Select("product_code", "product_name", "quantity").Order("product_name asc").Find(&datas)

	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return &datas, nil
}

func (r *cart_repo) DeleteCart(code string) error {
	var data *database.Cart
	var datas *database.Carts
	res := r.db.Where("product_code = ?", code).Find(&datas)

	if res.RowsAffected == 0 {
		return errors.New("data not found")
	}
	re := r.db.Model(datas).Where("product_code = ?", code).Delete(&data)
	if re.Error != nil {
		return errors.New("failed to delete data")
	}
	return nil
}

func (r *cart_repo) FindCart(code string) (*database.Cart, error) {
	var data *database.Cart
	var datas *database.Carts

	res := r.db.Model(&datas).Where("product_code = ?", code).Find(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	return data, nil
}

func (r *cart_repo) UpdateCart(data *database.Cart, code string) error {
	res := r.db.Model(&data).Where("product_code = ?", code).Updates(&data)

	if res.Error != nil {
		return errors.New("failed to update data")
	}
	return nil
}

func (r *cart_repo) GetByName(name string) (*database.Carts, error) {
	var datas database.Carts

	result := r.db.Where("LOWER(product_name) LIKE ?", "%"+name+"%").Find(&datas)
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return &datas, nil
}

func (r *cart_repo) GetByQty(qty int) (*database.Carts, error) {
	var datas database.Carts

	result := r.db.Where("quantity = ?", qty).Find(&datas).Order("quantity asc")
	if result.Error != nil {
		return nil, errors.New("failed obtain datas")
	}
	return &datas, nil
}

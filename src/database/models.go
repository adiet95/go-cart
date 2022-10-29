package database

type Cart struct {
	CartId      string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"cartId,omitempty"`
	ProductCode string `json:"kodeProduk,omitempty"`
	ProductName string `json:"namaProduk,omitempty"`
	Quantity    int    `json:"kuantitas,omitempty"`
}

type Carts []Cart

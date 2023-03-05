package models

type Cart struct {
	ID            int                 `json:"id" gorm:"primary_key:auto_increment"`
	UserID        int                 `json:"user_id" gorm:"type: int"`
	User          UsersCartResponse   `json:"user"`
	ProductID     int                 `json:"product_id" gorm:"type: int"`
	Product       ProductCartResponse `json:"product"`
	OrderQuantity int                 `json:"order_quantity" gorm:"type: int"`
	CreatedAt     int                 `json:"-"`
	UpdatedAt     int                 `json:"-"`
}

type CartUserResponse struct {
	ProductID     int `json:"product_id" gorm:"type: int"`
	OrderQuantity int `json:"order_quantity"`
	UserID        int `json:"-"`
}

type CartProductResponse struct {
	ProductID     int                 `json:"-"`
	Product       ProductCartResponse `json:"product"`
	OrderQuantity int                 `json:"order_quantity"`
	UserID        int                 `json:"user_id" gorm:"type: int"`
}

func (CartUserResponse) TableName() string {
	return "cart"
}

func (CartProductResponse) TableName() string {
	return "cart"
}

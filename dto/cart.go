package dto

type CartResponse struct {
	ProductID     int `json:"product_id" gorm:"type: int"`
	OrderQuantity int `json:"order_quantity" gorm:"type: int"`
	UserID        int `json:"user_id"`
}

type CartRequest struct {
	ProductID     int `json:"product_id" gorm:"type: int"`
	OrderQuantity int `json:"order_quantity" gorm:"type: int" validate:"required"`
	UserID        int `json:"user_id"`
}

package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindCarts() ([]models.Cart, error)
	// GetCart(ID int) (models.Cart, error)
	// CreateCart(cart models.Cart) (models.Cart, error)
	// UpdateCart(cart models.Cart) (models.Cart, error)
	// DeleteCart(cart models.Cart) (models.Cart, error)
}

func RepositoryCart(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCarts() ([]models.Cart, error) {
	var cart []models.Cart
	err := r.db.Preload("User").Preload("Product").Find(&cart).Error
	return cart, err
}

package dto

type ProfileResponse struct {
	ID      int    `json:"id" gorm:"primary_key:auto_increment"`
	Phone   string `json:"phone" gorm:"type: varchar(255)"`
	Address string `json:"address" gorm:"type: text"`
}

type CreateProfileRequest struct {
	Phone   string `json:"phone" form:"phone" validate:"required"`
	Address string `json:"address" form:"address" validate:"required"`
}

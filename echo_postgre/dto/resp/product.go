package dto

type ProductResponseDTO struct {
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Inventory   uint   `json:"inventory" gorm:"column:inventory"`
	Price       uint   `json:"price" gorm:"column:price"`
	Thumbnail   string `json:"thumbnail" gorm:"column:thumbnail"`
}

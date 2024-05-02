package dto

type OrderCartResponseDTO struct {
	ProductName string `json:"productName"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	Thumbnail   string `json:"thumbnail"`
}

type OrderDetailResponseDTO struct {
	ProductName string `json:"productName"`
	Username    string `json:"username"`
	Quantity    uint   `json:"quantity"`
	Price       uint   `json:"price"`
	Thumbnail   string `json:"thumbnail"`
}

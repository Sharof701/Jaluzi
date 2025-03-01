package models

type Product struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	Price        float64 `json:"price"`
	ProductImage string  `json:"product_image"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
	DeletedAt    string  `json:"delete_at,omitempty"`
}

type ProductCreate struct {
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	Price        float64 `json:"price"`
	ProductImage string  `json:"product_image"`
}

type ProductUpdate struct {
	Id           string  `json:"-"`
	Name         string  `json:"name"`
	Code         string  `json:"code"`
	Price        float64 `json:"price"`
	ProductImage string  `json:"product_image"`
}

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type ProductGetListRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ProductGetListResponse struct {
	Product []*Product `json:"product"`
	Total   int64      `json:"total"`
}

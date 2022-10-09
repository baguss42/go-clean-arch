package dto

// ProductRequest store data related request with controller & service
type ProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Qty         int    `json:"qty"`
}

// ProductResponse store data related response with controller & service
type ProductResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Qty         int    `json:"qty"`
}

// ProductParams store data related with repository
type ProductParams struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Qty         int    `json:"qty"`
}

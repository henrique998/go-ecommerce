package requests

type CreateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Price       int64  `json:"price"`
	StockQty    int    `json:"stock_qty"`
}

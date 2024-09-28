package dto

type CreateStockDto struct {
	StockName string `json:"stock_name"`
	Name      string `json:"name"`
	IsBalanz  bool   `json:"is_balanz"`
	Type      string `json:"type"`
}

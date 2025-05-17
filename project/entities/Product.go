package entities

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryID int     `json:"categoryId"`
	UnitPrice  float64 `json:"unitPrice"`
	Views      int     `json:"views"`
}

package entities

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryID int     `json:"categoryId"`
	UnitPrice  float64 `json:"unitPrice"`
	Views      int     `json:"views"`
}

func (p *Product) GetID() int {
	return p.ID
}

func (p *Product) SetID(id int) {
	p.ID = id
}

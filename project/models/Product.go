package entities

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	CategoryID int     `json:"categoryId"`
	UnitPrice  float64 `json:"unitPrice"`
	Views      int     `json:"views"`
}

func (p *Product) GetID() int {
	return p.Id
}

func (p *Product) SetID(id int) {
	p.Id = id
}

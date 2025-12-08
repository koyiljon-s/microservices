package data

import "encoding/json"
import "time"
import "io"

type Product struct {
	ID          int `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float32 `json:"price"`
	SKU         string `json:"sku"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return ProductList
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd34",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
}
package libs

import (
	"time"
)

type Products struct {
	ID            uint64      `json:"id"`
	Detail        ProductType `json:"detail"`
	Trademark     string      `json:"trademark"`
	Manufacturing string      `json:"manufacturing"`
	SKU           string      `json:"sku"`
	CreatedAt     time.Time   `json:"created_at"`
}

func (this pg) ListProducts() (products []Products, count int, err error) {
	rows, err := this.db.Query(`SELECT p.id, pt.id, pt.name, pt.detail, p.trademark, p.manufacturing, p."SKU", p.created_at FROM products as p JOIN product_type as pt on p.detail = pt.id;`)
	defer rows.Close()
	if err != nil {
		return products, -1, err
	}

	for rows.Next() {
		count ++
		var product Products
		rows.Scan(&product.ID, &product.Detail.ID, &product.Detail.Name, &product.Detail.Detail, &product.Trademark, &product.Manufacturing, &product.SKU, &product.CreatedAt)
		products = append(products, product)
	}

	return products, count, nil
}

func (this pg) GetProductById(id int64) (Products, int, error){
	var product Products
	row := this.db.QueryRow(`SELECT p.id, pt.id, pt.name, pt.detail, p.trademark, p.manufacturing, p."SKU", p.created_at FROM products as p JOIN product_type as pt on p.detail = pt.id WHERE p.id = $1 LIMIT 1;`, id)
	if row.Err() != nil{
		return product, -1, row.Err()
	}
	err := row.Scan(&product.ID, &product.Detail.ID, &product.Detail.Name, &product.Detail.Detail, &product.Trademark, &product.Manufacturing, &product.SKU, &product.CreatedAt)
	if err != nil {
		return Products{}, 0, err
	}
	return product, 1, nil
}

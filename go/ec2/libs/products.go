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

func (this pg) ListProducts() (products []Products, err error) {
	rows, err := this.db.Query("SELECT * FROM products;")
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product Products
		rows.Scan(&product.ID, &product.Detail, &product.Trademark, &product.Manufacturing, &product.SKU, &product.CreatedAt)
		products = append(products, product)
	}

	return products, nil
}

func (this pg) GetProductById(id int64) (Products, error){
	var product Products
	row := this.db.QueryRow("SELECT * FROM products WHERE id = $1 LIMIT 1;", id)
	if row.Err() != nil{
		return product, row.Err()
	}
	row.Scan(&product.ID, &product.Detail, &product.Trademark, &product.Manufacturing, &product.SKU, &product.CreatedAt)
	return product, nil
}

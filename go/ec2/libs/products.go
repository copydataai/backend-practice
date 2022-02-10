package libs

import (
	"time"
	"fmt"
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
	rows, err := this.db.Query("SELECT * FROM products;")
	defer rows.Close()
	if err != nil {
		return products, -1, err
	}

	for rows.Next() {
		count ++
		var product Products
		err := rows.Scan(&product.ID, &product.Detail, &product.Trademark, &product.Manufacturing, &product.SKU, &product.CreatedAt)
		fmt.Println(err)
		products = append(products, product)
	}

	return products, count, nil
}

func (this pg) GetProductById(id int64) (Products, int, error){
	var product Products
	row := this.db.QueryRow("SELECT * FROM products WHERE id = $1 LIMIT 1;", id)
	if row.Err() != nil{
		return product, -1, row.Err()
	}
	err := row.Scan(&product.ID, &product.Detail, &product.Trademark, &product.Manufacturing, &product.SKU, &product.CreatedAt)
	if err != nil {
		return Products{}, -1, err
	}
	return product, 1, nil
}

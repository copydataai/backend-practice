package libs


type ProductType struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func (this pg) ListProductTypes() (productTypes []ProductType, count int, err error) {
	rows, err := this.db.Query("SELECT * FROM product_type;")
	if err != nil {
		return productTypes, -1, err
	}

	for rows.Next() {
		count++
		var productType ProductType
		rows.Scan(&productType.ID, &productType.Name, &productType.Detail)
		productTypes = append(productTypes, productType)
	}

	return productTypes, count, nil
}


func (this pg) GetProductTypeById(id int64) (ProductType, int, error){
	var productType ProductType
	row := this.db.QueryRow("SELECT * FROM product_type WHERE id = $1 LIMIT 1;", id)
	if row.Err() != nil{
		return productType, -1, row.Err()
	}
	err := row.Scan(&productType.ID, &productType.Name, &productType.Detail)
	if err != nil {
		return ProductType{}, 0, err
	}
	return productType, 1, nil
}

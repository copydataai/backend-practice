package libs


type ProductType struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

func (this pg) ListProductTypes() (productTypes []ProductType, err error) {
	rows, err := this.db.Query("SELECT * FROM product_type;")
	if err != nil {
		return productTypes, err
	}

	for rows.Next() {
		var productType ProductType
		rows.Scan(&productType.ID, &productType.Name, &productType.Detail)
		productTypes = append(productTypes, productType)
	}

	return productTypes, nil
}


func (this pg) GetProductTypeById(id int64) (ProductType, error){
	var productType ProductType
	row := this.db.QueryRow("SELECT * FROM products WHERE id = $1 LIMIT 1;", id)
	if row.Err() != nil{
		return productType, row.Err()
	}
	row.Scan(&productType.ID, &productType.Name, &productType.Detail)
	return productType, nil
}

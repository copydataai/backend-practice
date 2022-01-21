package libs

import (
	"log"
	"time"
)

type Reviews struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	User       Users    `json:"author"`
	Product    Products `json:"product"`
	Created_at time.Time
}

type Users struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Email      *string `json:"email"`
	Password   string  ` json:"password"`
	Country    string  `json:"country"`
	Speciality string
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
}
type Products struct {
	ID            uint        `json:"id"`
	Detail        ProductType `json:"detail"`
	Trademark     string      `json:"trademark"`
	Manufacturing string      `json:"manufacturing"`
	SKU           string      `json:"sku"`
	CreatedAt     time.Time
}

type ProductType struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Detail string `json:'detail'`
}

func (this pg) ListReviews() (reviews []Reviews, err error) {
	rows, err := this.db.Query("SELECT * FROM reviews;")
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var review Reviews
		rows.Scan(&review.ID, &review.Title, &review.Content, &review.User, &review.Product, &review.Created_at)
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (this pg) ListProductTypes() (productTypes []ProductType, err error) {
	rows, err := this.db.Query("SELECT * FROM product_type;")
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var productType ProductType
		rows.Scan(&productType.ID, &productType.Name, &productType.Detail)
		productTypes = append(productTypes, productType)
	}

	return productTypes, nil
}

func (this pg) ListProducts() (products []Products, err error) {
	rows, err := this.db.Query("SELECT * FROM products;")
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var product Products
		rows.Scan(&product.ID, &product.Detail, &product.Trademark, &product.Manufacturing, &product.SKU, &product.CreatedAt)
		products = append(products, product)
	}

	return products, nil
}

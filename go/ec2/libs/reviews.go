package libs

import (
	"time"
)

type Reviews struct {
	ID         uint     `json:"id"`
	Title      string   `json:"title"`
	Content    string   `json:"content"`
	User       Users    `json:"author"`
	Product    Products `json:"product"`
	CreatedAt time.Time `json:"created_at"`
}

type Users struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Speciality string    `json:"speciality"`
}

func (this pg) ListReviews() (reviews []Reviews, count int, err error) {
	rows, err := this.db.Query(`SELECT rev.id, rev.title, rev.content, users.id, users.name, users.speciality, prod.id, type.id, type.name, type.detail, prod.trademark, prod.manufacturing, prod."SKU", prod.created_at, rev.created_at FROM reviews as rev JOIN users on rev.user = users.id JOIN products as prod on rev.product = prod.id JOIN product_type as type on prod.id = type.id;`)
	if err != nil {
		return reviews, -1, err
	}

	for rows.Next() {
		count++
		var review Reviews
		rows.Scan(&review.ID, &review.Title, &review.Content, &review.User.ID, &review.User.Name, &review.User.Speciality, &review.Product.ID, &review.Product.Detail.ID, &review.Product.Detail.Name, &review.Product.Detail.Detail, &review.Product.Trademark, &review.Product.Manufacturing, &review.Product.SKU, &review.Product.CreatedAt, &review.CreatedAt)
		reviews = append(reviews, review)
	}

	return reviews, count, nil
}

func (this pg) GetReviewById(id int64) (Reviews, int, error) {
	var review Reviews
	row := this.db.QueryRow(`SELECT rev.id, rev.title, rev.content, users.id, users.name, users.speciality, prod.id, type.id, type.name, type.detail, prod.trademark, prod.manufacturing, prod."SKU", prod.created_at, rev.created_at FROM reviews as rev JOIN users on rev.user = users.id JOIN products as prod on rev.product = prod.id JOIN product_type as type on prod.id = type.id WHERE rev.id = $1 LIMIT 1;`, id)
	if row.Err() != nil {
		return review, -1, row.Err()
	}
	err := row.Scan(&review.ID, &review.Title, &review.Content, &review.User.ID, &review.User.Name, &review.User.Speciality, &review.Product.ID, &review.Product.Detail.ID, &review.Product.Detail.Name, &review.Product.Detail.Detail, &review.Product.Trademark, &review.Product.Manufacturing, &review.Product.SKU, &review.Product.CreatedAt, &review.CreatedAt)
	if err != nil {
		return Reviews{}, 0, err
	}
	return review, 1, nil
}

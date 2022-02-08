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
	Email      *string   `json:"email"`
	Password   string    `json:"password"`
	Country    string    `json:"country"`
	Speciality string    `json:"speciality"`
	Role       string    `json:"role"`
	CreatedAt  time.Time `json:"created_at"`
}

func (this pg) ListReviews() (reviews []Reviews, err error) {
	rows, err := this.db.Query("SELECT * FROM reviews;")
	if err != nil {
		return reviews, err
	}

	for rows.Next() {
		var review Reviews
		rows.Scan(&review.ID, &review.Title, &review.Content, &review.User, &review.Product, &review.CreatedAt)
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (this pg) GetReviewById(id int64) (Reviews, error) {
	var review Reviews
	row := this.db.QueryRow("SELECT * FROM reviews WHERE id = $1 LIMIT 1;", id)
	if row.Err() == nil {
		return review, row.Err()
	}
	row.Scan(&review.ID, &review.Title, &review.Content, &review.User, &review.Product, &review.Product, &review.CreatedAt)
	return review, nil
}

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

func (this pg) ListReviews() (reviews []Reviews, count int, err error) {
	rows, err := this.db.Query("SELECT * FROM reviews LIMIT 5;")
	if err != nil {
		return reviews, -1, err
	}

	for rows.Next() {
		count++
		var review Reviews
		rows.Scan(&review.ID, &review.Title, &review.Content, &review.User, &review.Product, &review.CreatedAt)
		reviews = append(reviews, review)
	}

	return reviews, count, nil
}

func (this pg) GetReviewById(id int64) (Reviews, int, error) {
	var review Reviews
	row := this.db.QueryRow("SELECT * FROM reviews WHERE id = $1 LIMIT 1;", id)
	if row.Err() != nil {
		return review, -1, row.Err()
	}
	err := row.Scan(&review.ID, &review.Title, &review.Content, &review.User, &review.Product, &review.Product, &review.CreatedAt)
	if err != nil {
		return Reviews{}, -1, err
	}
	return review, 1, nil
}

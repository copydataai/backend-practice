package libs

type Posts interface {
	ListReviews() ([]Reviews, error)
	ListProducts() ([]Products, error)
	ListProductTypes() ([]ProductType, error)
}

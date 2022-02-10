package libs

type Posts interface {
	// Reviews
	ListReviews() ([]Reviews, int, error)
	GetReviewById(int64) (Reviews, int, error)

	// Products
	ListProducts() ([]Products, int, error)
	GetProductById(int64) (Products, int, error)

	// ProductTypes
	ListProductTypes() ([]ProductType, int, error)
	GetProductTypeById(int64) (ProductType, int, error)
}

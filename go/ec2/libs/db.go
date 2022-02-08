package libs

type Posts interface {
	// Reviews
	ListReviews() ([]Reviews, error)
	GetReviewById(int64) (Reviews, error)

	// Products
	ListProducts() ([]Products, error)
	GetProductById(int64) (Products, error)

	// ProductTypes
	ListProductTypes() ([]ProductType, error)
	GetProductTypeById(int64) (ProductType, error)
}

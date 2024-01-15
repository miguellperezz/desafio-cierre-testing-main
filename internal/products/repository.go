package products

import "errors"

var (
	ErrSellerIDRequired = errors.New("seller_id query param is required")
)

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	prodList := []Product{}

	if sellerID == "" {
		return prodList, ErrSellerIDRequired
	}

	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}

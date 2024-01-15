package products_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/stretchr/testify/assert"
)

func TestService_GetAllBySeller(t *testing.T) {
	t.Run("should return a list of products", func(t *testing.T) {
		// arrange
		var (
			expectedProducts = []products.Product{
				{ID: "mock", SellerID: "FEX112AC", Description: "generic product", Price: 123.55},
			}

			givenID = "FEX112AC"
		)

		repo := products.NewRepository()
		service := products.NewService(repo)

		// act
		products, err := service.GetAllBySeller(givenID)
		// assert
		assert.Equal(t, expectedProducts, products)
		assert.Equal(t, nil, err)
	})

	t.Run("should return an error if the seller ID is not provided", func(t *testing.T) {
		// arrange
		var (
			givenID = ""
		)

		repo := products.NewRepository()
		service := products.NewService(repo)

		// act
		_, err := service.GetAllBySeller(givenID)

		// assert

		//assert.Equal(t, nil, ps)
		assert.Equal(t, products.ErrSellerIDRequired, err)
	})

}

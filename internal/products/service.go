package products

import (
	"fmt"
)

type Service interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAllBySeller(sellerID string) ([]Product, error) {
	data, err := s.repo.GetAllBySeller(sellerID)

	if err != nil {
		fmt.Println("error in repository", err.Error(), "sellerId:", sellerID)
		return nil, err
	}
	return data, err
}

package service

import (
	"go-ecommerce/internal/model"
	"go-ecommerce/internal/repository"
	"strconv"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) ProductService {
	return ProductService{repo: repo}
}

func (s *ProductService) Create(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.repo.GetAll()
}
func (s *ProductService) FindByID(id uint) (*model.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) Update(idStr string, updated *model.Product) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	product, err := s.repo.FindByID(uint(id))
	if err != nil {
		return err
	}

	// Cập nhật thông tin
	product.Name = updated.Name
	product.Price = updated.Price
	product.Stock = updated.Stock

	return s.repo.Update(product)
}

func (s *ProductService) Delete(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	return s.repo.Delete(uint(id))
}
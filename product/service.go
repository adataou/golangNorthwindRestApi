package product

import (
	"golangNorthwindRestApi/helper"
)

type Service interface {
	GetProducts(param *getProductsRequest) (*ProductList, error)
	GetProductById(param *getProductByIDRequest) (*Product, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
	UpdateProduct(params *updateProductRequest) (int64, error)
	DeleteProduct(params *deleteProductRequest) (int64, error)
	GetBestSellers() (*ProductTopResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProducts(param *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(param)
	helper.Catch(err)
	totalProducts, err := s.repo.GetTotalProducts()
	helper.Catch(err)
	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}

func (s *service) GetProductById(param *getProductByIDRequest) (*Product, error) {
	//Business Logic
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) UpdateProduct(params *updateProductRequest) (int64, error) {
	return s.repo.UpdateProduct(params)
}

func (s *service) DeleteProduct(params *deleteProductRequest) (int64, error) {
	return s.repo.DeleteProduct(params.ProductID)
}

func (s *service) GetBestSellers() (*ProductTopResponse, error) {
	products, err := s.repo.GetBestSellers()
	helper.Catch(err)
	totalVentas, err := s.repo.GetTotalVentas()
	helper.Catch(err)

	return &ProductTopResponse{Data: products, TotalVentas: totalVentas}, err
}

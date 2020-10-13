package product

import (
	"context"
	"golangNorthwindRestApi/helper"

	"github.com/go-kit/kit/endpoint"
)

type getProductsRequest struct {
	Limit  int
	Offset int
}

type getProductByIDRequest struct {
	ProductID int
}

type getAddProductRequest struct {
	Category     string
	Description  string
	ListPrice    string
	StandardCost string
	ProductCode  string
	ProductName  string
}

type updateProductRequest struct {
	ID           int64
	Category     string
	Description  string
	ListPrice    float64
	StandardCost float64
	ProductCode  string
	ProductName  string
}

type deleteProductRequest struct {
	ProductID int
}

type getBestSellersRequest struct{}

// @Summary List of the  Products
// @Tags Product
// @Accept json
// @Produce  json
// @Param request body product.getProductsRequest true "User Data"
// @Success 200 {object} product.ProductList "ok"
// @Router /products/paginated [post]
func makeGetProductsEndPoint(s Service) endpoint.Endpoint {
	getProductsEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductsRequest)
		result, err := s.GetProducts(&req)
		helper.Catch(err)
		return result, nil
	}
	return getProductsEndpoint
}

//@Summary  Product by id
//@Tags Product
//@Accept json
//@Produce json
//@Param id path int true "Product Id"
//@Success 200 {object} product.Product "ok"
//@Router /products/{id} [get]
func makeGetProductByIdEndPoint(s Service) endpoint.Endpoint {
	getProductByIdEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProductByIDRequest)
		product, err := s.GetProductById(&req)
		helper.Catch(err)
		return product, nil
	}
	return getProductByIdEndpoint
}

//@Summary insert Product
//@Tags Product
//@Accept json
//@Produce json
//@Param request body product.getAddProductRequest true "User Data"
//@Success 200 {integer} int "ok"
//@Router /products/ [post]
func makeAddProductEndpoint(s Service) endpoint.Endpoint {
	addProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddProductRequest)
		id, err := s.InsertProduct(&req)
		helper.Catch(err)
		return id, nil
	}
	return addProductEndpoint
}

//@Summary update Product
//@Tags Product
//@Accept json
//@Produce json
//@Param request body product.updateProductRequest true "User Data"
//@Success 200 {interger} int "ok"
//@Router /products/ [put]
func makeUpdateProductEndpoint(s Service) endpoint.Endpoint {
	updateProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(updateProductRequest)
		id, err := s.UpdateProduct(&req)
		helper.Catch(err)
		return id, err
	}
	return updateProductEndpoint
}

//@Summary delete Product
//@Tags Product
//@Accept json
//@Produce json
//@Param id path int true "Product Id"
//@Success 200 {integer} int "ok"
//@Router /products/ [delete]
func makeDeleteProductEndpoint(s Service) endpoint.Endpoint {
	deleteProductEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductRequest)
		id, err := s.DeleteProduct(&req)
		helper.Catch(err)
		return id, err
	}
	return deleteProductEndpoint
}

// @Summary Best Sellers
// @Tags Product
// @Accept json
// @Produce  json
// @Success 200 {object} product.ProductTopResponse "ok"
// @Router /products/bestSellers [get]
func makeGetBestSellersEndpoint(s Service) endpoint.Endpoint {
	getBestSellersEndpoint := func(ctx context.Context, request interface{}) (interface{}, error) {
		products, err := s.GetBestSellers()
		helper.Catch(err)
		return products, err
	}
	return getBestSellersEndpoint
}

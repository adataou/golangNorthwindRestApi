package product

import (
	"context"
	"encoding/json"
	"golangNorthwindRestApi/helper"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

func MakeHttpHandler(s Service) http.Handler {
	r := chi.NewRouter()

	getProductsHandler := kithttp.NewServer(makeGetProductsEndPoint(s),
		getProductsRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/paginated", getProductsHandler)

	getProductByIdHandler := kithttp.NewServer(makeGetProductByIdEndPoint(s),
		getProductByIdRequestDecoder, kithttp.EncodeJSONResponse)

	r.Method(http.MethodGet, "/{id}", getProductByIdHandler)

	addProductHandler := kithttp.NewServer(makeAddProductEndpoint(s),
		addProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPost, "/", addProductHandler)

	updateProductHandler := kithttp.NewServer(makeUpdateProductEndpoint(s),
		updateProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodPut, "/", updateProductHandler)

	deleteProductHandler := kithttp.NewServer(makeDeleteProductEndpoint(s),
		deleteProductRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodDelete, "/{id}", deleteProductHandler)

	getBestSellerHandler := kithttp.NewServer(makeGetBestSellersEndpoint(s),
		getBestSellersRequestDecoder, kithttp.EncodeJSONResponse)
	r.Method(http.MethodGet, "/bestsellers", getBestSellerHandler)

	return r
}

func getProductsRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func getProductByIdRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	return getProductByIDRequest{
		ProductID: productId,
	}, nil
}

func addProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := getAddProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func updateProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	request := updateProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	helper.Catch(err)
	return request, nil
}

func deleteProductRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	productId, _ := strconv.Atoi(chi.URLParam(r, "id"))
	request := deleteProductRequest{
		ProductID: productId,
	}
	return request, nil
}

func getBestSellersRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return &getBestSellersRequest{}, nil
}

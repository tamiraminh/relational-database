package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/product"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
)


type ProductHandler struct {
	ProductService     product.ProductService
}

func ProvideProductHandler(productService product.ProductService) ProductHandler {
	return ProductHandler{
		ProductService: productService,
	}
}

func (h *ProductHandler) Router(r chi.Router) {
	r.Route("/product", func(r chi.Router) {
		// r.Group(func(r chi.Router) {
		// 	r.Post("/", h.CreateProduct)
		// })
		
		r.Post("/", h.CreateProduct)


	})
}


func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestFormat product.ProductRequestFormat
	err := decoder.Decode(&requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}



	product, err := h.ProductService.Create(requestFormat)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusCreated, product)
}





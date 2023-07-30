package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/evermos/boilerplate-go/internal/domain/product"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
	"github.com/gofrs/uuid"
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
		r.Put("/{id}", h.UpdateProduct)
		r.Delete("/", h.SoftDeleteProduct)
		r.Delete("/hard", h.HardDeleteProduct)
		r.Get("/pagination", h.ReadPagination)
		r.Get("/statussorted", h.ReadStatusSorted)
		r.Get("/by", h.ReadBy)


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



func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idString := chi.URLParam(r, "id")
	id, err := uuid.FromString(idString)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	var requestFormat product.ProductRequestFormat
	err = decoder.Decode(&requestFormat)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}


	product, err := h.ProductService.Update(id, requestFormat)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, product)
}



func (h *ProductHandler) SoftDeleteProduct(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	userIdString := r.URL.Query().Get("userId")
	id, err := uuid.FromString(idString)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	userId, err := uuid.FromString(userIdString)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}


	foo, err := h.ProductService.SoftDelete(id, userId)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, foo)
}


func (h *ProductHandler) HardDeleteProduct(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	userIdString := r.URL.Query().Get("userId")
	id, err := uuid.FromString(idString)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	userId, err := uuid.FromString(userIdString)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}


	product, err := h.ProductService.HardDelete(id, userId)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) ReadPagination(w http.ResponseWriter, r *http.Request){
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		response.WithError(w, err)
		return
	}

	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		response.WithError(w, err)
		return
	}

	products, err := h.ProductService.ReadPagination(limit, page - 1)
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) ReadStatusSorted(w http.ResponseWriter, r *http.Request){
	products, err := h.ProductService.ReadStatusSorted()
	if err != nil {
		response.WithError(w, err)
		return
	}

	response.WithJSON(w, http.StatusOK, products)
}


func (h *ProductHandler) ReadBy(w http.ResponseWriter, r *http.Request){
	brandName := r.URL.Query().Get("brandName")
	productName := r.URL.Query().Get("productName")
	variantName := r.URL.Query().Get("variantName")
	status := r.URL.Query().Get("status")

	if len(r.URL.Query()) > 1 {
		response.WithError(w, errors.New("params must only 1"))
	}

	products := []product.ProductStatus{}
	var err error
	switch {
		case brandName != "":
			products, err = h.ProductService.ReadByBrandName(brandName)
			if err != nil {
				response.WithError(w, err)
				return
			}
		case productName != "":
			products, err = h.ProductService.ReadByProductName(productName)
			if err != nil {
				response.WithError(w, err)
				return
			}
		case variantName != "":
			products, err = h.ProductService.ReadByVariantName(variantName)
			if err != nil {
				response.WithError(w, err)
				return
			}
		case status != "":
			products, err = h.ProductService.ReadByStatus(status)
			if err != nil {
				response.WithError(w, err)
				return
			}
	}


	response.WithJSON(w, http.StatusOK, products)
}

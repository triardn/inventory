package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/triardn/inventory/model"
)

type ProductResponse struct {
	ID       uint64 `json:"id"`
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Quantity int64  `json:"quantity,omitempty"`
	Created  int64  `json:"created,omitempty"`
}

type ProductPayload struct {
	Sku      string `json:"sku"`
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
}

func (h *Handler) GetAllProduct(w http.ResponseWriter, r *http.Request) (hErr error) {
	products, err := h.Service.Product.GetAllProduct()
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	var productResponse []ProductResponse
	for _, product := range products {
		productTemp := ProductResponse{
			ID:       product.ID,
			Sku:      product.Sku,
			Name:     product.Name,
			Quantity: product.Quantity,
			Created:  product.Created, // TODO: change to date time
		}

		productResponse = append(productResponse, productTemp)
	}

	response := NewAPIResponse(productResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

func (h *Handler) GetProductDetail(w http.ResponseWriter, r *http.Request) (hErr error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	product, err := h.Service.Product.GetProductDetail(id)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	productResponse := ProductResponse{
		ID:       product.ID,
		Sku:      product.Sku,
		Name:     product.Name,
		Quantity: product.Quantity,
		Created:  product.Created, // TODO: change to date time
	}

	response := NewAPIResponse(productResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) (hErr error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	var request ProductPayload
	err = json.Unmarshal(body, &request)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	// validation
	if request.Sku == "" {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if request.Name == "" {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if request.Quantity == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	newProduct := model.Product{}
	newProduct.Sku = request.Sku
	newProduct.Name = request.Name
	newProduct.Quantity = request.Quantity

	product, err := h.Service.Product.CreateProduct(newProduct)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	productResponse := ProductResponse{
		ID:       product.ID,
		Sku:      product.Sku,
		Name:     product.Name,
		Quantity: product.Quantity,
		Created:  product.Created, // TODO: change to date time
	}

	response := NewAPIResponse(productResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) (hErr error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	product, err := h.Service.Product.GetProductDetail(id)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	payload := make(map[string]interface{})
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	err = h.Service.Product.UpdateProduct(&product, payload)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	productResponse := ProductResponse{
		ID:       product.ID,
		Sku:      product.Sku,
		Name:     product.Name,
		Quantity: product.Quantity,
		Created:  product.Created, // TODO: change to date time
	}

	response := NewAPIResponse(productResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

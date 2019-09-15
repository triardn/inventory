package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) GetAllProduct(w http.ResponseWriter, r *http.Request) (hErr error) {
	products, err := h.Service.Product.GetAllProduct()
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	response := NewAPIResponse(products, nil)
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

	response := NewAPIResponse(product, nil)
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

	response := NewAPIResponse(product, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

package handler

import (
	"encoding/json"
	"net/http"
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

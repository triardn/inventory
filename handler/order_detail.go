package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetAllOrderDetail(w http.ResponseWriter, r *http.Request) (hErr error) {
	orderDetails, err := h.Service.OrderDetail.GetAllOrderDetail()
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	var allOrderDetail []OrderDetail
	for _, detail := range orderDetails {
		tempDetail := OrderDetail{
			ID:       detail.ID,
			ItemSku:  detail.Product.Sku,
			ItemName: detail.Product.Name,
			Price:    detail.Price,
			Quantity: detail.Quantity,
			Total:    detail.Total,
			Notes:    detail.Order.Notes,
			Created:  detail.Created, // TODO: change to date time
		}

		allOrderDetail = append(allOrderDetail, tempDetail)
	}

	response := NewAPIResponse(allOrderDetail, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

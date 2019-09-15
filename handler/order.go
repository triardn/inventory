package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderResponse struct {
	ID          uint64        `json:"id"`
	Invoice     string        `json:"invoice"`
	Total       uint64        `json:"total"`
	Notes       string        `json:"notes"`
	Created     uint64        `json:"created"`
	OrderDetail []OrderDetail `json:"order_detail"`
}

type OrderDetail struct {
	ID       uint64 `json:"id"`
	ItemSku  string `json:"item_sku"`
	ItemName string `json:"item_name"`
	Price    uint64 `json:"price"`
	Quantity uint64 `json:"quantity"`
	Total    uint64 `json:"total"`
	Notes    string `json:"notes,omitempty"`
	Created  uint64 `json:"created,omitempty"`
}

func (h *Handler) GetAllOrder(w http.ResponseWriter, r *http.Request) (hErr error) {
	orders, err := h.Service.Order.GetAllOrder()
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	var orderResponse []OrderResponse
	for _, order := range orders {
		var detailTemp []OrderDetail
		details, err := h.Service.OrderDetail.GetDetailByOrderID(order.ID)
		if err != nil {
			return StatusError{Code: http.StatusInternalServerError, Err: err}
		}

		for _, detail := range details {
			product, err := h.Service.Product.GetProductDetail(detail.ProductID)
			if err != nil {
				return StatusError{Code: http.StatusInternalServerError, Err: err}
			}

			detail := OrderDetail{
				ID:       detail.ID,
				ItemSku:  product.Sku,
				ItemName: product.Name,
				Price:    detail.Price,
				Quantity: detail.Quantity,
				Total:    detail.Total,
			}

			detailTemp = append(detailTemp, detail)
		}

		allOrder := OrderResponse{
			ID:          order.ID,
			Invoice:     order.Invoice,
			Total:       order.Total,
			Notes:       order.Notes,
			Created:     order.Created, // TODO: change to date time
			OrderDetail: detailTemp,
		}

		orderResponse = append(orderResponse, allOrder)
	}

	response := NewAPIResponse(orderResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) (hErr error) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		id, err = h.Service.Order.GetOrderIDByInvoice(vars["id"])
		if err != nil {
			return StatusError{Code: http.StatusBadRequest, Err: err}
		}
	}

	order, err := h.Service.Order.GetOrderByID(id)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	details, err := h.Service.OrderDetail.GetDetailByOrderID(order.ID)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	var allDetail []OrderDetail
	for _, detail := range details {
		tempDetail := OrderDetail{
			ID:       detail.ID,
			ItemSku:  detail.Product.Sku,
			ItemName: detail.Product.Name,
			Price:    detail.Price,
			Quantity: detail.Quantity,
			Total:    detail.Total,
		}

		allDetail = append(allDetail, tempDetail)
	}

	orderResponse := OrderResponse{
		ID:          order.ID,
		Invoice:     order.Invoice,
		Total:       order.Total,
		Notes:       order.Notes,
		Created:     order.Created, // TODO: change to date time
		OrderDetail: allDetail,
	}

	response := NewAPIResponse(orderResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

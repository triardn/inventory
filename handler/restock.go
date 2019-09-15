package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/triardn/inventory/model"
)

type RestockPayload struct {
	ProductSku       string `json:"product_sku"`
	ProductName      string `json:"product_name"`
	OrderedQuantity  int64  `json:"ordered_quantity"`
	ReceivedQuantity int64  `json:"received_quantity"`
	Price            int64  `json:"price"`
	Total            int64  `json:"total"`
	ReceiptNumber    string `json:"receipt_number"`
	Notes            string `json:"notes"`
}

type RestockResponse struct {
	ID               uint64          `json:"id"`
	OrderedQuantity  int64           `json:"ordered_quantity"`
	ReceivedQuantity int64           `json:"received_quantity"`
	Price            int64           `json:"price"`
	Total            int64           `json:"total"`
	ReceiptNumber    string          `json:"receipt_number"`
	Notes            string          `json:"notes"`
	Created          int64           `json:"created"`
	Product          ProductResponse `json:"product"`
}

func (h *Handler) GetAllRestock(w http.ResponseWriter, r *http.Request) (hErr error) {
	restocks, err := h.Service.Restock.GetAllRestockData()
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	var restockResponse []RestockResponse
	for _, restock := range restocks {
		tempRestock := RestockResponse{
			ID:               restock.ID,
			OrderedQuantity:  restock.OrderedQuantity,
			ReceivedQuantity: restock.ReceivedQuantity,
			Price:            restock.Price,
			Total:            restock.Total,
			ReceiptNumber:    restock.ReceiptNumber,
			Notes:            restock.Notes,
			Created:          restock.Created, // TODO: change to date time
			Product: ProductResponse{
				ID:   restock.Product.ID,
				Sku:  restock.Product.Sku,
				Name: restock.Product.Name,
			},
		}

		restockResponse = append(restockResponse, tempRestock)
	}

	response := NewAPIResponse(restockResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

func (h *Handler) CreateRestockData(w http.ResponseWriter, r *http.Request) (hErr error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	var request RestockPayload
	err = json.Unmarshal(body, &request)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	// validation
	if request.ProductSku == "" {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if request.OrderedQuantity == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if request.ReceivedQuantity == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if request.Price == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if request.Total == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	if request.ReceiptNumber == "" {
		request.ReceiptNumber = "(Hilang)"
	}

	productID, err := h.Service.Product.GetProductIDBySKU(request.ProductSku)
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return StatusError{Code: http.StatusInternalServerError, Err: err}
		} else {
			if request.ProductName == "" {
				return StatusError{Code: http.StatusBadRequest, Err: err}
			}

			productPayload := model.Product{
				Sku:  request.ProductSku,
				Name: request.ProductName,
			}

			newProduct, err := h.Service.Product.CreateProduct(productPayload)
			if err != nil {
				return StatusError{Code: http.StatusInternalServerError, Err: err}
			}

			productID = newProduct.ID
		}
	}

	var newRestock model.Restock
	newRestock.ProductID = productID
	newRestock.OrderedQuantity = request.OrderedQuantity
	newRestock.ReceivedQuantity = request.ReceivedQuantity
	newRestock.Price = request.Price
	newRestock.Total = request.Total
	newRestock.ReceiptNumber = request.ReceiptNumber
	newRestock.Notes = request.Notes

	restock, err := h.Service.Restock.CreateRestockData(newRestock)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	restockResponse := RestockResponse{
		ID:               restock.ID,
		OrderedQuantity:  restock.OrderedQuantity,
		ReceivedQuantity: restock.ReceivedQuantity,
		Price:            restock.Price,
		Total:            restock.Total,
		ReceiptNumber:    restock.ReceiptNumber,
		Notes:            restock.Notes,
		Created:          restock.Created, // TODO: change to date time
		Product: ProductResponse{
			ID:   restock.Product.ID,
			Sku:  restock.Product.Sku,
			Name: restock.Product.Name,
		},
	}

	response := NewAPIResponse(restockResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

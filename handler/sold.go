package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/model"
)

type SoldPayload struct {
	ProductSku string `json:"product_sku"`
	Quantity   int64  `json:"quantity"`
	Price      int64  `json:"price"`
	Notes      string `json:"notes"`
}

type SoldResponse struct {
	ID           uint64          `json:"id"`
	QuantitySold int64           `json:"quantity_sold"`
	Price        int64           `json:"price"`
	Total        int64           `json:"total"`
	Notes        string          `json:"notes"`
	Created      int64           `json:"created"`
	Product      ProductResponse `json:"product"`
}

func (h *Handler) GetAllSoldProduct(w http.ResponseWriter, r *http.Request) (hErr error) {
	solds, err := h.Service.Sold.GetAllSoldProducts()
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	var soldResponse []SoldResponse
	for _, sold := range solds {
		tempSold := SoldResponse{
			ID:           sold.ID,
			QuantitySold: sold.Quantity,
			Price:        sold.Price,
			Total:        sold.Total,
			Notes:        sold.Notes,
			Created:      sold.Created, // TODO: change to date time
			Product: ProductResponse{
				ID:   sold.Product.ID,
				Sku:  sold.Product.Sku,
				Name: sold.Product.Name,
			},
		}

		soldResponse = append(soldResponse, tempSold)
	}

	response := NewAPIResponse(soldResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

func (h *Handler) CreateSoldProductData(w http.ResponseWriter, r *http.Request) (hErr error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	var request SoldPayload
	err = json.Unmarshal(body, &request)
	if err != nil {
		return StatusError{Code: http.StatusBadRequest, Err: err}
	}

	// validation
	if request.ProductSku == "" {
		return StatusError{Code: http.StatusBadRequest, Err: common.ErrProductSkuCannotNull}
	}

	if request.Quantity == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: common.ErrProductQuantityCannotNull}
	}

	if request.Price == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: common.ErrProductPriceCannotNull}
	}

	productID, err := h.Service.Product.GetProductIDBySKU(request.ProductSku)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return StatusError{Code: http.StatusBadRequest, Err: err}
		} else {
			return StatusError{Code: http.StatusInternalServerError, Err: err}
		}
	}

	var newSold model.Sold
	newSold.ProductID = productID
	newSold.Quantity = request.Quantity
	newSold.Price = request.Price
	newSold.Total = newSold.Quantity * newSold.Price
	newSold.Notes = request.Notes

	sold, err := h.Service.Sold.CreateSoldRecord(newSold)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	soldResponse := SoldResponse{
		ID:           sold.ID,
		QuantitySold: sold.Quantity,
		Price:        sold.Price,
		Total:        sold.Total,
		Notes:        sold.Notes,
		Created:      sold.Created, // TODO: change to date time
		Product: ProductResponse{
			ID:   sold.Product.ID,
			Sku:  sold.Product.Sku,
			Name: sold.Product.Name,
		},
	}

	response := NewAPIResponse(soldResponse, nil)
	resp, err := json.Marshal(response)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

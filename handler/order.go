package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/triardn/inventory/common"
)

type OrderResponse struct {
	ID          uint64        `json:"id"`
	Invoice     string        `json:"invoice"`
	Total       int64         `json:"total"`
	Notes       string        `json:"notes"`
	Created     int64         `json:"created"`
	OrderDetail []OrderDetail `json:"order_detail"`
}

type OrderDetail struct {
	ID       uint64 `json:"id"`
	ItemSku  string `json:"item_sku"`
	ItemName string `json:"item_name"`
	Price    int64  `json:"price"`
	Quantity int64  `json:"quantity"`
	Total    int64  `json:"total"`
	Notes    string `json:"notes,omitempty"`
	Created  int64  `json:"created,omitempty"`
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

func (h *Handler) ExportOrder(w http.ResponseWriter, r *http.Request) (hErr error) {
	dateStart := r.URL.Query().Get("start")
	dateEnd := r.URL.Query().Get("end")

	layout := "2006-01-02"

	start, _ := time.Parse(layout, dateStart)
	end, _ := time.Parse(layout, dateEnd)

	data, totalProfit, err := h.Service.Order.PopulateExportData(dateStart, dateEnd)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	stats := h.Service.Order.ReportStatistics(dateStart, dateEnd)

	rangeTime := fmt.Sprintf("%s - %s", start.Format("02 January 2006"), end.Format("02 January 2006"))

	header := [][]string{
		{"LAPORAN PENJUALAN"},
		{""},
		{"Tanggal Cetak", time.Now().Format("02 January 2006")},
		{"Tanggal", rangeTime},
		{"Total Omzet", common.FormatCurrency("id_ID", stats["turnover"], true)},
		{"Total Laba Kotor", common.FormatCurrency("id_ID", totalProfit, true)},
		{"Total Penjualan", strconv.FormatInt(stats["orderCount"], 10)},
		{"Total Barang", strconv.FormatInt(stats["productCount"], 10)},
		{""},
		{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"},
	}

	for _, d := range data {
		header = append(header, d)
	}

	fileName := "Laporan Penjualan - " + rangeTime + ".csv"
	csvfile, err := os.Create(fileName)
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)

	err = writer.WriteAll(header) // flush everything into csvfile
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	Openfile, err := os.Open(fileName)
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "text/csv")

	//Send the file
	io.Copy(w, Openfile)

	return nil
}

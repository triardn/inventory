package handler

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/triardn/inventory/common"
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
		return StatusError{Code: http.StatusBadRequest, Err: common.ErrProductSkuCannotNull}
	}

	if request.Name == "" {
		return StatusError{Code: http.StatusBadRequest, Err: common.ErrProductNameCannotNull}
	}

	if request.Quantity == 0 {
		return StatusError{Code: http.StatusBadRequest, Err: common.ErrProductQuantityCannotNull}
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

func (h *Handler) ExportProduct(w http.ResponseWriter, r *http.Request) (hErr error) {
	data, total, err := h.Service.Product.PopulateExportData()
	if err != nil {
		return StatusError{Code: http.StatusInternalServerError, Err: err}
	}

	statistics := h.Service.Product.ProductStatistics()

	header := [][]string{
		{"LAPORAN NILAI BARANG"},
		{""},
		{"Tanggal Cetak", time.Now().Format("02 January 2006")},
		{"Jumlah SKU", strconv.Itoa(statistics["sku"])},
		{"Jumlah Total Barang", strconv.Itoa(statistics["stock"])},
		{"Total Nilai", common.FormatCurrency("id_ID", total, true)},
		{""},
		{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"},
	}

	for _, d := range data {
		header = append(header, d)
	}

	fileName := "Laporan Nilai Barang - " + time.Now().Format("02 January 2006") + ".csv"
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

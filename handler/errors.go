package handler

import (
	"github.com/triardn/inventory/common"
)

var errorMapEnglish map[error]string
var errorMapBahasa map[error]string

func init() {
	errorMapEnglish = make(map[error]string)
	errorMapBahasa = make(map[error]string)

	errorMapEnglish[common.ErrInvalidRequestEntity] = "Invalid request entity"
	errorMapBahasa[common.ErrInvalidRequestEntity] = "Entitas permintaan tidak valid"

	errorMapEnglish[common.ErrMissingRequestEntity] = "Missing request entity"
	errorMapBahasa[common.ErrMissingRequestEntity] = "Entitas permintaan tidak ditemukan"

	errorMapEnglish[common.ErrProductSkuCannotNull] = "Product SKU cannot be null"
	errorMapBahasa[common.ErrProductSkuCannotNull] = "SKU produk tidak boleh kosong"

	errorMapEnglish[common.ErrProductNameCannotNull] = "Product name cannot be null"
	errorMapBahasa[common.ErrProductNameCannotNull] = "Nama produk tidak boleh kosong"

	errorMapEnglish[common.ErrProductQuantityCannotNull] = "Quantity product cannot be null"
	errorMapBahasa[common.ErrProductQuantityCannotNull] = "Jumlah produk tidak boleh kosong"

	errorMapEnglish[common.ErrProductPriceCannotNull] = "Product price cannot be null"
	errorMapBahasa[common.ErrProductPriceCannotNull] = "Haga produk tidak boleh kosong"
}

package common

import (
	"errors"
	"fmt"
)

type ThirdPartyError struct {
	HTTPErrorCode int
	ErrorCode     string
	Description   string
}

func (tpe ThirdPartyError) Error() string {
	return fmt.Sprintf("%d - %s - %s", tpe.HTTPErrorCode, tpe.ErrorCode, tpe.Description)
}

var ErrInvalidRequestEntity = errors.New("Invalid request entity")
var ErrMissingRequestEntity = errors.New("Missing request entity")

var ErrProductSkuCannotNull = errors.New("Product SKU cannot be null")
var ErrProductNameCannotNull = errors.New("Product name cannot be null")
var ErrProductQuantityCannotNull = errors.New("Quantity cannot be null")
var ErrProductPriceCannotNull = errors.New("Product price cannot be null")

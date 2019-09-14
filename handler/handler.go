package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/jinzhu/gorm"

	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/service"
	"github.com/triardn/inventory/version"
)

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code       int
	Err        error
	ThirdParty bool
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}

type HttpHandler struct {
	Logger *common.APILogger
	H      func(w http.ResponseWriter, r *http.Request) error
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//w.Header().Add("Access-Control-Allow-Origin", "*")
	//w.Header().Add("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")

	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	err := h.H(w, r)
	if err != nil {
		h.Logger.HandleErrorWithTrace(err)

		switch e := err.(type) {
		case Error:
			if ee, ok := e.(StatusError); ok {

				if gorm.IsRecordNotFoundError(ee.Err) {
					emptyResp := NewAPIResponse(nil, nil)

					resp, err := json.Marshal(emptyResp)
					if err != nil {
						WriteErrorResponse(w, ee)
						return
					}

					w.WriteHeader(http.StatusOK)
					w.Write(resp)

					return
				}

				WriteErrorResponse(w, ee)
			}
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

type APIErrorResponse struct {
	Errors []APIError `json:"errors"`
}
type APIError struct {
	Code    *int    `json:"code,omitempty"`
	Details Details `json:"details"`
}

type Details struct {
	Id string `json:"id"`
	En string `json:"en"`
}

type APIResponse struct {
	Next *string     `json:"next,omitempty"`
	Data interface{} `json:"data"`
	Meta Meta        `json:"meta"`
}

type Meta struct {
	Version string `json:"version"`
	Status  string `json:"api_status"`
	APIEnv  string `json:"api_env"`
}

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func NewAPIResponse(data interface{}, nextPage *string) APIResponse {
	var response APIResponse
	voData := reflect.ValueOf(data)
	arrayData := []interface{}{}
	if voData.Kind() != reflect.Slice {
		if voData.IsValid() {
			arrayData = []interface{}{data}
		}
		response.Data = arrayData
	} else {
		if voData.Len() != 0 {
			response.Data = data
		} else {
			response.Data = arrayData
		}
	}
	response.Meta.Version = version.Version
	response.Meta.APIEnv = version.Environment
	response.Meta.Status = "unstable"
	response.Next = nextPage
	return response
}

func NewAPIErrorResponse(apiErr ...APIError) APIErrorResponse {
	apiErrorResponse := APIErrorResponse{}
	apiErrorResponse.Errors = apiErr
	return apiErrorResponse
}

func GetRequestBody(r *http.Request) (reqBody []byte, err error) {
	reqBody, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, common.ErrInvalidRequestEntity
	}

	return
}

func ParseRequestBody(reqBody []byte, target interface{}) (err error) {
	err = json.Unmarshal(reqBody, target)

	if err != nil {
		return common.ErrMissingRequestEntity
	}

	return
}

func WriteErrorResponse(w http.ResponseWriter, err StatusError) {
	w.WriteHeader(err.Status())

	apiError := APIError{}

	if err.ThirdParty {
		apiError.Code = &err.Code
		apiError.Details = Details{
			Id: err.Err.Error(),
			En: err.Err.Error(),
		}
	} else if _, ok := errorMapEnglish[err.Err]; !ok {
		apiError.Details = Details{
			Id: "Ups ada kesalahan, silakan coba dalam beberapa saat lagi.",
			En: err.Err.Error(),
			// En: "Oops something went wrong, try again later."
		}
	} else {
		apiError.Details = Details{
			Id: errorMapBahasa[err.Err],
			En: errorMapEnglish[err.Err],
		}
	}

	apiErrResponse := NewAPIErrorResponse(apiError)
	errResp, _ := json.Marshal(apiErrResponse)
	w.Write(errResp)
}

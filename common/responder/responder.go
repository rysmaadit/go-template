package responder

import (
	"encoding/json"
	"net/http"

	pkgErrors "github.com/pkg/errors"
	"github.com/rysmaadit/go-template/common/errors"
	"github.com/sirupsen/logrus"
)

type Template struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}

func NewHttpResponse(r *http.Request, w http.ResponseWriter, httpCode int, result interface{}, err error) {
	if err != nil {
		Error(r, w, err, httpCode)
	} else {
		Success(w, result, httpCode)
	}
}

func Error(r *http.Request, w http.ResponseWriter, err error, httpCode int) {
	switch err := pkgErrors.Cause(err).(type) {
	case *errors.BadRequestError:
		badRequestError(r, w, err)
	case *errors.UnauthorizedError:
		unauthorizedError(r, w, err)
	default:
		GenericError(r, w, err, err.Error(), httpCode)
	}
}

func Success(w http.ResponseWriter, successResponse interface{}, responseCode ...int) {
	w.Header().Set("Content-Type", "application/json")

	if len(responseCode) == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(responseCode[0])
	}

	t := Template{
		Status: true,
		Result: successResponse,
		Error:  nil,
	}

	if successResponse != nil {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(t)
	}
}

func GenericError(r *http.Request, w http.ResponseWriter, err error, errorResponse interface{}, responseCode int) {
	log := logrus.WithFields(logrus.Fields{
		"Method": r.Method,
		"Host":   r.Host,
		"Path":   r.URL.Path,
	}).WithField("ResponseCode", responseCode)

	if responseCode < 500 {
		log.Warn(err.Error())
	} else {
		log.Error(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)

	t := Template{
		Status: false,
		Result: nil,
		Error:  errorResponse,
	}

	if errorResponse != nil {
		_ = json.NewEncoder(w).Encode(t)
	}
}

func badRequestError(r *http.Request, w http.ResponseWriter, err *errors.BadRequestError) {
	GenericError(r, w, err, err.Error(), http.StatusBadRequest)
}

func unauthorizedError(r *http.Request, w http.ResponseWriter, err *errors.UnauthorizedError) {
	GenericError(r, w, err, err.Error(), http.StatusUnauthorized)
}

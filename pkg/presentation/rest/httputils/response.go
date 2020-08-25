package httputils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"net/http"
)

var ErrJsonFormat = errors.New("invalid json format")
var ErrInvalidRequest = errors.New(http.StatusText(http.StatusBadRequest))
var ErrValidation = errors.Wrap(ErrInvalidRequest, "validation error")

type errorData struct {
	errors string
	status int
}

type payloadData struct {
	data   interface{}
	status int
}

func Error(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	var data = errorData{
		errors: err.Error(),
		status: code,
	}

	json.NewEncoder(w).Encode(data)
}

func Response(w http.ResponseWriter, payload interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if code == http.StatusNoContent {
		return
	}

	var data = payloadData{
		data: payload,
	}

	json.NewEncoder(w).Encode(data)
}

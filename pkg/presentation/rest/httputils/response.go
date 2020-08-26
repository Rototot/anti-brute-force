package httputils

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrJSONFormat     = errors.New("invalid json format")
	ErrInvalidRequest = errors.New(http.StatusText(http.StatusBadRequest))
	ErrValidation     = errors.Wrap(ErrInvalidRequest, "validation error")
)

type errorData struct {
	Errors string
	Status int
}

type payloadData struct {
	Data interface{}
}

func Error(w http.ResponseWriter, err error, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	data := errorData{
		Errors: err.Error(),
		Status: code,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		zap.S().Warn(err)
	}
}

func Response(w http.ResponseWriter, payload interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if code == http.StatusNoContent {
		return
	}

	data := payloadData{
		Data: payload,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		zap.S().Warn(err)
	}
}

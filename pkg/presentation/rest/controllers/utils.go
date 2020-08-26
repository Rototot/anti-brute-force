package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/httputils"
	"github.com/pkg/errors"
)

type grabber struct {
	validator StructValidator
}

func (g *grabber) grabBodyAndValidate(res http.ResponseWriter, req *http.Request, targetDto interface{}) error {
	if err := json.NewDecoder(req.Body).Decode(&targetDto); err != nil {
		httputils.Error(res, httputils.ErrJSONFormat, http.StatusBadRequest)

		return err
	}

	// validate
	if err := g.validator.Struct(targetDto); err != nil {
		httputils.Error(res, errors.Wrap(httputils.ErrValidation, err.Error()), http.StatusBadRequest)

		return err
	}

	return nil
}

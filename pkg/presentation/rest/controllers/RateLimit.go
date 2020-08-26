package controllers

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/Rototot/anti-brute-force/pkg/domain/constants"

	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/httputils"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
)

type attemptHandler interface {
	Execute(useCase usecases.CheckLoginAttempt) error
}

type ResetAttemptsHandler interface {
	Execute(useCase usecases.ResetLoginAttempts) error
}

type RateLimitController struct {
	grabber

	attemptHandler       attemptHandler
	resetAttemptsHandler ResetAttemptsHandler
}

func NewRateLimitController(validator StructValidator, attemptHandler attemptHandler, resetAttemptsHandler ResetAttemptsHandler) *RateLimitController {
	return &RateLimitController{grabber: grabber{validator: validator}, attemptHandler: attemptHandler, resetAttemptsHandler: resetAttemptsHandler}
}

func (c *RateLimitController) Attempt(res http.ResponseWriter, req *http.Request) {
	var query queries.LoginAttemptQuery

	// parse
	if err := c.grabber.grabBodyAndValidate(res, req, &query); err != nil {
		httputils.Error(res, httputils.ErrJSONFormat, http.StatusBadRequest)

		return
	}

	useCase := usecases.CheckLoginAttempt{
		Login:    query.Login,
		Password: query.Password,
		IP:       net.ParseIP(query.IP),
	}

	err := c.attemptHandler.Execute(useCase)

	if err == nil {
		httputils.Response(res, true, http.StatusOK)

		return
	} else if err == constants.ErrAccessDenied || err == constants.ErrAttemptsIsExceeded {
		httputils.Response(res, false, http.StatusOK)

		return
	}

	httputils.Error(res, err, http.StatusInternalServerError)
}

func (c *RateLimitController) Reset(res http.ResponseWriter, req *http.Request) {
	var query queries.ResetAttemptsQuery

	if err := json.NewDecoder(req.Body).Decode(&query); err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)

		return
	}

	// validate
	if err := c.validator.Struct(query); err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)

		return
	}

	useCase := usecases.ResetLoginAttempts{
		Login: query.Login,
		IP:    net.ParseIP(query.IP),
	}

	if err := c.resetAttemptsHandler.Execute(useCase); err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)

		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

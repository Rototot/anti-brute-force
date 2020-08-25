package controllers

import (
	"encoding/json"
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/httputils"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
	"net"
	"net/http"
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
		httputils.Error(res, httputils.ErrJsonFormat, http.StatusBadRequest)
		return
	}

	var useCase = usecases.CheckLoginAttempt{
		Login:    query.Login,
		Password: query.Password,
		IP:       net.ParseIP(query.IP),
	}

	if err := c.attemptHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)
		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
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

	var useCase = usecases.ResetLoginAttempts{
		Login: query.Login,
		IP:    net.ParseIP(query.IP),
	}

	if err := c.resetAttemptsHandler.Execute(useCase); err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

package controllers

import (
	"encoding/json"
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
	"net"
	"net/http"
)

type attemptHandler interface {
	Execute(useCase usecases.CheckLoginAttempt) error
}

type resetAttemptsHandler interface {
	Execute(useCase usecases.ResetLoginAttempts) error
}

type RateLimitController struct {
	attemptHandler       attemptHandler
	resetAttemptsHandler resetAttemptsHandler

	validator StructValidator
}

func (c *RateLimitController) Attempt(res http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	var query queries.LoginAttemptQuery

	err := decoder.Decode(&query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	// validate
	err = c.validator.Struct(query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	var useCase = usecases.CheckLoginAttempt{
		Login:    query.Login,
		Password: query.Password,
		IP:       net.ParseIP(query.IP),
	}

	err = c.attemptHandler.Execute(useCase)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
	}
}

func (c *RateLimitController) Reset(res http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	var query queries.ResetAttemptsQuery

	err := decoder.Decode(&query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	// validate
	err = c.validator.Struct(query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	var useCase = usecases.ResetLoginAttempts{
		Login: query.Login,
		IP:    net.ParseIP(query.IP),
	}

	err = c.resetAttemptsHandler.Execute(useCase)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
	}
}

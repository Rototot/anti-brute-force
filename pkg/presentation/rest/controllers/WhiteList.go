package controllers

import (
	"encoding/json"
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
	"net"
	"net/http"
)

type createWhiteListHandler interface {
	Execute(useCase usecases.AddIPToWhiteList) error
}

type removeWhiteListHandler interface {
	Execute(useCase usecases.RemoveIpFromWhiteList) error
}

type WhiteListCrudController struct {
	createHandler createWhiteListHandler
	removeHandler removeWhiteListHandler

	validator StructValidator
}

func (c *WhiteListCrudController) Index(res http.ResponseWriter, req *http.Request) {

}

func (c *WhiteListCrudController) Create(res http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	var query queries.CreateWhiteListQuery

	err := decoder.Decode(&query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	// validate
	err = c.validator.Struct(query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	_, network, err := net.ParseCIDR(query.Subnet)
	if err != nil || network == nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	var useCase = usecases.AddIPToWhiteList{
		Subnet: *network,
	}

	err = c.createHandler.Execute(useCase)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
	}
}

func (c *WhiteListCrudController) Delete(res http.ResponseWriter, req *http.Request) {
	var decoder = json.NewDecoder(req.Body)
	var query queries.DeleteWhiteListQuery

	err := decoder.Decode(&query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	// validate
	err = c.validator.Struct(query)
	if err != nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	_, network, err := net.ParseCIDR(query.Subnet)
	if err != nil || network == nil {
		http.Error(res, "Invalid query", http.StatusBadRequest)
	}

	var useCase = usecases.RemoveIpFromWhiteList{
		Subnet: *network,
	}

	err = c.removeHandler.Execute(useCase)
	if err != nil {
		http.Error(res, "internal server error", http.StatusInternalServerError)
	}
}

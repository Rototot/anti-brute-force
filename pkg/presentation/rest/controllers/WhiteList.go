package controllers

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/httputils"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
	"net"
	"net/http"
)

type CreateWhiteListHandler interface {
	Execute(useCase usecases.AddIPToWhiteList) error
}

type RemoveWhiteListHandler interface {
	Execute(useCase usecases.RemoveIpFromWhiteList) error
}

type WhiteListCrudController struct {
	grabber

	createHandler CreateWhiteListHandler
	removeHandler RemoveWhiteListHandler
}

func NewWhiteListCrudController(validator StructValidator, createHandler CreateWhiteListHandler, removeHandler RemoveWhiteListHandler) *WhiteListCrudController {
	return &WhiteListCrudController{grabber: grabber{validator: validator}, createHandler: createHandler, removeHandler: removeHandler}
}

func (c *WhiteListCrudController) Index(res http.ResponseWriter, req *http.Request) {

}

func (c *WhiteListCrudController) Create(res http.ResponseWriter, req *http.Request) {
	var query queries.CreateWhiteListQuery

	// parse
	if err := c.grabber.grabBodyAndValidate(res, req, &query); err != nil {
		return
	}

	_, network, err := net.ParseCIDR(query.Subnet)
	if err != nil || network == nil {
		httputils.Error(res, httputils.ErrValidation, http.StatusBadRequest)
		return
	}

	var useCase = usecases.AddIPToWhiteList{
		Subnet: *network,
	}

	if err := c.createHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)
		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

func (c *WhiteListCrudController) Delete(res http.ResponseWriter, req *http.Request) {
	var query queries.DeleteWhiteListQuery

	// parse
	if err := c.grabber.grabBodyAndValidate(res, req, &query); err != nil {
		return
	}

	_, network, err := net.ParseCIDR(query.Subnet)
	if err != nil || network == nil {
		httputils.Error(res, httputils.ErrValidation, http.StatusBadRequest)
		return
	}

	var useCase = usecases.RemoveIpFromWhiteList{
		Subnet: *network,
	}

	if err := c.removeHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)
		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

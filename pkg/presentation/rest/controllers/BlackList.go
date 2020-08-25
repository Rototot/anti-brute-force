package controllers

import (
	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/httputils"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
	"net"
	"net/http"
)

type CreateBlackListHandler interface {
	Execute(useCase usecases.AddIpToBlacklist) error
}

type RemoveBlackListHandler interface {
	Execute(useCase usecases.RemoveIpFromBlackList) error
}

type BlackListCrudController struct {
	grabber

	createHandler CreateBlackListHandler
	removeHandler RemoveBlackListHandler
}

func NewBlackListCrudController(
	validator StructValidator,
	createHandler CreateBlackListHandler,
	removeHandler RemoveBlackListHandler,
) *BlackListCrudController {
	return &BlackListCrudController{grabber: grabber{validator: validator}, createHandler: createHandler, removeHandler: removeHandler}
}

func (c *BlackListCrudController) Index(res http.ResponseWriter, req *http.Request) {

}

func (c *BlackListCrudController) Create(res http.ResponseWriter, req *http.Request) {
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

	var useCase = usecases.AddIpToBlacklist{Subnet: *network}

	if err := c.createHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)
		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

func (c *BlackListCrudController) Delete(res http.ResponseWriter, req *http.Request) {
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

	var useCase = usecases.RemoveIpFromBlackList{Subnet: *network}

	if err := c.removeHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)
		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

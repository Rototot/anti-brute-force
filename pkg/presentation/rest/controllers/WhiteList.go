package controllers //nolint:dupl

import (
	"net"
	"net/http"

	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/httputils"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
)

type CreateWhiteListHandler interface {
	Execute(useCase usecases.AddIPToWhiteList) error
}

type RemoveWhiteListHandler interface {
	Execute(useCase usecases.RemoveIPFromWhiteList) error
}

type WhiteListCrudController struct {
	grabber

	createHandler CreateWhiteListHandler
	removeHandler RemoveWhiteListHandler
}

func NewWhiteListCrudController(
	validator StructValidator,
	createHandler CreateWhiteListHandler,
	removeHandler RemoveWhiteListHandler,
) *WhiteListCrudController {
	return &WhiteListCrudController{
		grabber:       grabber{validator: validator},
		createHandler: createHandler,
		removeHandler: removeHandler,
	}
}

func (c *WhiteListCrudController) ListWhitelists(res http.ResponseWriter, req *http.Request) {
}

func (c *WhiteListCrudController) CreateWhitelist(res http.ResponseWriter, req *http.Request) {
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

	useCase := usecases.AddIPToWhiteList{
		Subnet: *network,
	}

	if err := c.createHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)

		return
	}

	httputils.Response(res, nil, http.StatusCreated)
}

func (c *WhiteListCrudController) DeleteWhitelist(res http.ResponseWriter, req *http.Request) {
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

	useCase := usecases.RemoveIPFromWhiteList{
		Subnet: *network,
	}

	if err := c.removeHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)

		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

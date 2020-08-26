package controllers //nolint:dupl

import (
	"net"
	"net/http"

	"github.com/Rototot/anti-brute-force/pkg/application/usecases"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/httputils"
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/queries"
)

type CreateBlackListHandler interface {
	Execute(useCase usecases.AddIPToBlacklist) error
}

type RemoveBlackListHandler interface {
	Execute(useCase usecases.RemoveIPFromBlackList) error
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
	return &BlackListCrudController{
		grabber:       grabber{validator: validator},
		createHandler: createHandler,
		removeHandler: removeHandler,
	}
}

func (c *BlackListCrudController) ListBlacklists(res http.ResponseWriter, req *http.Request) {
}

func (c *BlackListCrudController) CreateBlacklist(res http.ResponseWriter, req *http.Request) {
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

	useCase := usecases.AddIPToBlacklist{Subnet: *network}

	if err := c.createHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)

		return
	}

	httputils.Response(res, nil, http.StatusCreated)
}

func (c *BlackListCrudController) DeleteBlacklist(res http.ResponseWriter, req *http.Request) {
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

	useCase := usecases.RemoveIPFromBlackList{Subnet: *network}

	if err := c.removeHandler.Execute(useCase); err != nil {
		httputils.Error(res, err, http.StatusInternalServerError)

		return
	}

	httputils.Response(res, nil, http.StatusNoContent)
}

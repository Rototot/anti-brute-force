package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type whitelistCrudService interface {
	ListWhitelists(res http.ResponseWriter, req *http.Request)
	CreateWhitelist(res http.ResponseWriter, req *http.Request)
	DeleteWhitelist(res http.ResponseWriter, req *http.Request)
}

type blacklistCrudService interface {
	ListBlacklists(res http.ResponseWriter, req *http.Request)
	CreateBlacklist(res http.ResponseWriter, req *http.Request)
	DeleteBlacklist(res http.ResponseWriter, req *http.Request)
}

type rateLimiterService interface {
	Attempt(res http.ResponseWriter, req *http.Request)
	Reset(res http.ResponseWriter, req *http.Request)
}

type Router struct {
	whitelist   whitelistCrudService
	blacklist   blacklistCrudService
	rateLimiter rateLimiterService
}

func NewRouter(whitelist whitelistCrudService, blacklist blacklistCrudService, rateLimiter rateLimiterService) *Router {
	return &Router{whitelist: whitelist, blacklist: blacklist, rateLimiter: rateLimiter}
}

func (r *Router) Create() *mux.Router {
	muxRouter := mux.NewRouter()

	// whitelist
	muxRouter.HandleFunc("/whitelists", r.whitelist.ListWhitelists).
		Methods("GET")
	muxRouter.HandleFunc("/whitelist", r.whitelist.CreateWhitelist).
		Methods("POST")
	muxRouter.HandleFunc("/whitelist/{id:[0-9\\./]+}", r.whitelist.DeleteWhitelist).
		Methods("DELETE")

	// blacklist
	muxRouter.HandleFunc("/blacklists", r.blacklist.ListBlacklists).
		Methods("GET")
	muxRouter.HandleFunc("/blacklist", r.blacklist.CreateBlacklist).
		Methods("POST")
	muxRouter.HandleFunc("/blacklist/{id:[0-9\\./]+}", r.blacklist.DeleteBlacklist).
		Methods("DELETE")

	// rate limiter
	muxRouter.HandleFunc("/login/attempt", r.rateLimiter.Attempt).
		Methods("POST")

	muxRouter.HandleFunc("/login/attempt", r.rateLimiter.Reset).
		Methods("DELETE")

	return muxRouter
}

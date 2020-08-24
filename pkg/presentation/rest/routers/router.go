package routers

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ipCrudService interface {
	Index(res http.ResponseWriter, req *http.Request)
	Create(res http.ResponseWriter, req *http.Request)
	Delete(res http.ResponseWriter, req *http.Request)
}

type rateLimiterService interface {
	Attempt(res http.ResponseWriter, req *http.Request)
	Reset(res http.ResponseWriter, req *http.Request)
}

type Router struct {
	whitelist   ipCrudService
	blacklist   ipCrudService
	rateLimiter rateLimiterService
}

func NewRouter(whitelist ipCrudService, blacklist ipCrudService, rateLimiter rateLimiterService) *Router {
	return &Router{whitelist: whitelist, blacklist: blacklist, rateLimiter: rateLimiter}
}
func (r *Router) Create() *mux.Router {
	muxRouter := mux.NewRouter()

	// whitelist
	muxRouter.HandleFunc("/whitelists", r.whitelist.Index).
		Methods("GET")
	muxRouter.HandleFunc("/whitelist", r.whitelist.Create).
		Methods("POST")
	muxRouter.HandleFunc("/whitelist/{id:[0-9\\./]+}", r.whitelist.Delete).
		Methods("DELETE")

	// blacklist
	muxRouter.HandleFunc("/blacklists", r.blacklist.Index).
		Methods("GET")
	muxRouter.HandleFunc("/blacklist", r.blacklist.Create).
		Methods("POST")
	muxRouter.HandleFunc("/blacklist/{id:[0-9\\./]+}", r.blacklist.Delete).
		Methods("DELETE")

	// rate limiter
	muxRouter.HandleFunc("/login/attempt", r.rateLimiter.Attempt).
		Methods("POST")

	muxRouter.HandleFunc("/login/attempt", r.rateLimiter.Reset).
		Methods("DELETE")

	return muxRouter
}

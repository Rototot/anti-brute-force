package routers

import (
	"github.com/Rototot/anti-brute-force/pkg/presentation/rest/controllers"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	// whitelist
	whitelist := &controllers.WhiteListCrudController{}
	r.HandleFunc("/whitelists", whitelist.Index).
		Methods("GET")
	r.HandleFunc("/whitelist", whitelist.Create).
		Methods("POST")
	r.HandleFunc("/whitelist/{id:[0-9\\./]+}", whitelist.Delete).
		Methods("DELETE")

	// blacklist
	blacklist := &controllers.BlackListCrudController{}
	r.HandleFunc("/blacklists", blacklist.Index).
		Methods("GET")
	r.HandleFunc("/blacklist", blacklist.Create).
		Methods("POST")
	r.HandleFunc("/blacklist/{id:[0-9\\./]+}", blacklist.Delete).
		Methods("DELETE")

	// rate limiter
	ratelimit := &controllers.RateLimitController{}
	r.HandleFunc("/login/attempt", ratelimit.Attempt).
		Methods("POST")

	r.HandleFunc("/login/attempt", ratelimit.Reset).
		Methods("DELETE")

	return r
}

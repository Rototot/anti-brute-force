package controllers

import "net/http"

type RateLimitController struct {
}

func (c *RateLimitController) Attempt(res http.ResponseWriter, req *http.Request) {

}

func (c *RateLimitController) Reset(res http.ResponseWriter, req *http.Request) {

}

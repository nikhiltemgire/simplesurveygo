package servicehandlers

import (
	"fmt"
	"net/http"
	"time"
)

type HttpServiceHandler interface {
	Get(*http.Request) SrvcRes
	Put(*http.Request) SrvcRes
	Post(*http.Request) SrvcRes
}

func methodRouter(p HttpServiceHandler, w http.ResponseWriter, r *http.Request) interface{} {
	var response interface{}

	if r.Method == "GET" {
		start := time.Now()
		response = p.Get(r)
		elapsed := time.Since(start)
		fmt.Println(elapsed)

	} else if r.Method == "PUT" {
		start := time.Now()
		response = p.Put(r)
		elapsed := time.Since(start)
		fmt.Println(elapsed)

	} else if r.Method == "POST" {
		start := time.Now()
		response = p.Post(r)
		elapsed := time.Since(start)
		fmt.Println(elapsed)

	}
	return response
}

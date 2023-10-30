package server

import (
	"net/http"

	"k8s.io/klog/v2"
)


func handleWrapChain(handle http.HandlerFunc) http.HandlerFunc {
	var final http.HandlerFunc

	final = rejectNotJson(handle)
	final = logRequest(handle)

	return final
}

func logRequest(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		klog.InfoS("revice request", "uri", r.RequestURI)
		handle(w, r)
	}
}

func rejectNotJson(handle http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if r.Header.Get()
		handle(w, r)
	}
}
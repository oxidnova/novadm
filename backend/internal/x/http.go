package x

import "net/http"

var offers = []string{"text/html", "text/*", "*/*", "application/json"}
var defaultOffer = "text/html"

func IsJSONRequest(r *http.Request) bool {
	return NegotiateContentType(r, offers, defaultOffer) == "application/json" ||
		r.Header.Get("Content-Type") == "application/json"
}

func IsBrowserRequest(r *http.Request) bool {
	return NegotiateContentType(r, offers, defaultOffer) == "text/html"
}

func AcceptsJSON(r *http.Request) bool {
	return NegotiateContentType(r, []string{
		"text/html",
		"application/json",
	}, "text/html") == "application/json"
}

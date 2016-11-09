// Package httputil provides http-related utilities
package httputil

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Echo encapsulates http request headers, form values and body, if any, and formats them into a string
func Echo(r *http.Request) string {
	e := fmt.Sprintf("%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		e += fmt.Sprintf("Header[%q] = %q\n", k, v)
	}
	e += fmt.Sprintf("Host = %q\n", r.Host)
	e += fmt.Sprintf("RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	e += "\n"
	for k, v := range r.Form {
		e += fmt.Sprintf("Form[%q] = %q\n", k, v)
	}
	e += "\n"

	if body, _ := ioutil.ReadAll(r.Body); len(body) != 0 {
		e += fmt.Sprintf("r.Body = %s\n", body)
	}
	return e
}

package httputil

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Echo(r *http.Request) string {
	var e string
	e += fmt.Sprintf("%s %s %s\n", r.Method, r.URL, r.Proto)
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
	body, _ := ioutil.ReadAll(r.Body)
	e += fmt.Sprintf("r.Body = %s\n", body)
	return e
}

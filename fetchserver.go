package gae

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/urlfetch"
	"io"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "url is :%t", req.TLS != nil)
}
func fetch(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	url := fmt.Sprintf("%s//%s/%s", "", "", "")
	resp, err := client.Get(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}

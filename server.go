package main

import (
	// "crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct{}

func copyHeaders(dst, src http.Header) {
	for k, _ := range dst {
		dst.Del(k)
	}
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Printf("Request url: %s", req.RequestURI)
	client := &http.Client{}

	request, err := http.NewRequest(req.Method, req.RequestURI, nil)
	if err != nil {
		fmt.Println(err)
	}
	// copyHeaders(request.Header, req.Header)
	for _, k := range []string{"Referer", "Cookie"} {
		request.Header[k] = req.Header[k]
	}
	resp, err := client.Do(request)

	if err == nil {
		defer resp.Body.Close()
		// copy header

		copyHeaders(w.Header(), resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
	} else {
		fmt.Println(err)
	}
}

func main() {
	server := http.Server{
		Addr:    ":8888",
		Handler: &handler{},
	}
	log.Printf("Server is running on: %s", "8888")
	server.ListenAndServe()
}

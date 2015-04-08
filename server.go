package main

import (
	// "crypto/tls"
	"fmt"
	"github.com/elazarl/goproxy"
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
	// fmt.Println(req.RequestURI)
	// var resp *http.Response
	// var err error
	// if req.TLS != nil || req.Method == "CONNECT" {
	// 	fmt.Println(req.Method)
	// 	w.Write([]byte("HTTP/1.0 200 OK\r\n\r\n"))
	// 	return
	// 	tr := &http.Transport{
	// 		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	// 		DisableCompression: true,
	// 	}
	// 	client := &http.Client{Transport: tr}
	// 	fmt.Println(req.Method, req.RequestURI)
	// 	request, err := http.NewRequest(req.Method, "https://"+req.RequestURI, nil)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	resp, err = client.Do(request)
	// 	resp.Write(w)
	// 	return
	// } else {
	log.Printf("Request url: %s", req.RequestURI)
	client := &http.Client{}

	request, err := http.NewRequest(req.Method, req.RequestURI, nil)
	if err != nil {
		fmt.Println(err)
	}
	// copyHeaders(request.Header, req.Header)

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

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe(":8888", proxy))
}

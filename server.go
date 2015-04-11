package main

import (
	// "crypto/tls"
	// "bufio"
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
	log.Printf("Request url: %s, method: %s", req.RequestURI, req.Method)
	client := &http.Client{}

	url := req.RequestURI
	// if req.Method == "CONNECT" || true {
	// 	url = "https://" + url

	// 	// w.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	// 	hj, ok := w.(http.Hijacker)
	// 	if !ok {
	// 		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
	// 		return
	// 	}

	// 	_, bufrw, err := hj.Hijack()
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	// Don't forget to close the connection:
	// 	// defer conn.Close()
	// 	//conn.Write([]byte("HTTP/1.0 200 OK\r\n\r\n"))
	// 	bufrw.WriteString("HTTP/1.0 200 OK\r\n\r\nhello world")

	// 	bufrw.Flush()
	// 	// go func(bufrw *bufio.ReadWriter) {
	// 	// 	for {
	// 	// 		fmt.Fprintf(bufrw, "You said: %q\nBye.\n", "hello world")
	// 	// 		bufrw.Flush()
	// 	// 		// s, err := bufrw.ReadString()
	// 	// 		// if err != nil {
	// 	// 		// 	log.Printf("error reading string: %v", err)
	// 	// 		// 	return
	// 	// 		// }
	// 	// 	}
	// 	// }(bufrw)
	// 	return
	// }
	request, err := http.NewRequest(req.Method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	// copyHeaders(request.Header, req.Header)
	for _, k := range []string{"Referer", "Cookie"} {
		request.Header[k] = req.Header[k]
	}
	resp, err := client.Do(request)
	log.Println(resp)
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

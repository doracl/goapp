package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("Connect", "https://www.baidu.com", nil)
	if err != nil {
		log.Println("new request err:")
		log.Fatal(err)
	}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("client do err:")
		log.Fatal(err)
	}
	log.Println("request back")
	defer resp.Body.Close()
	log.Println(resp)

}

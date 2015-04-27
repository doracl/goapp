package main

import (
	"crypto/tls"
	// "crypto/x509"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

func to_tag(key, value string) string {
	return fmt.Sprintf("<%s>%s</%s>", key, value, key)
}

func send(data map[string]string) {

}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RequestURI)
	decoder := json.NewDecoder(req.Body)
	var t map[string]string
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	data := map[string]string{
		"mch_billno":   "billno",
		"mch_id":       "",
		"wxappid":      "",
		"nick_name":    "Test",
		"send_name":    "Test",
		"re_openid":    "",
		"total_amount": "0",
		"min_value":    "0",
		"max_value":    "0",
		"total_num":    "1",
		"wishing":      "1",
		"client_ip":    "192.168.29.154",
		"act_name":     "1",
		"remark":       "1",
		"nonce_str":    "d2asf1323242sdf1a",
	}
	for key, val := range t {
		if key == "key" {
			continue
		}
		data[key] = val
	}
	counter++
	data["mch_billno"] = fmt.Sprintf("%s%s%010d", data["mch_id"], time.Now().Format("20060102"), counter)

	cert, err := tls.LoadX509KeyPair("apiclient_cert.pem", "apiclient_key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	config = config

	// config.BuildNameToCertificate()

	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	query_string := ""
	for _, k := range keys {
		query_string += k + "=" + data[k] + "&"
	}
	query_string += "key=" + t["key"]

	sign := fmt.Sprintf("%x", md5.Sum([]byte(query_string)))
	sign = strings.ToUpper(sign)
	data["sign"] = sign

	xmlstring := "<xml>"
	for k := range data {
		xmlstring += to_tag(k, data[k])
	}
	xmlstring += "</xml>"
	transport := &http.Transport{TLSClientConfig: &config}
	client := &http.Client{Transport: transport}
	b := bytes.NewBufferString(xmlstring)
	response, err := client.Post("https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", "application/xml", b)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	log.Print(string(content))
}

var counter = 0

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

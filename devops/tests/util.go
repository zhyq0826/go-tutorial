package tests

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//formatJSON and print beautiful
func formatJSON(f interface{}) {

	if v, ok := f.([]interface{}); ok {
		for _, vv := range v {
			formatJSON(vv)
		}
		return
	}

	for k, v := range f.(map[string]interface{}) {
		switch vv := v.(type) {
		case int:
			fmt.Println(k, "==>", vv)
		case string:
			fmt.Println(k, "==>", vv)
		case float64:
			fmt.Println(k, "==>", vv)
		}
	}
}

// decode json data to list
func decodeList(f interface{}) []map[string]interface{} {
	ret := make([]map[string]interface{}, 0)
	if v, ok := f.([]interface{}); ok {
		for _, vv := range v {
			ret = append(ret, vv.(map[string]interface{}))
		}
	}

	return ret
}

// decode json data to map
func decodeMap(f interface{}) map[string]interface{} {
	var ret map[string]interface{}
	if v, ok := f.(map[string]interface{}); ok {
		ret = v
	}

	return ret
}

// RequestClient client for test
type RequestClient struct {
	HOST        string
	ContentType string
	Client      http.Client
	Prefix      string
}

// JoinURL composie url and host
func (client RequestClient) JoinURL(url string) string {
	if strings.Contains(url, "http") || strings.Contains(url, "https") {
		return url
	}

	return strings.Join([]string{client.HOST, client.Prefix, url}, "")
}

//Get client get
func (client RequestClient) Get(url string) *http.Response {
	resp, err := http.Get(client.JoinURL(url))
	if err != nil {
		log.Println(err)
		return nil
	}

	return resp
}

//Post client post
func (client RequestClient) Post(url string, body io.Reader) *http.Response {
	resp, err := http.Post(client.JoinURL(url), client.ContentType, body)
	if err != nil {
		log.Println(err)
		return nil
	}

	return resp
}

//Put client pust
func (client RequestClient) Put(url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(http.MethodPut, client.JoinURL(url), body)
	if err != nil {
		log.Println(err)
		return nil
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}

	return resp

}

//Delete client pust
func (client RequestClient) Delete(url string, body io.Reader) *http.Response {
	req, err := http.NewRequest(http.MethodDelete, client.JoinURL(url), body)
	if err != nil {
		log.Println(err)
		return nil
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		log.Println(err)
		return nil
	}

	return resp

}

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/zhyq0826/go-tutorial/devops/resources"
)

var host = "http://127.0.0.1:8000"

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

func TestList(t *testing.T) {
	resp, err := http.Get(host + "/domain/list")
	if err != nil {
		t.Error(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var v []resources.DomainForm
	json.Unmarshal(data, &v)
	for _, d := range v {
		if d.URL == "" || d.Name == "" {
			t.Fail()
		}
	}
}

func TestCreate(t *testing.T) {
	form := resources.DomainForm{
		Name:    "i am name",
		URL:     "i am url",
		Private: 1,
	}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp, _ := http.Post(host+"/domain/create", "application/json", body)
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestUpdate(t *testing.T) {
	form := resources.DomainForm{
		Name:    "i am name",
		URL:     "i am url",
		Private: 1,
	}
	client := http.Client{}
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	req, _ := http.NewRequest(http.MethodPut, host+"/domain/1", body)
	resp, _ := client.Do(req)
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, host+"/domain/10", io.Reader(nil))
	resp, _ := client.Do(req)
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

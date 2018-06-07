package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

var host = "http://127.0.0.1:8000"

func TestCreate(t *testing.T) {

}

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

func decodeList(f interface{}) []map[string]interface{} {
	ret := make([]map[string]interface{}, 0)
	if v, ok := f.([]interface{}); ok {
		for _, vv := range v {
			ret = append(ret, vv.(map[string]interface{}))
		}
	}

	return ret
}

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
	var v interface{}
	json.Unmarshal(data, &v)
	for _, value := range decodeList(v) {
		ret := decodeMap(value)
		fmt.Println(ret)
		_, ok := ret["Name"]
		if ok != true {
			t.Fail()
		}
		_, ok = ret["URL"]
		if ok != true {
			t.Fail()
		}
	}
}

func TestUpdate(t *testing.T) {

}

func TestDelete(t *testing.T) {

}

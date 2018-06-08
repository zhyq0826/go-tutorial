package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"testing"

	"github.com/zhyq0826/go-tutorial/devops/resources"
)

var CLIENT = RequestClient{
	HOST:        "http://127.0.0.1:8000",
	Prefix:      "/domain",
	ContentType: "application/json",
}

func TestList(t *testing.T) {
	url := "/list"
	resp := CLIENT.Get(url)
	if resp == nil {
		t.Fail()
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var v []*resources.DomainForm
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
	resp := CLIENT.Post("/create", body)
	if resp == nil {
		t.Fail()
	}
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
	data, _ := json.Marshal(form)
	body := bytes.NewBuffer(data)
	resp := CLIENT.Put("/1", body)
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	resp := CLIENT.Delete("/10", io.Reader(nil))
	if resp == nil {
		t.Fail()
	}
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

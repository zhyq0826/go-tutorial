package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
)

type Domain struct {
	Name    string `gorm:"column:name;type:varchar(256)"`
	URL     string `gorm:"column:url;type:varchar(256)"`
	Private uint   `gorm:"column:private;type:tinyint;not null;default:0"`
}

type DomainForm struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Private   uint      `json:"private"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func main() {
	// fmt.Println(reflect.ValueOf(Domain{}).Elem())
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Println(v.String())
	fmt.Println(v.Type())
	x := v.Interface()
	i := x.(int)
	fmt.Println(i)
	sv := reflect.ValueOf(&Domain{
		Name: "name domain",
	}).Elem()
	df := &DomainForm{}
	dv := reflect.ValueOf(df).Elem()
	for i := 0; i < sv.NumField(); i++ {
		value := sv.Field(i)
		name := sv.Type().Field(i).Name
		fmt.Println(name)
		fmt.Println(value)

		dvalue := dv.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}

	fmt.Println(df)
}

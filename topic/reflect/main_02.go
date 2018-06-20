package main

import (
	"fmt"
	"reflect"
	"time"
)

type Base struct {
	ID        int       `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Domain struct {
	Base
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
	sv := reflect.ValueOf(&Domain{
		Name:    "name domain",
		URL:     "url",
		Private: 1,
		Base: Base{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}).Elem()
	df := &DomainForm{}
	dv := reflect.ValueOf(df).Elem()
	for i := 0; i < sv.NumField(); i++ {
		// reflect.Value
		value := sv.Field(i)
		name := sv.Type().Field(i).Name
		fmt.Println(name, "===", value)
		if name == "Base" {
			for j := 0; j < value.NumField(); j++ {
				baseValue := value.Field(j)
				baseName := value.Type().Field(j).Name
				dvalue := dv.FieldByName(baseName)
				if dvalue.IsValid() == false {
					continue
				}

				dvalue.Set(baseValue)
			}
		}
		dvalue := dv.FieldByName(name)
		// 检查是否存在
		if dvalue.IsValid() == false {
			continue
		}
		//设置新值
		dvalue.Set(value)
	}

	fmt.Println(df)
}

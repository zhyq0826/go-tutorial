package devops

import (
	"log"
	"time"

	"github.com/zhyq0826/go-tutorial/devops/db"
	model "github.com/zhyq0826/go-tutorial/devops/models/devops"
)

// CreateDomain create
func CreateDomain(name, url string, private uint) {
	domain := model.Domain{
		Name:    name,
		URL:     url,
		Private: private,
	}
	now := time.Now()
	domain.UpdatedAt = now
	domain.CreatedAt = now
	db.DB.Create(&domain)
}

// UpdateDomain update
func UpdateDomain(id int, name, url string, private uint) {
	domain := model.Domain{
		BaseModel: model.BaseModel{
			ID: id,
		},
	}
	db.DB.Model(&domain).Updates(&model.Domain{Name: name, URL: url})
}

// QueryDomain domain list
func QueryDomain(page, limit int, name, url string) []model.Domain {
	if limit == 0 {
		limit = 10
	}
	domains := make([]model.Domain, limit)
	conditon := make([]interface{}, 2)
	query := ""
	if name != "" {
		query += "name = ?"
		conditon = append(conditon, name)
	}
	if url != "" {
		query += "and url = ?"
		conditon = append(conditon, url)
	}
	db.DB.Where(query, conditon...).Find(&domains)

	log.Fatal(domains)
	return domains
}

package devops

import (
	"fmt"
	"time"

	"github.com/zhyq0826/go-tutorial/devops/db"
	model "github.com/zhyq0826/go-tutorial/devops/models/devops"
	"github.com/zhyq0826/go-tutorial/devops/resources"
)

// CreateApp create
func CreateApp(form resources.AppForm) {
	service := model.App{
		Name:          form.Name,
		URL:           form.URL,
		Desc:          form.Desc,
		DeployDir:     form.DeployDir,
		RepositoryURL: form.RepositoryURL,
		MonitorURL:    form.MonitorURL,
	}
	now := time.Now()
	service.UpdatedAt = now
	service.CreatedAt = now
	fmt.Println(service)
	db.DB.Create(&service)
}

// UpdateApp update
func UpdateApp(id int, form resources.AppForm) {
	service := model.App{
		BaseModel: model.BaseModel{
			ID: id,
		},
	}
	db.DB.Model(&service).Updates(&model.App{
		Name:          form.Name,
		URL:           form.URL,
		Desc:          form.Desc,
		DeployDir:     form.DeployDir,
		RepositoryURL: form.RepositoryURL,
		MonitorURL:    form.MonitorURL,
	})
}

// QueryApp app list
func QueryApp(page, limit int, name, url string) []model.App {
	if limit == 0 {
		limit = 10
	}
	ret := make([]model.App, limit)
	conditon := make([]interface{}, 0)
	query := ""
	if name != "" {
		query += "name = ?"
		conditon = append(conditon, name)
	}
	if url != "" {
		query += "and url = ?"
		conditon = append(conditon, url)
	}
	db.DB.Where(query, conditon...).Limit(limit).Offset((page - 1) * limit).Find(&ret)

	return ret
}

// DeleteApp app delete
func DeleteApp(id int) {
	db.DB.Delete(&model.App{}, "id = ?", id)
}

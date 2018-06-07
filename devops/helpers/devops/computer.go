package devops

import (
	"time"

	"github.com/zhyq0826/go-tutorial/devops/db"
	model "github.com/zhyq0826/go-tutorial/devops/models/devops"
	"github.com/zhyq0826/go-tutorial/devops/resources"
)

// CreateComputer create computer
func CreateComputer(form resources.ComputerForm) {
	computer := model.Computer{
		CPU:       form.CPU,
		RAM:       form.RAM,
		PrivateIP: form.PrivateIP,
		PublicIP:  form.PublicIP,
	}
	now := time.Now()
	computer.UpdatedAt = now
	computer.CreatedAt = now
	db.DB.Create(&computer)
}

// UpdateComputer update
func UpdateComputer(id int, form resources.ComputerForm) {
	computer := model.Computer{
		CPU:       form.CPU,
		RAM:       form.RAM,
		PrivateIP: form.PrivateIP,
		PublicIP:  form.PublicIP,
	}
	db.DB.Model(&computer).Where("id=?", id).Updates(computer)
}

// QueryComputer computer list
func QueryComputer(page, limit int) []model.Computer {
	if limit == 0 {
		limit = 10
	}
	ret := make([]model.Computer, limit)
	db.DB.Where("").Limit(limit).Offset((page - 1) * limit).Find(&ret)

	return ret
}

func DeleteComputer(id int) {
	db.DB.Delete(&model.Computer{}, "id = ?", id)
}

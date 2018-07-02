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
func QueryComputer(page, limit, appid int) []model.Computer {
	if limit == 0 {
		limit = 10
	}
	conditon := make([]interface{}, 0)
	query := ""
	if appid != 0 {
		query += "app_id = ?"
		conditon = append(conditon, appid)
	}
	ret := make([]model.Computer, limit)
	db.DB.Where(query, conditon...).Limit(limit).Offset((page - 1) * limit).Find(&ret)

	return ret
}

// DeleteComputer for computer delete
func DeleteComputer(id int) {
	db.DB.Delete(&model.Computer{}, "id = ?", id)
}

// CreateDisk for disk create
func CreateDisk(form resources.DiskForm) {
	disk := model.Disk{
		Size:       form.Size,
		Left:       form.Left,
		ComputerID: form.ComputerID,
	}
	now := time.Now()
	disk.UpdatedAt = now
	disk.CreatedAt = now
	db.DB.Create(&disk)
}

// UpdateDisk for disk update
func UpdateDisk(id int, form resources.DiskForm) {
	disk := model.Disk{
		Size:       form.Size,
		Left:       form.Left,
		ComputerID: form.ComputerID,
	}
	disk.UpdatedAt = time.Now()
	db.DB.Model(&disk).Where("id=?", id).Updates(disk)
}

// DeleteDisk for disk delete
func DeleteDisk(id int) {
	db.DB.Delete(&model.Disk{}, "id = ?", id)
}

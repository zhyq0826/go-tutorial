package devops

import (
	"time"
)

// BaseModel for all devops model
type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null"`
	ID        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
}

// Domain model 域名
type Domain struct {
	BaseModel
	Name    string `gorm:"column:name;type:varchar(256)"`
	URL     string `gorm:"column:url;type:varchar(256)"`
	Private uint   `gorm:"column:private;type:tinyint;not null;default:0"`
}

// TableName for domain in sql
func (Domain) TableName() string {
	return "domain"
}

// Computer 机器
type Computer struct {
	BaseModel
	CPU       uint   `gorm:"column:cpu;type:int(11)"`
	RAM       uint   `gorm:"column:ram;type:int(11)"`
	PrivateIP string `gorm:"column:private_ip;type:varchar(256)"`
	PublicIP  string `gorm:"column:public_ip;type:varchar(256)"`
}

// Disk 磁盘
type Disk struct {
	BaseModel
	// 容量
	Size int `gorm:"column:size;type:int(11)"`
	// 剩余容量
	Left int `gorm:"column:left;type:int(11)"`
	// 机器 id
	ComputerID int `gorm:"column:computer_id;type:int(11)"`
}

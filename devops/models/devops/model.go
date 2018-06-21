package devops

import (
	"bytes"
	"database/sql/driver"
	"errors"
	"time"
)

type JSON []byte

// Value value
func (j JSON) Value() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

//Scan value
func (j *JSON) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}
	s, ok := value.([]byte)
	if !ok {
		return errors.New("Invalid Scan Source")
	}
	*j = append((*j)[0:0], s...)
	return nil
}

// MarshalJSON value
func (j JSON) MarshalJSON() ([]byte, error) {
	if j == nil {
		return []byte("null"), nil
	}
	return j, nil
}

//UnmarshalJSON value
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("null point exception")
	}
	*j = append((*j)[0:0], data...)
	return nil
}
func (j JSON) IsNull() bool {
	return len(j) == 0 || string(j) == "null"
}
func (j JSON) Equals(j1 JSON) bool {
	return bytes.Equal([]byte(j), []byte(j1))
}

// BaseModel for all devops model
type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"updated_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"created_at,omitempty"`
	ID        int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id,omitempty"`
}

// Domain model 域名 顶级或二级域名
type Domain struct {
	BaseModel
	Name    string `gorm:"column:name;type:varchar(256)" json:"name"`
	Host    string `gorm:"column:host;type:varchar(256)" json:"host"`
	Private uint   `gorm:"column:private;type:tinyint;not null;default:0" json:"private"`
	IP      JSON   `gorm:"column:ip;type:json;" json:"ip"`
}

// TableName for domain in sql
func (Domain) TableName() string {
	return "domain"
}

// Computer 机器
type Computer struct {
	BaseModel
	CPU       uint   `gorm:"column:cpu;type:int(11);default:0" json:"cpu"`
	RAM       uint   `gorm:"column:ram;type:int(11);default:0" json:"ram"`
	PrivateIP string `gorm:"column:private_ip;type:varchar(256):default:''" json:"private_ip"`
	PublicIP  string `gorm:"column:public_ip;type:varchar(256):default:''" json:"public_ip"`
	// 服务 id 一般来说一台机器最好部署单一的服务
	ServiceID uint `gorm:"column:service_id;type:int(11);default:0" json:"service_id"`
}

// Disk 磁盘
type Disk struct {
	BaseModel
	// 容量
	Size int `gorm:"column:size;type:int(11);default:0" json:"size"`
	// 剩余容量
	Left int `gorm:"column:left;type:int(11):default:0" json:"left"`
	// 机器 id
	ComputerID int `gorm:"column:computer_id;type:int(11)" json:"computer_id"`
}

// Service 服务 一般对应一个地址，可以外部地址或内部地址，地址可以是一个域名或一个路径。服务包括基础服务和应用服务
// 应用服务 执行特定功能的服务，比如认证鉴权、商品列表、订单交易等
// 基础服务 提供应用服务基础功能的服务，比如数据库服务、消息队列服务、缓存服务等等
type Service struct {
	BaseModel
	// 服务名称
	Name string `gorm:"column:name;type:varchar(256)" json:"name"`
	// 服务地址
	URL string `gorm:"column:url;type:varchar(256)" json:"url"`
	// 服务功能描述
	Desc string `gorm:"column:desc;type:varchar(512)" json:"desc"`
	// 是否核心链路
	IsImportant uint `gorm:"column:is_important;type:tinyint(11):default:0" json:"is_important"`
	// 仓库地址
	RepositoryURL string `gorm:"column:repository_url;type:varchar(256)" json:"repository_url"`
	// 部署目录
	DeployDir string `gorm:"column:deploy_dir;type:varchar(256)" json:"deploy_dir"`
	// 服务监控地址
	MonitorURL string `gorm:"column:monitor_url;type:varchar(256)" json:"monitor_url"`
}

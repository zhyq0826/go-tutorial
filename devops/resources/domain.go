package resources

import (
	"github.com/zhyq0826/go-tutorial/devops/models/devops"
)

type (
	// DomainForm for json data
	DomainForm struct {
		devops.Domain
		// ID        int       `json:"id,omitempty"`
		// Name      string    `json:"name"`
		// URL       string    `json:"url"`
		// Private   uint      `json:"private"`
		// CreatedAt time.Time `json:"created_at,omitempty"`
		// UpdatedAt time.Time `json:"updated_at,omitempty"`
	}

	// ComputerForm for json data
	ComputerForm struct {
		devops.Computer
		// ID        int       `json:"id"`
		// CPU       uint      `json:"cpu"`
		// RAM       uint      `json:"ram"`
		// PrivateIP string    `json:"private_ip"`
		// PublicIP  string    `json:"public_ip"`
		// CreatedAt time.Time `json:"created_at,omitempty"`
		// UpdatedAt time.Time `json:"updated_at,omitempty"`
	}

	// DiskForm for json data
	DiskForm struct {
		devops.Disk
		// ID int `json:"id,omitempty"`
		// // 容量
		// Size int `json:"size"`
		// // 剩余容量
		// Left int `json:"left"`
		// // 机器 id
		// ComputerID int       `json:"computer_id"`
		// CreatedAt  time.Time `json:"created_at,omitempty"`
		// UpdatedAt  time.Time `json:"updated_at,omitempty"`
	}

	// ServiceForm for json data
	ServiceForm struct {
		devops.Service
	}
)

package devops

import (
	"time"
)

// BaseModel for all devops model
type BaseModel struct {
	CreatedAt *time.Time `gorm:"column:created_at;type:datetime;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;type:datetime;not null"`
	ID        int        `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT"`
}

// Domain model
type Domain struct {
	BaseModel
	Name    string `gorm:"column:name;type:varchar(256)"`
	URL     string `gorm:"column:url;type:varchar(256)"`
	Private uint   `gorm:"column:private;type:tinyint;not null;default:0"`
}

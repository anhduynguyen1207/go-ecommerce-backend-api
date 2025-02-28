package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID        uuid.UUID `gorm:"column:id; type:int; not null; primary_key; auto_increment; comment:'Primary Key is ID'"`
	RoleName  string    `gorm:"column:role_name"`
	RolesNote string    `gorm:"column:role_note; type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}

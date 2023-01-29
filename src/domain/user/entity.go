package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"name"`
	Email       string    `gorm:"column:email;unique"`
	Password    string    `gorm:"column:password"`
	IsAdmin     bool      `gorm:"column:is_admin;default:false"`
	CreatedDate time.Time `json:"created_date" gorm:"<-:create;column:created_date;default:null"`
	UpdatedDate time.Time `json:"updated_date" gorm:"column:updated_date;default:null"`
}

func (p User) TableName() string {
	return "users"
}

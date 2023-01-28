package review

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/domain/user"
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID          uuid.UUID        `gorm:"type:uuid;primaryKey"`
	Name        string           `gorm:"name"`
	Rating      int              `gorm:"column:rating"`
	Comment     string           `gorm:"column:comment"`
	CreatedDate time.Time        `json:"created_date" gorm:"<-:create;column:created_date;default:null"`
	Product     *product.Product `gorm:"foreignKey:ProductId"`
	ProductId   string           ``
	User        *user.User       `gorm:"foreignKey:UserId"`
	UserId      string
}

func (r Review) TableName() string {
	return "reviews"
}

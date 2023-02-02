package product

import (
	"time"

	"github.com/google/uuid"

	"shoeshop-backend/src/domain/user"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"name"`
	Description string    `gorm:"column:description"`
	ImageUrl    string    `gorm:"column:image_url"`
	Rating      int8      `gorm:"column:rating"`
	Price       float64   `gorm:"column:price"`
	StockCount  float64   `gorm:"column:stock"`
	Reviews     []*Review `gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	CreatedDate time.Time `json:"created_date" gorm:"<-:create;column:created_date;default:null"`
	DeletedDate time.Time `json:"deleted_date" gorm:"column:deleted_date;default:null"`
	UpdatedDate time.Time `json:"updated_date" gorm:"column:updated_date;default:null"`
}

func (p Product) TableName() string {
	return "products"
}

type Review struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey"`
	Rating      int        `gorm:"column:rating"`
	Comment     string     `gorm:"column:comment"`
	CreatedDate time.Time  `json:"created_date" gorm:"<-:create;column:created_date;default:null"`
	Product     *Product   `gorm:"foreignKey:ProductId"`
	ProductId   string     ``
	User        *user.User `gorm:"foreignKey:UserId"`
	UserId      string
}

func (r Review) TableName() string {
	return "reviews"
}

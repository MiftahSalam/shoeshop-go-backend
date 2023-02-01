package order

import (
	"shoeshop-backend/src/domain/product"
	"shoeshop-backend/src/domain/user"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID              uuid.UUID     `gorm:"type:uuid;primaryKey"`
	User            *user.User    `gorm:"foreignKey:UserId"`
	UserId          string        ``
	Items           []*Item       `gorm:"foreignKey:OrderId;constraint:OnDelete:CASCADE"`
	ShippingAddress Shipping      `gorm:"embedded;embeddedPrefix:shipping_"`
	PaymentMethod   string        `gorm:"column:payment_method"`
	PaymentStatus   PaymentResult `gorm:"embedded;embeddedPrefix:payment_"`
	TaxPrice        float64       `gorm:"column:tax_price"`
	ShippingPrice   float64       `gorm:"column:shipping_price"`
	TotalPrice      float64       `gorm:"column:total_price"`
	IsPaid          bool          `gorm:"column:is_paid"`
	PaidAt          time.Time     `json:"paid_date" gorm:"column:paid_at"`
	IsDelivered     bool          `gorm:"column:is_delivered"`
	DeliveredAt     time.Time     `json:"delivered_at" gorm:"delivered_at"`
	CreatedAt       time.Time     `json:"created_at" gorm:"<-:create;column:created_at;default:null"`
}

type Item struct {
	ID          uuid.UUID        `gorm:"type:uuid;primaryKey"`
	Name        string           `gorm:"name"`
	Quantity    int              `gorm:"column:quantity"`
	Price       float64          `gorm:"column:price"`
	CreatedAt   time.Time        `json:"created_at" gorm:"<-:create;column:created_at;default:null"`
	UpdatedDate time.Time        `json:"updated_at" gorm:"column:updated_at;default:null"`
	Product     *product.Product `gorm:"foreignKey:ProductId"`
	ProductId   string           ``
	OrderId     string
}

type Shipping struct {
	Address    string `gorm:"name"`
	City       string `gorm:"column:city"`
	PostalCode string `gorm:"column:postal_code"`
	Country    string `gorm:"column:country"`
}

type PaymentResult struct {
	ID         string    `gorm:"id"`
	Status     string    `gorm:"status"`
	UpdateTime time.Time `gorm:"column:update"`
	Email      string    `gorm:"column:email"`
}

func (p Order) TableName() string {
	return "orders"
}

func (i Item) TableName() string {
	return "items"
}

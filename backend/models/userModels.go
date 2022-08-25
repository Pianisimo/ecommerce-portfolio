package models

import (
	"time"
)

type User struct {
	ID             int64     `json:"id" sql:"id"`
	FirstName      string    `json:"first_name" validate:"required,min=2,max=30" sql:"first_name"`
	LastName       string    `json:"last_name" validate:"required,min=2,max=30" sql:"last_name"`
	Password       string    `json:"password" validate:"required,min=6,max=30" sql:"password"`
	Email          string    `json:"email" validate:"required,email" sql:"email"`
	Phone          string    `json:"phone" sql:"phone"`
	Token          string    `json:"token" sql:"token"`
	RefreshToken   string    `json:"refresh_token" sql:"refresh_token"`
	CSRF           string    `json:"csrf" sql:"csrf"`
	CreatedAt      time.Time `json:"created_at" sql:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" sql:"updated_at"`
	UserCart       []Product `json:"user_cart" sql:"user_cart"`
	AddressDetails []Address `json:"address" sql:"address_details"`
	OrderStatus    []Order   `json:"orders" sql:"order_status"`
}

type Product struct {
	ID     *int64  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Price  *uint64 `json:"price,omitempty"`
	Rating *uint8  `json:"rating,omitempty"`
	Image  *string `json:"image,omitempty"`
}

type Address struct {
	ID      *int64  `json:"id,omitempty"`
	House   *string `json:"house,omitempty"`
	Street  *string `json:"street,omitempty"`
	City    *string `json:"city,omitempty"`
	Pincode *string `json:"pincode,omitempty"`
}

type Order struct {
	ID            *int64    `json:"id,omitempty"`
	Cart          []Product `json:"cart,omitempty"`
	OrderedAt     time.Time `json:"ordered_at"`
	Price         *uint64   `json:"price,omitempty"`
	Discount      *int      `json:"discount,omitempty"`
	PaymentMethod Payment   `json:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital,omitempty"`
	COD     bool `json:"cod,omitempty"`
}

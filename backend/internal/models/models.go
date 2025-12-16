package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	CustomerName    string         `json:"customer_name" gorm:"not null"`
	Whatsapp        string         `json:"whatsapp" gorm:"not null"`
	ProductName     string         `json:"product_name" gorm:"not null"`
	PackageType     string         `json:"package_type" gorm:"not null"` // starter, growth, pro
	Price           int            `json:"price" gorm:"not null"`
	Status          string         `json:"status" gorm:"default:'pending'"` // pending, paid, completed
	PaymentProofURL string         `json:"payment_proof_url"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

type UmkmPage struct {
	ID              uint           `json:"id" gorm:"primaryKey"`
	OrderID         uint           `json:"order_id" gorm:"not null"`
	Slug            string         `json:"slug" gorm:"uniqueIndex;not null"`
	VideoURL        string         `json:"video_url"`
	Photos          string         `json:"photos" gorm:"type:jsonb"` // JSON array of photo URLs
	Description     string         `json:"description" gorm:"type:text"`
	Price           string         `json:"price"`
	WhatsappLink    string         `json:"whatsapp_link"`
	MarketplaceLink string         `json:"marketplace_link"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"-" gorm:"index"`
}

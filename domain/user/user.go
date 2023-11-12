package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// User model
type User struct {
	ID           string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserName     string         `gorm:"type:char(70);unique" json:"username"`
	FirstName    string         `gorm:"type:char(30)" json:"firstname"`
	LastName     string         `gorm:"type:char(30)" json:"lastname"`
	Password     string         `json:"password,omitempty"`
	Language     Language       `gorm:"not null;type:char(3)" json:"language"`
	Email        string         `gorm:"type:char(70)" json:"email"`
	Phone        string         `gorm:"type:char(30)" json:"phone"`
	Photo        string         `gorm:"type:char(100)" json:"photo"`
	Roles  []Roles         `json:"-"`
	ClientID     string         `json:"client_id"`
	ClientSecret string         `json:"client_secret"`
	Token        string         `json:"token"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `json:"-"`
}

func (c *User) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}

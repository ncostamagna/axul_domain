package user

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//User model
type User struct {
	ID           string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserName     string         `gorm:"size:70;unique" json:"username"`
	FirstName    string         `gorm:"size:30" json:"firstname"`
	LastName     string         `gorm:"size:30" json:"lastname"`
	Password     string         `json:"password"`
	Email        string         `gorm:"size:70" json:"email"`
	Phone        string         `gorm:"size:30" json:"phone"`
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
package contact

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// Contact model
type Contact struct {
	ID         string         `json:"id" gorm:"type:char(36);not null;primary_key;unique_index"`
	UserID     string         `json:"user_id" gorm:"type:char(36);"`
	Firstname  string         `gorm:"size:30" json:"firstname"`
	Lastname   string         `gorm:"size:30" json:"lastname"`
	Nickname   string         `gorm:"size:30" json:"nickname"`
	Gender     string         `gorm:"size:1" json:"gender"`
	TypeID     string         `gorm:"size:40" json:"type"`
	Phone      string         `gorm:"size:30" json:"phone"`
	Instagram  string         `gorm:"size:40" json:"instagram"`
	Photo      string         `gorm:"size:200" json:"photo"`
	Birthday   time.Time      `json:"birthday"`
	Days       int64          `gorm:"-" json:"days"`
	TemplateID string         `json:"template_id" gorm:"type:char(36);"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}

func (c *Contact) BeforeCreate(tx *gorm.DB) (err error) {

	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	return
}

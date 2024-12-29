package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID 		`gorm:"primaryKey;column:id;not null"`
	Name        string 		`gorm:"column:name;not null"`
	PhoneNumber string 		`gorm:"column:phone_number"`
	Email       string 		`gorm:"column:email"`
	Address     string 		`gorm:"column:address"`
	CreatedAt   time.Time	`gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt	time.Time	`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Books		[]Book		`gorm:"foreignKey:user_id;references:id"`
}

func (user *User) TableName() string {
	return "users"
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	

	return nil
}

func (user *User) AfterUpdate(tx *gorm.DB) error {
	for _, userBook := range user.Books {
		tx.Model(&Book{}).Where("id = ? ", userBook.ID).Update("return_date", time.Now().Add(time.Hour * 336)).
		Update("availability", false)
	}
	
	return nil
}

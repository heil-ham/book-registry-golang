package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        		uuid.UUID	`gorm:"primaryKey;column:id;not null"`
	UserId    		string		`gorm:"column:user_id;default:null"`
	Title     		string	    `gorm:"column:title"`
	Availability	bool		`gorm:"column:availability"`
	CreatedAt 		time.Time 	`gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt 		time.Time 	`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	ReturnDate 		time.Time 	`gorm:"column:return_date;autoCreateTime;default:null"`
}

func (book *Book) TableName() string {
	return "books"
}

func (book *Book) BeforeCreate(tx *gorm.DB) error {
	if book.ID == uuid.Nil {
		book.ID = uuid.New()
	}
	
	if !book.Availability {
		book.Availability = true
	}
	
	book.ReturnDate = time.Now()

	return nil

}

func (book *Book) BeforeSave(tx *gorm.DB) error {
	if !book.Availability {
		book.Availability = true
	}
	
	book.ReturnDate = time.Now()

	return nil

}
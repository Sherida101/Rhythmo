// Rhythmo — Build better habits with rhythm
// https://github.com/Sherida101/Rhythmo
//
// See LICENSE for copyright details.

package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Habit represents a user's habit tracking record in the application.
type Habit struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:100;not null" json:"name" validate:"required,max=100"`
	Description string         `gorm:"size:255" json:"description" validate:"max=255"`
	Category    string         `gorm:"size:50" json:"category" validate:"required,max=50"`
	Status      string         `gorm:"size:20" json:"status" validate:"required,oneof=active inactive"`
	Priority    string         `gorm:"size:20" json:"priority" validate:"required,oneof=low medium high"`
	Frequency   string         `gorm:"size:20" json:"frequency" validate:"required,oneof=daily weekly monthly"`
	Reminder    string         `gorm:"size:100" json:"reminder" validate:"max=100"`
	Streak      int            `gorm:"default:0" json:"streak" validate:"gte=0"`
	StartDate   time.Time      `gorm:"not null" json:"startDate" validate:"required"`
	EndDate     time.Time      `gorm:"not null" json:"endDate" validate:"required,gtfield=StartDate"`
	LastDone    time.Time      `gorm:"not null" json:"lastDone" validate:"required,ltfield=EndDate"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
}

// Validate performs custom validation logic on the Habit struct.
func (h *Habit) Validate() error {
	validate := validator.New()
	return validate.Struct(h)
}

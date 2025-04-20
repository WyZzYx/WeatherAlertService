package models

import (
	"gorm.io/gorm"
	"time"
)

type Subscription struct {
	gorm.Model
	City       string
	Condition  string
	LastSentAt *time.Time
	Users      []User `gorm:"many2many:user_subscriptions;"`
}

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string
	Email         string         `gorm:"uniqueIndex"`
	Subscriptions []Subscription `gorm:"many2many:user_subscriptions;"`
}

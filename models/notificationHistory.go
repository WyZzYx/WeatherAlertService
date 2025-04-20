package models

type NotificationHistory struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	City      string
	Condition string
	SentAt    int64
}

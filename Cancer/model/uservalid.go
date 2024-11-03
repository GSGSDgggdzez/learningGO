package model

type UserValid struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

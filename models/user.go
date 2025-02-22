package models

type User struct {
	ID        string `json:"id" gorm:"primaryKey"`
	FirstName string `json:"first_name" gorm:"type:varchar(100)"`
	LastName  string `json:"last_name" gorm:"type:varchar(100)"`
}

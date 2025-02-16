package data

type Password struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Pass string
	Login string
	Theme string
	Description string
}
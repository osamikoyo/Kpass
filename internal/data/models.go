package data

type Password struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Hash string
	Login string
	Theme string
	Description string
}

type Account struct{
	ID uint `gorm:"primaryKey;autoIncrement"`
	Password string
	Name string
}
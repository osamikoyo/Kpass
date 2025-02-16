package data

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct{
	db *gorm.DB
}

func New(path string) (*Storage, error) {
	db, err := gorm.Open(sqlite.Open(path))
	if err != nil{
		return nil, fmt.Errorf("error open db: %w", err)
	}

	if err = db.AutoMigrate(&Account{}, &Password{});err != nil{
		return nil, fmt.Errorf("error automigrate: %w", err)
	}

	return &Storage{
		db,
	}, nil
}

func (s *Storage) AddAccount(acc *Account) error {
	return s.db.Create(acc).Error
}

func (s *Storage) CheckPass(accname, pass string) (bool, error){
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil{
		return false, fmt.Errorf("cant get hash: %w", err)
	}

	var acc Account

	result := s.db.Where(&Account{Name: accname}).Find(&acc)
	if result.Error != nil{
		return false, fmt.Errorf("cant get user: %w", err)
	}

	if acc.Password == string(hash){
		return true, nil
	}
	return false, nil
}

func (s *Storage) AddPassword(pass *Password, Accpassword string) error {
	key, err := bcrypt.GenerateFromPassword([]byte(pass.Hash), bcrypt.DefaultCost)
	if err != nil{
		return fmt.Errorf("cant get hash: %w", err)
	}

	
}
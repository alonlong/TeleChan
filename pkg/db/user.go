package db

import (
	"TeleChan/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableUser - database table 'users'
	TableUser = "users"
)

// FindUserByEmail queries the user by 'email'
func (s *Handler) FindUserByEmail(email string) (*model.User, error) {
	u := model.User{}

	// query the user rows by email
	res := s.db.Table(TableUser).Where("email = ?", email).First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// FindUserByID queries the user by unique key 'user_id'
func (s *Handler) FindUserByUK(userID string) (*model.User, error) {
	u := model.User{}

	// query the user rows by user id
	res := s.db.Table(TableUser).Where("user_id = ?", userID).First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// CreateUser insert the user into database
func (s *Handler) CreateUser(u *model.User) error {
	return s.db.Table(TableUser).Create(u).Error
}

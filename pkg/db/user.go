package db

import (
	"TeleChan/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableUser - database table 'users'
	TableUser = "users"
)

// FindUserByKey queries the user by unique key 'email'
func (s *Handler) FindUserByKey(email string) (*model.User, error) {
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

// CreateUser insert the user into database
func (s *Handler) CreateUser(u *model.User) error {
	return s.db.Table(TableUser).Create(u).Error
}

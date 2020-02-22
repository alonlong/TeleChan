package db

import (
	"TeleChan/pkg/db/model"

	"github.com/jinzhu/gorm"
)

const (
	// TableChannel - database table 'channels'
	TableChannel = "channels"
	// TableUserChannel - database table 'users_channels'
	TableUserChannel = "users_channels"
)

// FindChannelByName queries the channel by 'name'
func (s *Handler) FindChannelByName(name string) (*model.Channel, error) {
	u := model.Channel{}

	// query the channel rows by channel name
	res := s.db.Table(TableChannel).Where("name = ?", name).First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// FindChannelByLikeName queries the channel by  ike 'name'
func (s *Handler) FindChannelsByLikeName(input string) ([]model.Channel, error) {
	u := []model.Channel{}

	// query the channel rows by like channel name
	res := s.db.Table(TableChannel).Where("name like ?", "%"+input+"%").Find(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return u, nil
}

// FindChannelByID queries the channels by unique key 'channel_id'
func (s *Handler) FindChannelByUK(channelID string) (*model.Channel, error) {
	u := model.Channel{}

	// query the channel rows by channel id
	res := s.db.Table(TableChannel).Where("channel_id = ?", channelID).First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// FindChannelsByUser queries the user's owned channels by 'user_id'
func (s *Handler) FindChannelsByUser(userID string) ([]model.Channel, error) {
	u := []model.Channel{}

	// query the channel rows by channel id
	res := s.db.Table(TableChannel).Where("user_id = ?", userID).Find(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return u, nil
		}
		return nil, res.Error
	}

	return u, nil
}

// CreateChannel insert the channel into database
func (s *Handler) CreateChannel(u *model.Channel) error {
	return s.db.Table(TableChannel).Create(u).Error
}

// CreateUserChannel insert the user's joining channel into database
func (s *Handler) CreateUserChannel(u *model.UserChannel) error {
	return s.db.Table(TableUserChannel).Create(u).Error
}

// FindUserChannelByUK queries the user channel by unique key 'user_id' and 'channel_id'
func (s *Handler) FindUserChannelByUK(userID string, channelID string) (*model.UserChannel, error) {
	u := model.UserChannel{}

	// query the user channel rows by user id and channel id
	res := s.db.Table(TableUserChannel).
		Where("user_id = ?", userID).
		Where("channel_id = ?", channelID).
		First(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, res.Error
	}

	return &u, nil
}

// DeleteUserChannelByUK deletes the user channel by unique key 'user_id' and 'channel_id'
func (s *Handler) DeleteUserChannelByUK(userID string, channelID string) error {
	u := model.UserChannel{}
	res := s.db.Table(TableUserChannel).
		Where("user_id = ?", userID).
		Where("channel_id = ?", channelID).
		Delete(u)

	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return nil
		}
		return res.Error
	}
	return nil
}

// FindUserChannelsByUser queries the user's owned and joined channels by 'user_id'
func (s *Handler) FindUserChannelsByUser(userID string) ([]model.UserChannel, error) {
	u := []model.UserChannel{}

	// query the user channel rows by user id
	res := s.db.Table(TableUserChannel).Where("user_id = ?", userID).Find(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return u, nil
		}
		return nil, res.Error
	}

	return u, nil
}

// FindJoinedUsersByChannel queries the users who joined the channel by 'channel_id'
func (s *Handler) FindJoinedUsersByChannel(channelID string) ([]model.UserChannel, error) {
	u := []model.UserChannel{}

	// query the users rows by channel id
	res := s.db.Table(TableUserChannel).
		Where("channel_id = ?", channelID).
		Where("owned = ?", false).
		Find(&u)
	if res.Error != nil {
		// if there is no record found
		if res.Error == gorm.ErrRecordNotFound {
			return u, nil
		}
		return nil, res.Error
	}

	return u, nil
}

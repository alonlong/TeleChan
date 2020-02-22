package db

import (
	"TeleChan/pkg/db/model"
)

const (
	// TableMessage - database table 'messages'
	TableMessage = "messages"
	// TableChannelMessage - database table 'channels_messages'
	TableChannelMessage = "channels_messages"
)

// CreateMessage insert the message into database
func (s *Handler) CreateMessage(m *model.Message) error {
	res, err := s.db.DB().Exec("insert into messages (content) VALUES (?)", m.Content)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	m.ID = int32(id)

	return nil
}

// CreateChannelMessage insert the channel message into database
func (s *Handler) CreateChannelMessage(u *model.ChannelMessage) error {
	return s.db.Table(TableChannelMessage).Create(u).Error
}

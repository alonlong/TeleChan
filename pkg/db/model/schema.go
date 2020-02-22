package model

// User represents the database model for table 'users'
type User struct {
	UserID   string // the user id
	Email    string // the user's email
	Password string // the user's password
}

// Channel represents the database model for table 'channels'
type Channel struct {
	ChannelID string // the channel id
	Name      string // the channel name
	UserID    string // the user id who create the channel
}

// UserChannel represents the database model for table 'users_channels'
type UserChannel struct {
	UserID    string // the user id
	ChannelID string // the channel id
	Name      string // the channel name
	Owned     bool   // user owned or joined the channel
}

// Message represents the database model for table 'messages'
type Message struct {
	ID      int32    // the message sequence number
	Content string // the message content
}

// ChannelMessage represents the database model for table 'channels_messages'
type ChannelMessage struct {
	ChannelID string // the channel id
	MessageID int32  // the sequence number of message
	Brief     string // the brief of message content
}

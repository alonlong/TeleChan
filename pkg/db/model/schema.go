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

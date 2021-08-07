package repo

import "time"

type Session struct {
	ID         string    `gorm:"type:VARCHAR(50)"` // session id , will be in the cookie
	AccountId  string    `gorm:"type:VARCHAR(50)"` // account owner
	ExpiryTime time.Time // time when session is no longer valid
	CreateTime time.Time // time when session was made
}

// ISession : User Login Session Repository
type ISession interface {
	// FindById : Find specific session that is not expired
	FindById(id string) *Session
	// Create : Create the login session , the create time is handled for you
	Create(session *Session) (*Session, bool)
	// Update : Update a session
	Update(session *Session) (*Session, bool)
	// Delete : Delete the login session
	Delete(id string) bool
}

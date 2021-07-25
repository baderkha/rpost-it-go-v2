package repo

type Session struct {
	ID        string
	AccountId string
}

// ISession : User Login Session Repository
type ISession interface {
	// FindById : Find specific session
	FindById(id string) *Session
	// Create : Create the login session
	Create(session *Session) (*Session, bool)
	// Delete : Delete the login session
	Delete(id string) bool
}

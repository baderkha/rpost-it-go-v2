package repo

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type MYSQLSession struct {
	db *gorm.DB
}

func NewMYSQLSessionRepo(db *gorm.DB) *MYSQLSession {
	return &MYSQLSession{
		db: db,
	}
}

func (ms *MYSQLSession) FindById(id string) *Session {
	var record Session
	ms.db.First(&record, "id=?", id)
	return &record
}

func (ms *MYSQLSession) Create(session *Session) (*Session, bool) {
	session.ID = uuid.NewV4().String()
	isCreated := ms.db.Create(session).Error == nil
	return session, isCreated
}

func (ms *MYSQLSession) Delete(id string) bool {
	return ms.db.Unscoped().Model(&Session{}).Where("id=?", id).Delete(&Session{}).Error == nil
}

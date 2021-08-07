package repo

import (
	"time"

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
	ms.db.Where("expiry_time>=?", time.Now()).First(&record, "id=?", id)
	return &record
}

func (ms *MYSQLSession) Create(session *Session) (*Session, bool) {
	session.ID = uuid.NewV4().String()
	session.CreateTime = time.Now()
	isCreated := ms.db.Create(session).Error == nil

	return session, isCreated
}

func (ms *MYSQLSession) Update(session *Session) (*Session, bool) {
	isUpdated := ms.db.Updates(session).Error == nil
	return session, isUpdated
}

func (ms *MYSQLSession) Delete(id string) bool {
	return ms.db.Unscoped().Model(&Session{}).Where("id=?", id).Delete(&Session{}).Error == nil
}

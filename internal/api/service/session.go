package service

import (
	"rpost-it-go/internal/api/repo"
	"time"
)

const (
	SessionExpiryMinutesTimeDefault = 20
)

type session struct {
	BaseService
	repo repo.ISession
}

func (s *session) getTimeByDefaultMinFromNow() time.Time {
	return time.
		Now().
		Add(time.Minute * time.Duration(SessionExpiryMinutesTimeDefault))
}

func (s *session) isExpiredSession(session *repo.Session) bool {
	return session.ExpiryTime.Before(time.Now())
}

func (s *session) create(accountId string) (*repo.Session, error) {
	session, isCreated := s.repo.Create(&repo.Session{
		AccountId:  accountId,
		ExpiryTime: s.getTimeByDefaultMinFromNow(),
	})

	if !isCreated {
		return nil, s.Error().InternalError()
	}
	return session, nil
}

func (s *session) get(id string) (*repo.Session, error) {
	if id == "" {
		return nil, s.Error().UnAuthorized()
	}
	session := s.repo.FindById(id)
	if session == nil || session.ID != id {
		return nil, s.Error().UnAuthorized()
	} else if s.isExpiredSession(session) {
		_ = s.delete(id)
		return nil, s.Error().ExpiredLogin()
	}
	return session, nil
}

func (s *session) refresh(session *repo.Session) (*repo.Session, error) {
	session.ExpiryTime = s.getTimeByDefaultMinFromNow()
	ses, isUpdated := s.repo.Update(session)
	if !isUpdated {
		return nil, s.Error().InternalError()
	}
	return ses, nil
}

func (s *session) delete(id string) error {
	isDeleted := s.repo.Delete(id)
	if !isDeleted {
		return s.Error().InternalError()
	}
	return nil
}

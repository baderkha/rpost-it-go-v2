package service

import "rpost-it-go/internal/api/repo"

type session struct {
	BaseService
	repo repo.ISession
}

func (s *session) create(accountId string) (*repo.Session, error) {
	session, isCreated := s.repo.Create(&repo.Session{
		AccountId: accountId,
	})
	if !isCreated {
		return nil, s.Error().InternalError()
	}
	return session, nil
}

func (s *session) get(id string) (*repo.Session, error) {
	session := s.repo.FindById(id)
	if session == nil || session.ID != id {
		return nil, s.Error().UnAuthorized()
	}
	return session, nil
}

func (s *session) delete(id string) error {
	isDeleted := s.repo.Delete(id)
	if !isDeleted {
		return s.Error().InternalError()
	}
	return nil
}

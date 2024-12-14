package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type SessionService interface {
	GetSessionByID(sessionID string) (*models.Session, error)
	CreateSession(session *models.Session) error
	DeleteSession(sessionID string) error
	DeleteExpiredSessions() error
	GetSessionByToken(token string) (*models.Session, error)
}

type sessionService struct {
	repo repositories.SessionRepository
}

func NewSessionService(repo repositories.SessionRepository) SessionService {
	return &sessionService{repo: repo}
}

func (s *sessionService) GetSessionByID(sessionID string) (*models.Session, error) {
	return s.repo.GetSessionByID(sessionID)
}

func (s *sessionService) CreateSession(session *models.Session) error {
	return s.repo.CreateSession(session)
}

func (s *sessionService) DeleteSession(sessionID string) error {
	return s.repo.DeleteSession(sessionID)
}

func (s *sessionService) DeleteExpiredSessions() error {
	return s.repo.DeleteExpiredSessions()
}

func (s *sessionService) GetSessionByToken(token string) (*models.Session, error) {
	return s.repo.GetSessionByToken(token)
}

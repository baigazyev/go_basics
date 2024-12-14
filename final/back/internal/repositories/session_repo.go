package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type SessionRepository interface {
	GetSessionByID(sessionID string) (*models.Session, error)
	CreateSession(session *models.Session) error
	DeleteSession(sessionID string) error
	DeleteExpiredSessions() error
	SaveSession(session *models.Session) error
	GetSessionByToken(token string) (*models.Session, error)
	DeleteSessionByToken(token string) error
}

type sessionRepository struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

func (r *sessionRepository) GetSessionByID(sessionID string) (*models.Session, error) {
	var session models.Session
	err := r.db.First(&session, "session_id = ?", sessionID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &session, err
}

func (r *sessionRepository) CreateSession(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepository) DeleteSession(sessionID string) error {
	return r.db.Delete(&models.Session{}, "session_id = ?", sessionID).Error
}

func (r *sessionRepository) DeleteExpiredSessions() error {
	return r.db.Where("expires_at < NOW()").Delete(&models.Session{}).Error
}

func (r *sessionRepository) SaveSession(session *models.Session) error {
	return r.db.Create(session).Error
}

func (r *sessionRepository) GetSessionByToken(token string) (*models.Session, error) {
	var session models.Session
	err := r.db.Where("token = ?", token).First(&session).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &session, err
}

func (r *sessionRepository) DeleteSessionByToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&models.Session{}).Error
}

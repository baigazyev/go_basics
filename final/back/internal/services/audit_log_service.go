package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type AuditLogService interface {
	GetAllAuditLogs() ([]models.AuditLog, error)
	GetAuditLogsByUserID(userID int) ([]models.AuditLog, error)
	CreateAuditLog(log *models.AuditLog) error
}

type auditLogService struct {
	repo repositories.AuditLogRepository
}

func NewAuditLogService(repo repositories.AuditLogRepository) AuditLogService {
	return &auditLogService{repo: repo}
}

func (s *auditLogService) GetAllAuditLogs() ([]models.AuditLog, error) {
	return s.repo.GetAllAuditLogs()
}

func (s *auditLogService) GetAuditLogsByUserID(userID int) ([]models.AuditLog, error) {
	return s.repo.GetAuditLogsByUserID(userID)
}

func (s *auditLogService) CreateAuditLog(log *models.AuditLog) error {
	return s.repo.CreateAuditLog(log)
}

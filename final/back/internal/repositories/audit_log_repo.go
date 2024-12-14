package repositories

import (
	"e-commerce/internal/models"

	"gorm.io/gorm"
)

type AuditLogRepository interface {
	GetAllAuditLogs() ([]models.AuditLog, error)
	GetAuditLogsByUserID(userID int) ([]models.AuditLog, error)
	CreateAuditLog(log *models.AuditLog) error
}

type auditLogRepository struct {
	db *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

func (r *auditLogRepository) GetAllAuditLogs() ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Find(&logs).Error
	return logs, err
}

func (r *auditLogRepository) GetAuditLogsByUserID(userID int) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.db.Where("user_id = ?", userID).Find(&logs).Error
	return logs, err
}

func (r *auditLogRepository) CreateAuditLog(log *models.AuditLog) error {
	return r.db.Create(log).Error
}

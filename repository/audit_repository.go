package repository

import (
	"backend-avanzada/models"
	"errors"

	"gorm.io/gorm"
)

type AuditRepository struct {
	db *gorm.DB
}

func NewAuditRepository(db *gorm.DB) *AuditRepository {
	return &AuditRepository{db: db}
}

func (r *AuditRepository) FindAll() ([]*models.Audit, error) {
	var audits []*models.Audit
	err := r.db.Preload("Alchemist").Find(&audits).Error
	return audits, err
}

func (r *AuditRepository) FindById(id uint) (*models.Audit, error) {
	var audit models.Audit
	err := r.db.Preload("Alchemist").First(&audit, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &audit, nil
}

func (r *AuditRepository) FindByAlchemistID(alchemistID uint) ([]*models.Audit, error) {
	var audits []*models.Audit
	err := r.db.Where("alchemist_id = ?", alchemistID).Preload("Alchemist").Find(&audits).Error
	return audits, err
}

func (r *AuditRepository) FindByType(auditType models.AuditType) ([]*models.Audit, error) {
	var audits []*models.Audit
	err := r.db.Where("type = ?", auditType).Preload("Alchemist").Find(&audits).Error
	return audits, err
}

func (r *AuditRepository) FindUnresolved() ([]*models.Audit, error) {
	var audits []*models.Audit
	err := r.db.Where("resolved = ?", false).Preload("Alchemist").Find(&audits).Error
	return audits, err
}

func (r *AuditRepository) Save(audit *models.Audit) (*models.Audit, error) {
	err := r.db.Save(audit).Error
	return audit, err
}

func (r *AuditRepository) Create(audit *models.Audit) (*models.Audit, error) {
	err := r.db.Create(audit).Error
	return audit, err
}

func (r *AuditRepository) Delete(audit *models.Audit) error {
	return r.db.Delete(audit).Error
}

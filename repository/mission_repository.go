package repository

import (
	"backend-avanzada/models"
	"errors"

	"gorm.io/gorm"
)

type MissionRepository struct {
	db *gorm.DB
}

func NewMissionRepository(db *gorm.DB) *MissionRepository {
	return &MissionRepository{db: db}
}

func (r *MissionRepository) FindAll() ([]*models.Mission, error) {
	var missions []*models.Mission
	err := r.db.Preload("Alchemist").Find(&missions).Error
	return missions, err
}

func (r *MissionRepository) FindById(id uint) (*models.Mission, error) {
	var mission models.Mission
	err := r.db.Preload("Alchemist").First(&mission, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &mission, nil
}

func (r *MissionRepository) FindByAlchemistID(alchemistID uint) ([]*models.Mission, error) {
	var missions []*models.Mission
	err := r.db.Where("alchemist_id = ?", alchemistID).Preload("Alchemist").Find(&missions).Error
	return missions, err
}

func (r *MissionRepository) FindByStatus(status models.MissionStatus) ([]*models.Mission, error) {
	var missions []*models.Mission
	err := r.db.Where("status = ?", status).Preload("Alchemist").Find(&missions).Error
	return missions, err
}

func (r *MissionRepository) Save(mission *models.Mission) (*models.Mission, error) {
	err := r.db.Save(mission).Error
	return mission, err
}

func (r *MissionRepository) Create(mission *models.Mission) (*models.Mission, error) {
	err := r.db.Create(mission).Error
	return mission, err
}

func (r *MissionRepository) Delete(mission *models.Mission) error {
	return r.db.Delete(mission).Error
}

package repository

import (
	"backend-avanzada/models"
	"errors"

	"gorm.io/gorm"
)

type TransmutationRepository struct {
	db *gorm.DB
}

func NewTransmutationRepository(db *gorm.DB) *TransmutationRepository {
	return &TransmutationRepository{db: db}
}

func (r *TransmutationRepository) FindAll() ([]*models.Transmutation, error) {
	var transmutations []*models.Transmutation
	err := r.db.Preload("Alchemist").
		Preload("InputMaterials.Material").
		Preload("OutputMaterials.Material").
		Find(&transmutations).Error
	return transmutations, err
}

func (r *TransmutationRepository) FindById(id uint) (*models.Transmutation, error) {
	var transmutation models.Transmutation
	err := r.db.Preload("Alchemist").
		Preload("InputMaterials.Material").
		Preload("OutputMaterials.Material").
		First(&transmutation, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &transmutation, nil
}

func (r *TransmutationRepository) FindByAlchemistID(alchemistID uint) ([]*models.Transmutation, error) {
	var transmutations []*models.Transmutation
	err := r.db.Where("alchemist_id = ?", alchemistID).
		Preload("Alchemist").
		Preload("InputMaterials.Material").
		Preload("OutputMaterials.Material").
		Find(&transmutations).Error
	return transmutations, err
}

func (r *TransmutationRepository) FindByStatus(status models.TransmutationStatus) ([]*models.Transmutation, error) {
	var transmutations []*models.Transmutation
	err := r.db.Where("status = ?", status).
		Preload("Alchemist").
		Preload("InputMaterials.Material").
		Preload("OutputMaterials.Material").
		Find(&transmutations).Error
	return transmutations, err
}

func (r *TransmutationRepository) Save(transmutation *models.Transmutation) (*models.Transmutation, error) {
	err := r.db.Save(transmutation).Error
	return transmutation, err
}

func (r *TransmutationRepository) Create(transmutation *models.Transmutation) (*models.Transmutation, error) {
	err := r.db.Create(transmutation).Error
	return transmutation, err
}

func (r *TransmutationRepository) Delete(transmutation *models.Transmutation) error {
	return r.db.Delete(transmutation).Error
}

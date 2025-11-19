package repository

import (
	"backend-avanzada/models"
	"errors"

	"gorm.io/gorm"
)

type AlchemistRepository struct {
	db *gorm.DB
}

func NewAlchemistRepository(db *gorm.DB) *AlchemistRepository {
	return &AlchemistRepository{db: db}
}

func (r *AlchemistRepository) FindAll() ([]*models.Alchemist, error) {
	var alchemists []*models.Alchemist
	err := r.db.Find(&alchemists).Error
	return alchemists, err
}

func (r *AlchemistRepository) FindById(id uint) (*models.Alchemist, error) {
	var alchemist models.Alchemist
	err := r.db.First(&alchemist, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &alchemist, nil
}

func (r *AlchemistRepository) FindByEmail(email string) (*models.Alchemist, error) {
	var alchemist models.Alchemist
	err := r.db.Where("email = ?", email).First(&alchemist).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &alchemist, nil
}

func (r *AlchemistRepository) Save(alchemist *models.Alchemist) (*models.Alchemist, error) {
	err := r.db.Save(alchemist).Error
	return alchemist, err
}

func (r *AlchemistRepository) Create(alchemist *models.Alchemist) (*models.Alchemist, error) {
	err := r.db.Create(alchemist).Error
	return alchemist, err
}

func (r *AlchemistRepository) Delete(alchemist *models.Alchemist) error {
	return r.db.Delete(alchemist).Error
}

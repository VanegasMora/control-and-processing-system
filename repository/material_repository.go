package repository

import (
	"backend-avanzada/models"
	"errors"

	"gorm.io/gorm"
)

type MaterialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) *MaterialRepository {
	return &MaterialRepository{db: db}
}

func (r *MaterialRepository) FindAll() ([]*models.Material, error) {
	var materials []*models.Material
	err := r.db.Find(&materials).Error
	return materials, err
}

func (r *MaterialRepository) FindById(id uint) (*models.Material, error) {
	var material models.Material
	err := r.db.First(&material, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &material, nil
}

func (r *MaterialRepository) FindByName(name string) (*models.Material, error) {
	var material models.Material
	err := r.db.Where("name = ?", name).First(&material).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &material, nil
}

func (r *MaterialRepository) Save(material *models.Material) (*models.Material, error) {
	err := r.db.Save(material).Error
	return material, err
}

func (r *MaterialRepository) Create(material *models.Material) (*models.Material, error) {
	err := r.db.Create(material).Error
	return material, err
}

func (r *MaterialRepository) Delete(material *models.Material) error {
	return r.db.Delete(material).Error
}

func (r *MaterialRepository) UpdateStock(materialID uint, quantity float64) error {
	return r.db.Model(&models.Material{}).Where("id = ?", materialID).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}

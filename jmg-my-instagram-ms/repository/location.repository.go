package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{db: config.GetConnection()}
}

func (p *LocationRepository) Save(Location *models.Location) error {
	return p.db.Create(Location).Error
}

func (p *LocationRepository) Update(Location *models.Location) error {
	return p.db.Updates(Location).Error
}

func (p *LocationRepository) Delete(id uint) error {
	return p.db.Delete(&models.Location{}, id).Error
}

func (p *LocationRepository) FindById(id uint) (*models.Location, error) {
	var Location models.Location
	if err := p.db.First(&Location, id).Error; err != nil {
		return nil, err
	}
	return &Location, nil
}

func (p *LocationRepository) FindByName(name string) (*models.Location, error) {
	var Location models.Location
	if err := p.db.Where("name = ?", name).First(&Location).Error; err != nil {
		return nil, err
	}
	return &Location, nil
}

func (p *LocationRepository) FindAll() ([]*models.Location, error) {
	var products []*models.Location
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *LocationRepository) Search(name string, top int) ([]models.Location, error) {
	var locations []models.Location
	result := p.db.Where("concat(city,country) ILIKE ?", "%"+name+"%").Limit(top).Find(&locations)

	if result.Error != nil {
		return nil, result.Error
	}

	return locations, nil
}

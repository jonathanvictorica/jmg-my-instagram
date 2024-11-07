package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type MultimediaContentRepository struct {
	db *gorm.DB
}

func NewMultimediaContentRepository() *MultimediaContentRepository {
	return &MultimediaContentRepository{db: config.GetConnection()}
}

func (p *MultimediaContentRepository) Save(MultimediaContent *models.MultimediaContent) error {
	return p.db.Create(MultimediaContent).Error
}

func (p *MultimediaContentRepository) Update(MultimediaContent *models.MultimediaContent) error {
	return p.db.Updates(MultimediaContent).Error
}

func (p *MultimediaContentRepository) Delete(id uint) error {
	return p.db.Delete(&models.MultimediaContent{}, id).Error
}

func (p *MultimediaContentRepository) FindById(id uint) (*models.MultimediaContent, error) {
	var MultimediaContent models.MultimediaContent
	if err := p.db.First(&MultimediaContent, id).Error; err != nil {
		return nil, err
	}
	return &MultimediaContent, nil
}

func (p *MultimediaContentRepository) FindByName(name string) (*models.MultimediaContent, error) {
	var MultimediaContent models.MultimediaContent
	if err := p.db.Where("name = ?", name).First(&MultimediaContent).Error; err != nil {
		return nil, err
	}
	return &MultimediaContent, nil
}

func (p *MultimediaContentRepository) FindAll() ([]*models.MultimediaContent, error) {
	var products []*models.MultimediaContent
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

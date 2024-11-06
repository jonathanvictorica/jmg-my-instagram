package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type HashtagRepository struct {
	db *gorm.DB
}

func NewHashtagRepository() *HashtagRepository {
	return &HashtagRepository{db: config.GetConnection()}
}

func (p *HashtagRepository) Save(Hashtag *models.Hashtag) error {
	return p.db.Create(Hashtag).Error
}

func (p *HashtagRepository) Update(Hashtag *models.Hashtag) error {
	return p.db.Updates(Hashtag).Error
}

func (p *HashtagRepository) Delete(id uint) error {
	return p.db.Delete(&models.Hashtag{}, id).Error
}

func (p *HashtagRepository) FindById(id uint) (*models.Hashtag, error) {
	var Hashtag models.Hashtag
	if err := p.db.First(&Hashtag, id).Error; err != nil {
		return nil, err
	}
	return &Hashtag, nil
}

func (p *HashtagRepository) FindByName(name string) (*models.Hashtag, error) {
	var Hashtag models.Hashtag
	if err := p.db.Where("name = ?", name).First(&Hashtag).Error; err != nil {
		return nil, err
	}
	return &Hashtag, nil
}

func (p *HashtagRepository) FindAll() ([]*models.Hashtag, error) {
	var products []*models.Hashtag
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

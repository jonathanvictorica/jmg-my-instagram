package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type StoryContentHashtagRepository struct {
	db *gorm.DB
}

func NewStoryContentHashtagRepository() *StoryContentHashtagRepository {
	return &StoryContentHashtagRepository{db: config.GetConnection()}
}

func (p *StoryContentHashtagRepository) Save(StoryContentHashtag *models.StoryContentHashtag) error {
	return p.db.Create(StoryContentHashtag).Error
}

func (p *StoryContentHashtagRepository) Update(StoryContentHashtag *models.StoryContentHashtag) error {
	return p.db.Updates(StoryContentHashtag).Error
}

func (p *StoryContentHashtagRepository) Delete(id uint) error {
	return p.db.Delete(&models.StoryContentHashtag{}, id).Error
}

func (p *StoryContentHashtagRepository) FindById(id uint) (*models.StoryContentHashtag, error) {
	var StoryContentHashtag models.StoryContentHashtag
	if err := p.db.First(&StoryContentHashtag, id).Error; err != nil {
		return nil, err
	}
	return &StoryContentHashtag, nil
}

func (p *StoryContentHashtagRepository) FindByName(name string) (*models.StoryContentHashtag, error) {
	var StoryContentHashtag models.StoryContentHashtag
	if err := p.db.Where("name = ?", name).First(&StoryContentHashtag).Error; err != nil {
		return nil, err
	}
	return &StoryContentHashtag, nil
}

func (p *StoryContentHashtagRepository) FindAll() ([]*models.StoryContentHashtag, error) {
	var products []*models.StoryContentHashtag
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

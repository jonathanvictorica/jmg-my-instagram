package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type StoryContentRepository struct {
	db *gorm.DB
}

func NewStoryContentRepository() *StoryContentRepository {
	return &StoryContentRepository{db: config.GetConnection()}
}

func (p *StoryContentRepository) Save(StoryContent *models.StoryContent) error {
	return p.db.Create(StoryContent).Error
}

func (p *StoryContentRepository) Update(StoryContent *models.StoryContent) error {
	return p.db.Updates(StoryContent).Error
}

func (p *StoryContentRepository) Delete(id uint) error {
	return p.db.Delete(&models.StoryContent{}, id).Error
}

func (p *StoryContentRepository) FindById(id uint) (*models.StoryContent, error) {
	var StoryContent models.StoryContent
	if err := p.db.First(&StoryContent, id).Error; err != nil {
		return nil, err
	}
	return &StoryContent, nil
}

func (p *StoryContentRepository) FindByName(name string) (*models.StoryContent, error) {
	var StoryContent models.StoryContent
	if err := p.db.Where("name = ?", name).First(&StoryContent).Error; err != nil {
		return nil, err
	}
	return &StoryContent, nil
}

func (p *StoryContentRepository) FindAll() ([]*models.StoryContent, error) {
	var products []*models.StoryContent
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

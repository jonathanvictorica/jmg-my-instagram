package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type StoryRepository struct {
	db *gorm.DB
}

func NewStoryRepository() *StoryRepository {
	return &StoryRepository{db: config.GetConnection()}
}

func (p *StoryRepository) Save(Story *models.Story) error {
	return p.db.Create(Story).Error
}

func (p *StoryRepository) Update(Story *models.Story) error {
	return p.db.Updates(Story).Error
}

func (p *StoryRepository) Delete(id uint) error {
	return p.db.Delete(&models.Story{}, id).Error
}

func (p *StoryRepository) FindById(id uint) (*models.Story, error) {
	var Story models.Story
	if err := p.db.First(&Story, id).Error; err != nil {
		return nil, err
	}
	return &Story, nil
}

func (p *StoryRepository) FindByName(name string) (*models.Story, error) {
	var Story models.Story
	if err := p.db.Where("name = ?", name).First(&Story).Error; err != nil {
		return nil, err
	}
	return &Story, nil
}

func (p *StoryRepository) FindAll() ([]*models.Story, error) {
	var products []*models.Story
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

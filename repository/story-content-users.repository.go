package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type StoryContentUsersRepository struct {
	db *gorm.DB
}

func NewStoryContentUsersRepository() *StoryContentUsersRepository {
	return &StoryContentUsersRepository{db: config.GetConnection()}
}

func (p *StoryContentUsersRepository) Save(StoryContentUsers *models.StoryContentUsers) error {
	return p.db.Create(StoryContentUsers).Error
}

func (p *StoryContentUsersRepository) Update(StoryContentUsers *models.StoryContentUsers) error {
	return p.db.Updates(StoryContentUsers).Error
}

func (p *StoryContentUsersRepository) Delete(id uint) error {
	return p.db.Delete(&models.StoryContentUsers{}, id).Error
}

func (p *StoryContentUsersRepository) FindById(id uint) (*models.StoryContentUsers, error) {
	var StoryContentUsers models.StoryContentUsers
	if err := p.db.First(&StoryContentUsers, id).Error; err != nil {
		return nil, err
	}
	return &StoryContentUsers, nil
}

func (p *StoryContentUsersRepository) FindByName(name string) (*models.StoryContentUsers, error) {
	var StoryContentUsers models.StoryContentUsers
	if err := p.db.Where("name = ?", name).First(&StoryContentUsers).Error; err != nil {
		return nil, err
	}
	return &StoryContentUsers, nil
}

func (p *StoryContentUsersRepository) FindAll() ([]*models.StoryContentUsers, error) {
	var products []*models.StoryContentUsers
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

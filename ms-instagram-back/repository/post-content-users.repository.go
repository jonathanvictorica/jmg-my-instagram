package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type PostContentUsersRepository struct {
	db *gorm.DB
}

func NewPostContentUsersRepository() *PostContentUsersRepository {
	return &PostContentUsersRepository{db: config.GetConnection()}
}

func (p *PostContentUsersRepository) Save(PostContentUsers *models.PostContentUsers) error {
	return p.db.Create(PostContentUsers).Error
}

func (p *PostContentUsersRepository) Update(PostContentUsers *models.PostContentUsers) error {
	return p.db.Updates(PostContentUsers).Error
}

func (p *PostContentUsersRepository) Delete(id uint) error {
	return p.db.Delete(&models.PostContentUsers{}, id).Error
}

func (p *PostContentUsersRepository) FindById(id uint) (*models.PostContentUsers, error) {
	var PostContentUsers models.PostContentUsers
	if err := p.db.First(&PostContentUsers, id).Error; err != nil {
		return nil, err
	}
	return &PostContentUsers, nil
}

func (p *PostContentUsersRepository) FindByName(name string) (*models.PostContentUsers, error) {
	var PostContentUsers models.PostContentUsers
	if err := p.db.Where("name = ?", name).First(&PostContentUsers).Error; err != nil {
		return nil, err
	}
	return &PostContentUsers, nil
}

func (p *PostContentUsersRepository) FindAll() ([]*models.PostContentUsers, error) {
	var products []*models.PostContentUsers
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository() *PostRepository {
	return &PostRepository{db: config.GetConnection()}
}

func (p *PostRepository) Save(Post *models.Post) error {
	return p.db.Create(Post).Error
}

func (p *PostRepository) Update(Post *models.Post) error {
	return p.db.Updates(Post).Error
}

func (p *PostRepository) Delete(id uint) error {
	return p.db.Delete(&models.Post{}, id).Error
}

func (p *PostRepository) FindById(id uint) (*models.Post, error) {
	var Post models.Post
	if err := p.db.First(&Post, id).Error; err != nil {
		return nil, err
	}
	return &Post, nil
}

func (p *PostRepository) FindByName(name string) (*models.Post, error) {
	var Post models.Post
	if err := p.db.Where("name = ?", name).First(&Post).Error; err != nil {
		return nil, err
	}
	return &Post, nil
}

func (p *PostRepository) FindAll() ([]*models.Post, error) {
	var products []*models.Post
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

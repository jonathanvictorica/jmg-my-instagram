package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type PostContentRepository struct {
	db *gorm.DB
}

func NewPostContentRepository() *PostContentRepository {
	return &PostContentRepository{db: config.GetConnection()}
}

func (p *PostContentRepository) Save(PostContent *models.PostContent) error {
	return p.db.Create(PostContent).Error
}

func (p *PostContentRepository) Update(PostContent *models.PostContent) error {
	return p.db.Updates(PostContent).Error
}

func (p *PostContentRepository) Delete(id uint) error {
	return p.db.Delete(&models.PostContent{}, id).Error
}

func (p *PostContentRepository) FindById(id uint) (*models.PostContent, error) {
	var PostContent models.PostContent
	if err := p.db.First(&PostContent, id).Error; err != nil {
		return nil, err
	}
	return &PostContent, nil
}

func (p *PostContentRepository) FindByName(name string) (*models.PostContent, error) {
	var PostContent models.PostContent
	if err := p.db.Where("name = ?", name).First(&PostContent).Error; err != nil {
		return nil, err
	}
	return &PostContent, nil
}

func (p *PostContentRepository) FindAll() ([]*models.PostContent, error) {
	var products []*models.PostContent
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

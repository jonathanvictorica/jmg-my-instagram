package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type PostContentHashtagRepository struct {
	db *gorm.DB
}

func NewPostContentHashtagRepository() *PostContentHashtagRepository {
	return &PostContentHashtagRepository{db: config.GetConnection()}
}

func (p *PostContentHashtagRepository) Save(PostContentHashtag *models.PostContentHashtag) error {
	return p.db.Create(PostContentHashtag).Error
}

func (p *PostContentHashtagRepository) Update(PostContentHashtag *models.PostContentHashtag) error {
	return p.db.Updates(PostContentHashtag).Error
}

func (p *PostContentHashtagRepository) Delete(id uint) error {
	return p.db.Delete(&models.PostContentHashtag{}, id).Error
}

func (p *PostContentHashtagRepository) FindById(id uint) (*models.PostContentHashtag, error) {
	var PostContentHashtag models.PostContentHashtag
	if err := p.db.First(&PostContentHashtag, id).Error; err != nil {
		return nil, err
	}
	return &PostContentHashtag, nil
}

func (p *PostContentHashtagRepository) FindByName(name string) (*models.PostContentHashtag, error) {
	var PostContentHashtag models.PostContentHashtag
	if err := p.db.Where("name = ?", name).First(&PostContentHashtag).Error; err != nil {
		return nil, err
	}
	return &PostContentHashtag, nil
}

func (p *PostContentHashtagRepository) FindAll() ([]*models.PostContentHashtag, error) {
	var products []*models.PostContentHashtag
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

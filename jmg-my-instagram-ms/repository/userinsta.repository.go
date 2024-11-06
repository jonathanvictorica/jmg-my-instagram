package repository

import (
	"gorm.io/gorm"
	"jmg-my-instagram-ms/config"
	models "jmg-my-instagram-ms/model"
)

type UserInstaRepository struct {
	db *gorm.DB
}

func NewUserInstaRepository() *UserInstaRepository {
	return &UserInstaRepository{db: config.GetConnection()}
}

func (p *UserInstaRepository) Save(UserInsta *models.UserInsta) error {
	return p.db.Create(UserInsta).Error
}

func (p *UserInstaRepository) Update(UserInsta *models.UserInsta) error {
	return p.db.Updates(UserInsta).Error
}

func (p *UserInstaRepository) Delete(id uint) error {
	return p.db.Delete(&models.UserInsta{}, id).Error
}

func (p *UserInstaRepository) FindById(id uint) (*models.UserInsta, error) {
	var UserInsta models.UserInsta
	if err := p.db.First(&UserInsta, id).Error; err != nil {
		return nil, err
	}
	return &UserInsta, nil
}

func (p *UserInstaRepository) FindByName(name string) (*models.UserInsta, error) {
	var UserInsta models.UserInsta
	if err := p.db.Where("name = ?", name).First(&UserInsta).Error; err != nil {
		return nil, err
	}
	return &UserInsta, nil
}

func (p *UserInstaRepository) FindAll() ([]*models.UserInsta, error) {
	var products []*models.UserInsta
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p *UserInstaRepository) Search(name string, top int) ([]models.UserInsta, error) {
	var usersInstas []models.UserInsta
	result := p.db.Where("Username ILIKE ?", "%"+name+"%").Limit(top).Find(&usersInstas)

	if result.Error != nil {
		return nil, result.Error
	}

	return usersInstas, nil
}

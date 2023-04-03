package repositories

import (
	"github.com/hyperyuri/webapi-with-go/models"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	CreateApplication(application *models.Application) error
}

type applicationRepositoryImpl struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &applicationRepositoryImpl{db: db}
}

func (a *applicationRepositoryImpl) CreateApplication(application *models.Application) error {
	return a.db.Create(application).Error
}

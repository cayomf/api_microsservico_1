package repositories

import (
	"github.com/hyperyuri/webapi-with-go/models"
	"gorm.io/gorm"
)

type VacancyRepository interface {
	FindAll() ([]models.Vacancy, error)
	FindLast(quantity int) ([]models.Vacancy, error)
	FindById(id int) (models.Vacancy, error)
	Create(vacancy *models.Vacancy) error
	Delete(id int) error
	Update(vacancy *models.Vacancy) error
}

type vacancyRepositoryImpl struct {
	db *gorm.DB
}

func NewVacancyRepository(db *gorm.DB) VacancyRepository {
	return &vacancyRepositoryImpl{db: db}
}

func (vr *vacancyRepositoryImpl) FindAll() ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	err := vr.db.Find(&vacancies).Error
	return vacancies, err
}

func (vr *vacancyRepositoryImpl) FindLast(quantity int) ([]models.Vacancy, error) {
	var vacancies []models.Vacancy
	err := vr.db.Limit(quantity).Find(&vacancies).Error
	return vacancies, err
}

func (vr *vacancyRepositoryImpl) FindById(id int) (models.Vacancy, error) {
	var vacancy models.Vacancy
	err := vr.db.First(&vacancy, id).Error
	return vacancy, err
}

func (vr *vacancyRepositoryImpl) Create(vacancy *models.Vacancy) error {
	return vr.db.Create(vacancy).Error
}

func (vr *vacancyRepositoryImpl) Delete(id int) error {
	return vr.db.Delete(&models.Vacancy{}, id).Error
}

func (vr *vacancyRepositoryImpl) Update(vacancy *models.Vacancy) error {
	return vr.db.Save(vacancy).Error
}

package services

import (
	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/repositories"
)

type VacancyService interface {
	GetAll() ([]models.Vacancy, error)
	GetLast(quantity int) ([]models.Vacancy, error)
	GetById(id int) (models.Vacancy, error)
	Create(vacancy *models.Vacancy) error
	Delete(id int) error
	Update(vacancy *models.Vacancy) error
}

type vacancyServiceImpl struct {
	repo repositories.VacancyRepository
}

func NewVacancyService(repo repositories.VacancyRepository) VacancyService {
	return &vacancyServiceImpl{repo: repo}
}

func (vs *vacancyServiceImpl) GetAll() ([]models.Vacancy, error) {
	return vs.repo.FindAll()
}

func (vs *vacancyServiceImpl) GetLast(quantity int) ([]models.Vacancy, error) {
	return vs.repo.FindLast(quantity)
}

func (vs *vacancyServiceImpl) GetById(id int) (models.Vacancy, error) {
	return vs.repo.FindById(id)
}

func (vs *vacancyServiceImpl) Create(vacancy *models.Vacancy) error {
	return vs.repo.Create(vacancy)
}

func (vs *vacancyServiceImpl) Delete(id int) error {
	return vs.repo.Delete(id)
}

func (vs *vacancyServiceImpl) Update(vacancy *models.Vacancy) error {
	return vs.repo.Update(vacancy)
}

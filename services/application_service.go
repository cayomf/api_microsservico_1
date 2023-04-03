package services

import (
	"github.com/hyperyuri/webapi-with-go/models"
	"github.com/hyperyuri/webapi-with-go/repositories"
)

type ApplicationService interface {
	ApplyForVacancy(application *models.Application) error
}

type applicationServiceImpl struct {
	applicationRepository repositories.ApplicationRepository
}

func NewApplicationService(applicationRepository repositories.ApplicationRepository) ApplicationService {
	return &applicationServiceImpl{applicationRepository: applicationRepository}
}

func (a *applicationServiceImpl) ApplyForVacancy(application *models.Application) error {
	return a.applicationRepository.CreateApplication(application)
}

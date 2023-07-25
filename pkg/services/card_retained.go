package services

import (
	"EJM/dto"
	"EJM/pkg/models"
	"EJM/pkg/repository"
	"gorm.io/gorm"
)

type CardRetainedService struct {
	*gorm.DB
	CardRetainedRepository repository.CardRetainedRepository
}

func NewCardRetainedService(service *CardRetainedService) *CardRetainedService {
	return service
}

func (cardRetained *CardRetainedService) GetCardRetained(cardRetainedDto *dto.GetCardRetained) ([]models.Term, error) {

	var cardRetainedRepo  repository.CardRetainedRepository = cardRetained.CardRetainedRepository

	data, err := cardRetainedRepo.GetCardRetained(cardRetainedDto) 
	if err != nil {
		return []models.Term{}, err
	}

	return data, nil
}

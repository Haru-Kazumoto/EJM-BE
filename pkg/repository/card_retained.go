package repository

import (
	"EJM/dto"
	"EJM/pkg/models"
	"gorm.io/gorm"
)

type CardRetainedRepository interface {
	GetCardRetained(filter *dto.GetCardRetained) ([]models.Term, error)}

type CardRetained struct {
	db *gorm.DB
}

func NewCardRetainedRepository(db *gorm.DB) *CardRetained {
	return &CardRetained{db: db}
}

func (cardRetainedObject *CardRetained) CardRetainedModel() (tx *gorm.DB) {
	return cardRetainedObject.db.Model(&models.Term{})
}

func (cardRetainedObject *CardRetained) GetCardRetained(filter *dto.GetCardRetained) ([]models.Term, error) {
	var cardRetained []models.Term
	data := cardRetainedObject.CardRetainedModel()

		if filter.TermID != "" {
			data = data.Where("term_id = ?", filter.TermID)
		}
		if filter.ATMKey != "" {
			data = data.Where("atm_key = ?", filter.ATMKey)
		}
		if filter.TermType != "" {
			data = data.Where("term_type = ?", filter.TermType)
		}
		if filter.Merk != "" {
			data = data.Where("merk = ?", filter.Merk)
		}
		if filter.ManagedBy != "" {
			data = data.Where("managed_by = ?", filter.ManagedBy)
		}
		if filter.Jarkom != "" {
			data = data.Where("jarkom = ?", filter.Jarkom)
		}
		if filter.Region != "" {
			data = data.Where("region = ?", filter.Region)
		}
		if filter.Area != "" {
			data = data.Where("area = ?", filter.Area)
		}

	// cari data
	data.Scopes().Debug().
		Find(&cardRetained)

	// checking errors
	if err := data.Error; err != nil {
		return []models.Term{}, err
	}

	return cardRetained, nil
}

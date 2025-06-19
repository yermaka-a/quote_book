package repository

import "quote_book/internal/core/models"

type QuoteRepository interface {
	GetRandom() (*models.Quote, error)
	Add(*models.Quote)
	List() ([]*models.Quote, error)
	ListByAuthor(string) ([]*models.Quote, error)
	DelByID(int) (int, error)
}

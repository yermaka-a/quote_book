package service

import (
	"errors"
	"log/slog"
	"quote_book/internal/core/models"
	"quote_book/internal/core/repository"
	"quote_book/internal/storage"
)

type QuoteService struct {
	repo repository.QuoteRepository
	log  *slog.Logger
}

func NewQuoteService(repo repository.QuoteRepository, log *slog.Logger) *QuoteService {
	return &QuoteService{
		repo: repo,
		log:  log}
}

func (s *QuoteService) GetRandom() (*models.Quote, error) {
	quote, err := s.repo.GetRandom()
	if errors.Is(err, storage.ErrMapIsEmpty) {
		return nil, ErrHasNoQuotes
	}
	if err != nil {
		s.log.Error("error getting random quote", "err", err)
		return nil, ErrGettingRandQuote
	}
	return quote, nil
}

func (s *QuoteService) Add(quote *models.Quote) {
	s.repo.Add(quote)
}

func (s *QuoteService) List() ([]*models.Quote, error) {
	list, err := s.repo.List()
	if errors.Is(err, storage.ErrListIsEmpty) {
		return nil, ErrListOfQuotesIsEmpty
	}
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *QuoteService) ListByAuthor(author string) ([]*models.Quote, error) {
	list, err := s.repo.ListByAuthor(author)
	if errors.Is(err, storage.ErrListIsEmpty) {
		return nil, ErrListOfQuotesIsEmpty
	}
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *QuoteService) DelByID(id int) (int, error) {
	id, err := s.repo.DelByID(id)
	if err != nil {
		s.log.Info("error deleting quote", "err", err)
		return 0, errors.New("quote isn't deleted")
	}
	return id, nil
}

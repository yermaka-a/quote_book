package storage

import (
	"math/rand"
	"quote_book/internal/core/models"
	"sync"
)

type MemoryQuoteStorage struct {
	CounterID int
	quotes    map[string]map[int]*models.Quote
	mx        sync.RWMutex
}

func New() *MemoryQuoteStorage {
	return &MemoryQuoteStorage{
		CounterID: 1,
		quotes:    make(map[string]map[int]*models.Quote),
	}
}

func (s *MemoryQuoteStorage) GetRandom() (*models.Quote, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	if len(s.quotes) == 0 {
		return nil, ErrMapIsEmpty
	}

	authors := make([]string, 0, len(s.quotes))
	for author := range s.quotes {
		authors = append(authors, author)
	}

	randomAuthor := authors[rand.Intn(len(authors))]

	authorQuotes := s.quotes[randomAuthor]
	if len(authorQuotes) == 0 {
		return nil, ErrHasNoSubRecords
	}

	quoteIDs := make([]int, 0, len(authorQuotes))
	for id := range authorQuotes {
		quoteIDs = append(quoteIDs, id)
	}

	randomQuoteID := quoteIDs[rand.Intn(len(quoteIDs))]

	return authorQuotes[randomQuoteID], nil
}

func (s *MemoryQuoteStorage) Add(quote *models.Quote) {
	s.mx.Lock()
	defer s.mx.Unlock()
	quote.ID = s.CounterID
	s.CounterID++

	if v, isExists := s.quotes[quote.Author]; isExists {
		v[quote.ID] = quote
		return
	}

	s.quotes[quote.Author] = make(map[int]*models.Quote)
	s.quotes[quote.Author][quote.ID] = quote
}

func (s *MemoryQuoteStorage) List() ([]*models.Quote, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	listQuotes := make([]*models.Quote, 0)
	for _, v := range s.quotes {
		for _, quote := range v {
			listQuotes = append(listQuotes, quote)
		}
	}
	if len(listQuotes) == 0 {
		return nil, ErrListIsEmpty
	}
	return listQuotes, nil
}

func (s *MemoryQuoteStorage) ListByAuthor(author string) ([]*models.Quote, error) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	listQuotes := make([]*models.Quote, 0)
	if v, isExists := s.quotes[author]; isExists {
		for _, quote := range v {
			listQuotes = append(listQuotes, quote)
		}
	} else {
		return nil, ErrRecordNotFound
	}

	if len(listQuotes) == 0 {
		return nil, ErrListIsEmpty
	}

	return listQuotes, nil
}

func (s *MemoryQuoteStorage) DelByID(id int) (int, error) {
	s.mx.Lock()
	defer s.mx.Unlock()
	for author, v := range s.quotes {
		if _, isExists := v[id]; isExists {
			delete(v, id)
			if len(v) == 0 {
				delete(s.quotes, author)
			}
			return id, nil
		}
	}
	return 0, ErrRecordNotFound
}

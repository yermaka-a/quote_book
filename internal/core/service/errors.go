package service

import "errors"

var (
	ErrListOfQuotesIsEmpty = errors.New("list of quotes is empty")
	ErrHasNoQuotes         = errors.New("No one quote has now")
	ErrGettingRandQuote = errors.New("error getting random quote") 
)

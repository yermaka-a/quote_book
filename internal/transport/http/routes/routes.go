package routes

import (
	"net/http"
	"quote_book/internal/transport/http/handler"
)

func SetupRoutes(r *http.ServeMux, quoteHandler *handler.QuoteHandler) {
	r.HandleFunc("GET /quotes/random", quoteHandler.GetQuote)
	r.HandleFunc("DELETE /quotes/{id}", quoteHandler.DelQuote)
	r.HandleFunc("POST /quotes", quoteHandler.AddQuote)
	r.HandleFunc("GET /quotes", quoteHandler.ListQuotes)
}

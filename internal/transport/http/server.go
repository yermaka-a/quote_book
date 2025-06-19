package http

import (
	"net/http"
	"quote_book/internal/transport/http/handler"
	"quote_book/internal/transport/http/routes"
)

type Server struct {
	router *http.ServeMux
}

func NewServer(quoteHandler *handler.QuoteHandler) *Server {
	r := http.NewServeMux()
	routes.SetupRoutes(r, quoteHandler)
	return &Server{router: r}
}

func (s *Server) Run(port string) error {
	return http.ListenAndServe(port, s.router)
}

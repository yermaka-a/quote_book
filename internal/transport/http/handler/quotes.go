package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"quote_book/internal/core/models"
	"quote_book/internal/core/service"
	"strconv"
	"strings"
)

type QuoteHandler struct {
	service *service.QuoteService
	log     *slog.Logger
}

func NewQuoteHandler(service *service.QuoteService, log *slog.Logger) *QuoteHandler {
	return &QuoteHandler{
		service: service,
		log:     log}
}

func (h *QuoteHandler) responseJSON(httpcode int, msg any, w http.ResponseWriter) {
	type response struct {
		Reponse *any `json:"reponse,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpcode)
	json.NewEncoder(w).Encode(
		response{
			Reponse: &msg,
		},
	)
}

func (h *QuoteHandler) GetQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := h.service.GetRandom()
	if errors.Is(err, service.ErrHasNoQuotes) {
		h.responseJSON(http.StatusNotFound, "no quotes exist", w)
		return
	}

	if err != nil {
		h.responseJSON(http.StatusInternalServerError, "internal error", w)
		return
	}

	h.responseJSON(http.StatusOK, quote, w)
}

func (h *QuoteHandler) AddQuote(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		h.responseJSON(http.StatusUnsupportedMediaType, "Content-Type must be application/json", w)
		return
	}

	var quote models.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		h.responseJSON(http.StatusBadRequest, err.Error(), w)
		return
	}

	h.service.Add(&quote)
	h.responseJSON(http.StatusCreated, "quote created", w)
}

func (h *QuoteHandler) DelQuote(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	strId := pathParts[len(pathParts)-1]
	id, err := strconv.Atoi(strId)
	if err != nil || id <= 0 {
		h.log.Debug("incorrect id parameter", "err", err, "id", id)
		h.responseJSON(http.StatusBadRequest, "incorrect id parameter", w)
		return
	}

	id, err = h.service.DelByID(id)
	if err != nil {
		h.log.Error("deleting error", "err", err)
		h.responseJSON(http.StatusInternalServerError, "quote isn't deleted", w)
		return
	}

	h.responseJSON(http.StatusOK, fmt.Sprintf("quote with id: %d has been deleted", id), w)
}

func (h *QuoteHandler) ListQuotes(w http.ResponseWriter, r *http.Request) {
	author, err := url.QueryUnescape(r.URL.Query().Get("author"))

	if err != nil {
		h.responseJSON(http.StatusBadRequest, "invalid author parameter", w)
		return
	}

	if author != "" {
		list, err := h.service.ListByAuthor(author)
		if err != nil {
			h.log.Error("error getting by author", "err", err)
			h.responseJSON(http.StatusInternalServerError, "error getting by author", w)
			return
		}
		h.responseJSON(http.StatusOK, list, w)
		return
	}

	list, err := h.service.List()
	if err != nil {
		if errors.Is(err, service.ErrListOfQuotesIsEmpty) {
			h.responseJSON(http.StatusNotFound, service.ErrListOfQuotesIsEmpty.Error(), w)
		} else {
			h.log.Error("error getting list of quotes", "err", err)
			h.responseJSON(http.StatusInternalServerError, "internal error", w)
		}
		return
	}
	h.responseJSON(http.StatusOK, list, w)
}

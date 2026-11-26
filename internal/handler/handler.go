package handler

import (
	"net/http"
	"strings"

	models "github.com/valeryvash/metrics/internal/model"
)

type Service interface {
	UpdateMetric(models.Metrics) error
}

type Handler struct {
	Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		badRequest(w)
		return
	}

	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	switch {
	case len(parts) != 4:
		notFoundRequest(w)
		return
	case parts[0] != "update":
		badRequest(w)
		return
	case parts[2] == "":
		notFoundRequest(w)
		return
	}

	var (
		m models.Metrics
		err error
	)

	m, err = models.NewMetric(parts[1], parts[3])
	if err != nil {
		badRequest(w)
		return
	}

	err = h.UpdateMetric(m)
	if err != nil {
		badRequest(w)
		return
	}

	ok(w)
}

func badRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func notFoundRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

func ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

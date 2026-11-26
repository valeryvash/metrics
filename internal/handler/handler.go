package handler

import (
	"net/http"
	"strconv"
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

	m := models.Metrics{}
	switch parts[1] {
	case models.Counter:
		m.MType = models.Counter
		delta, err := strconv.ParseInt(parts[3], 10, 64)
		if err != nil {
			badRequest(w)
			return
		}
		m.Delta = &delta
	case models.Gauge:
		m.MType = models.Gauge
		value, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			badRequest(w)
			return
		}
		m.Value = &value
	default:
		badRequest(w)
		return
	}

	err := h.UpdateMetric(m)
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

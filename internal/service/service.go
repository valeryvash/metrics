package service

import (
	"fmt"

	models "github.com/valeryvash/metrics/internal/model"
)

type Repository interface {
	UpsertGauge(models.Metrics) error
	AddCounter(models.Metrics) error
}

type Service struct {
	Repository
}

func NewService(r Repository) *Service {
	return &Service{
		Repository: r,
	}
}

func (s *Service) UpdateMetric(m models.Metrics) (err error) {
	err = m.Validate()
	if err != nil {
		return err
	}

	switch m.MType {
	case models.Gauge:
		err = s.UpsertGauge(m)
	case models.Counter:
		err = s.AddCounter(m)
	default:
		return fmt.Errorf("%w: %s", models.ErrUnknowMetricType, m.MType)
	}

	return err
}

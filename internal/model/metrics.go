package models

import (
	"errors"
	"fmt"
)

const (
	Counter = "counter"
	Gauge   = "gauge"
)

// NOTE: Не усложняем пример, вводя иерархическую вложенность структур.
// Органичиваясь плоской моделью.
// Delta и Value объявлены через указатели,
// что бы отличать значение "0", от не заданного значения
// и соответственно не кодировать в структуру.
type Metrics struct {
	ID    string   `json:"id"`
	MType string   `json:"type"`
	Delta *int64   `json:"delta,omitempty"`
	Value *float64 `json:"value,omitempty"`
	Hash  string   `json:"hash,omitempty"`
}

var (
	ErrMetricNameIsEmpty = errors.New("metric name is empty")
	ErrUnknowMetricType  = errors.New("unkown metric type")
	ErrMustNotBeNil      = errors.New("must not be nil")
)

func (m *Metrics) Validate() (err error) {
	if m.ID == "" {
		errors.Join(err, ErrMetricNameIsEmpty)
	}
	switch {
	case m.MType == Counter && m.Delta == nil:
		errors.Join(err, fmt.Errorf("%s %w", "delta", ErrMustNotBeNil))
	case m.MType == Gauge && m.Value == nil:
		errors.Join(err, fmt.Errorf("%s %w", "value", ErrMustNotBeNil))
	default:
		errors.Join(err, fmt.Errorf("%w: %s", ErrUnknowMetricType, m.MType))
	}

	return
}

type MetricName struct {
	Name string
	Type string
}

func (m *Metrics) ToMetricName() MetricName {
	return MetricName{
		Name: m.ID,
		Type: m.MType,
	}
}

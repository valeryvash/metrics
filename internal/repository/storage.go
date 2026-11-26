package repository

import (
	"sync"

	models "github.com/valeryvash/metrics/internal/model"
)

type MemStorage struct {
	mu      sync.Mutex
	storage map[models.MetricName]models.Metrics
}

func NewMemStorage() *MemStorage {
	return &MemStorage{storage: make(map[models.MetricName]models.Metrics, 0)}
}

func (m *MemStorage) UpsertGauge(g models.Metrics) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	name := g.ToMetricName()
	v := m.storage[name]
	v.Value = g.Value

	m.storage[name] = v

	return nil
}

func (m *MemStorage) AddCounter(c models.Metrics) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	name := c.ToMetricName()
	v, ok := m.storage[name]
	if !ok {
		v.Delta = c.Delta
	} else {
		*v.Delta += *c.Delta
	}
	m.storage[name] = v

	return nil
}

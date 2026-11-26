package repository_test

import (
	"github.com/valeryvash/metrics/internal/repository"
	"github.com/valeryvash/metrics/internal/service"
)

var _ service.Repository = (*repository.MemStorage)(nil)

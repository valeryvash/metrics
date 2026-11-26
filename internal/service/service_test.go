package service_test

import (
	"github.com/valeryvash/metrics/internal/handler"
	"github.com/valeryvash/metrics/internal/service"
)

var _ handler.Service = (*service.Service)(nil)
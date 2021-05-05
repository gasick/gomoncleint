package service

import "gomonclient/pkg/config"

// структура сервиса мониторинга
type Service struct {
	config *config.Config // ссылка на конфигурацию
}

// конструктор для Service
func New(cfg *config.Config) *Service {
	return &Service{
		config: cfg,
	}
}

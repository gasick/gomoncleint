package service

import (
	"gomonclient/pkg/config"
	"gomonclient/pkg/hw_listeners"
)

// структура сервиса мониторинга
type Service struct {
	config       *config.Config             // ссылка на конфигурацию
	cpuListener  *hw_listeners.CPUListener  // ссылка на слушателя CPU
	diskListener *hw_listeners.DiskListener // ссылка на слушателя CPU
	ramListener  *hw_listeners.RAMListener  // ссылка на слушателя CPU
}

// конструктор для Service
func New(cfg *config.Config) *Service {
	return &Service{
		config:       cfg,
		cpuListener:  &hw_listeners.CPUListener{Delay: cfg.CPUDelay, Timeout: cfg.Timeout},
		diskListener: &hw_listeners.DiskListener{Delay: cfg.DiskDelay, Timeout: cfg.Timeout, Path: cfg.Path},
		ramListener:  &hw_listeners.RAMListener{Delay: cfg.RAMDelay, Timeout: cfg.Timeout},
	}
}

func RunAndServe() [
	
]
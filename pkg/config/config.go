package config

import (
	"os"
	"strconv"
)

// получение значения из env и подстановка дефолтного значения
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Cтруктура для хранения конфига из env
type Config struct {
	CPUDelay    int // задержка опроса для CPU
	DiskDelay   int // задержка опроса для DISK
	RAMDelay    int // задержка опроса для RAM
	CommonDelay int // общая задержка запроса
	Timeout     int // таймаут ожидания ответа от опрашивателя
}

// конструктор для структуры конфигурации
func New() *Config {
	cpuDelay, _ := strconv.Atoi(getEnv("CPU_DELAY", "500"))
	diskDelay, _ := strconv.Atoi(getEnv("DISK_DELAY", "500"))
	ramDelay, _ := strconv.Atoi(getEnv("RAM_DELAY", "500"))
	commonDelay, _ := strconv.Atoi(getEnv("COMMON_DELAY", "500"))
	timeout, _ := strconv.Atoi(getEnv("TIMEOUT", "500"))

	return &Config{
		CPUDelay:    cpuDelay,
		DiskDelay:   diskDelay,
		RAMDelay:    ramDelay,
		CommonDelay: commonDelay,
		Timeout:     timeout,
	}
}

package hw_listeners

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// структура респонса от слушателя CPU
type CPUStatus struct {
	Count byte      `json:"cpu_count"`
	Usage []float64 `json:"usage"`
}

// реализация интерфейса Responser для CPUStatus
func (cs *CPUStatus) String() string {
	result := fmt.Sprintf("NumCPU: %v, CPU Usage: %v", cs.Count, cs.Usage)
	return result
}

// структура слушателя CPU
type CPUListener struct {
	Delay   int // задержка слушателя
	Timeout int // таймаут опроса
}

// реализация интерфейса Listener для CPUListener
func (c *CPUListener) getStatus() (CPUStatus, error) {
	usage, err := cpu.Percent(time.Microsecond*0, true)
	if err != nil {
		log.Panic("CPUListener error: ", err)
		return CPUStatus{}, err
	}

	count, err := cpu.Counts(true)
	if err != nil {
		log.Panic("CPUListener error: ", err)
		return CPUStatus{}, err
	}

	result := CPUStatus{Count: byte(count), Usage: usage}
	return result, nil
}

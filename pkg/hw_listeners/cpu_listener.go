package hw_listeners

import (
	"fmt"
	"log"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// структура респонса от слушателя CPU
type CPUStatus struct {
	Num   int
	Usage []float64
}

// реализация интерфейса Responser для CPUStatus
func (cs *CPUStatus) Read() string {
	result := fmt.Sprintf("NumCPU: %v, CPU Usage: %v", cs.Num, cs.Usage)
	return result
}

// структура слушателя CPU
type CPUListener struct {
	Delay int
}

// реализация интерфейса Listener для CPUListener
func (c *CPUListener) getStatus() (CPUStatus, error) {
	usage, err := cpu.Percent(time.Microsecond*0, true)
	if err != nil {
		log.Panic("CPUListener error: ", err)
		return CPUStatus{}, err
	}

	num, err := cpu.Counts(true)
	if err != nil {
		log.Panic("CPUListener error: ", err)
		return CPUStatus{}, err
	}

	result := CPUStatus{Num: num, Usage: usage}
	return result, nil
}

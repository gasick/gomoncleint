package hw_listeners

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

// структура респонса от слушателя Disk
type RAMStatus struct {
	Uptime    time.Time `json:"updtime"`    // uptime seconds since boot
	Loads     [3]uint64 `json:"loads"`      // 1, 5, and 15 minute load averages
	Total     float64   `json:"total"`      // Total usable main memory size
	Free      float64   `json:"free"`       // Available memory size
	Shared    float64   `json:"shared"`     // Amount of shared memory
	TotalSwap float64   `json:"total_swap"` // Total swap space size
	FreeSwap  float64   `json:"free_swap"`  // Swap space still available
}

// реализация интерфейса Responser для RAMStatus
func (rs *RAMStatus) String() string {
	result := fmt.Sprintf(`Uptime: %v, Loads: %v, Total: %v GB, 
						Free: %v GB, Shared: %v GB, TotalSwap: %v GB, FreeSwap: %v GB`,
		rs.Uptime, rs.Loads, rs.Total, rs.Free, rs.Shared, rs.TotalSwap, rs.FreeSwap)

	return result
}

// структура слушателя RAM
type RAMListener struct {
	Delay   int // задержка слушателя
	Timeout int // таймаут опроса
}

// реализация интерфейса Listener для RAMListener
func (d *RAMListener) getStatus() (RAMStatus, error) {

	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)
	if err != nil {
		log.Panic("RAMListener error:", err)
		return RAMStatus{}, err
	}

	// uptime seconds since boot
	t := time.Unix(info.Uptime, 0)

	result := RAMStatus{
		Uptime:    t,
		Loads:     info.Loads,
		Total:     float64(info.Totalram) / float64(GB),
		Free:      float64(info.Freeram) / float64(GB),
		Shared:    float64(info.Sharedram) / float64(GB),
		TotalSwap: float64(info.Totalswap) / float64(GB),
		FreeSwap:  float64(info.Freeswap) / float64(GB),
	}

	return result, nil

}

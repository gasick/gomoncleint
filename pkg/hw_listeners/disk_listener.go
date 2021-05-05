package hw_listeners

import (
	"fmt"
	"log"
	"syscall"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// структура респонса от слушателя Disk
type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

// реализация интерфейса Responser для DiskStatus
func (ds *DiskStatus) String() string {
	result := fmt.Sprintf("All: %v, Used: %v, Free: %v", ds.All, ds.Used, ds.Free)
	return result
}

// структура слушателя Disk
type DiskListener struct {
	Delay   int    // задержка слушателя
	Timeout int    // таймаут опроса
	Path    string // хз
}

// реализация интерфейса Listener для DiskListener
func (d *DiskListener) getStatus() (DiskStatus, error) {

	fs := syscall.Statfs_t{}
	err := syscall.Statfs(d.Path, &fs)
	if err != nil {
		log.Panic("DiskListener error: ", err)
		return DiskStatus{}, err
	}

	result := DiskStatus{
		All:  fs.Blocks * uint64(fs.Bsize),
		Used: fs.Bfree * uint64(fs.Bsize),
	}
	result.Used = result.All - result.Free

	return result, nil
}

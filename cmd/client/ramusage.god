package main

import (
	"syscall"
)

type MemStatus struct {
	All  uint32 `json:"all"`
	Used uint32 `json:"used"`
	Free uint32 `json:"free"`
	Self uint64 `json:"self"`
}

func RamUsage() MemStatus {
	mem := MemStatus{}
	//System occupied, only valid under linux/mac
	//system memory usage
	sysInfo := new(syscall.Sysinfo_t)
	err := syscall.Sysinfo(sysInfo)
	if err == nil {
		mem.All = sysInfo.Totalram * uint32(syscall.Getpagesize())
		mem.Free = sysInfo.Freeram * uint32(syscall.Getpagesize())
		mem.Used = mem.All - mem.Free
	}
	return mem
}

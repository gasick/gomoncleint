package main

import (
	"fmt"
	"syscall"
	"time"
)

func RamUsage() {
	info := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&info)

	if err != nil {
		fmt.Println("Error:", err)
	}
	// uptime seconds since boot
	t := time.Unix(info.Uptime, 0)
	uptime := t.Format("15:04:05 2006-01-02")

	fmt.Println("uptime:", uptime)

	// 1, 5, and 15 minute load averages
	fmt.Println("loads:", info.Loads)

	// Total usable main memory size
	fmt.Printf("total ram:%.2f GB \n", float64(info.Totalram)/float64(GB))

	// Available memory size
	fmt.Printf("free ram: %.2f GB \n", float64(info.Freeram)/float64(GB))

	// Amount of shared memory
	fmt.Printf("shared ram: %.2f GB \n", float64(info.Sharedram)/float64(GB))

	// Total swap space size
	fmt.Printf("total swap: %.2f GB \n", float64(info.Totalswap)/float64(GB))

	// Swap space still available
	fmt.Printf("free swap: %.2f GB \n", float64(info.Freeswap)/float64(GB))
}

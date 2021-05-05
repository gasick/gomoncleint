package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("Docker state: ")
		printDockerState()
		fmt.Println()
		fmt.Println("Disk usage")
		disk := DiskUsage("/")
		fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
		fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
		fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
		fmt.Println()
		fmt.Println("RamUsage")
		RamUsage()
		fmt.Println()
		fmt.Println("CpuUsage")
		CpuUsage()
		fmt.Println()
		time.Sleep(8 * time.Second)
	}
}

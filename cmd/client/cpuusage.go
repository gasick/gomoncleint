package main

import (
	"fmt"
	"runtime"
)

func CpuUsage() {
	cpu := runtime.NumCPU()
	fmt.Println("CPU amount:", cpu)
}

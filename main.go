package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	// Получение списка запуцщенных контейнеров(docker ps)
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	// Вывод всех идентификаторов контейнеров
	for _, container := range containers {
		fmt.Print(container.Names)
		fmt.Print(" -- ")
		fmt.Print(container.Image)
		fmt.Print(" -- ")
		fmt.Print(container.State)
		fmt.Print(" -- ")
		fmt.Println(container.Ports)
	}

	disk := DiskUsage("/")
	fmt.Printf("All: %.2f GB\n", float64(disk.All)/float64(GB))
	fmt.Printf("Used: %.2f GB\n", float64(disk.Used)/float64(GB))
	fmt.Printf("Free: %.2f GB\n", float64(disk.Free)/float64(GB))
}

package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func printDockerState() {
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
}

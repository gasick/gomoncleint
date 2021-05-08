package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type RunContainers struct {
	contName  []string
	contState string
	contPorts []types.Port
}

func getDockerState() []RunContainers {
	var RunContainersState []RunContainers

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
		fmt.Println("docker from main")
		fmt.Print(container.Names)
		fmt.Print(" -- ")
		fmt.Print(container.Image)
		fmt.Print(" -- ")
		fmt.Print(container.State)
		fmt.Print(" -- ")
		fmt.Println(container.Ports)
		RunContainersState = append(RunContainersState, RunContainers{
			contName:  container.Names,
			contState: container.State,
			contPorts: container.Ports,
		})

	}
	return RunContainersState
}

func checkDockerState() {
	var Reference []RunContainers
	Reference = getDockerState()

	for {
		if Reference != nil {
			var currentDockerState []RunContainers
			currentDockerState = getDockerState()
			var i = 0
			var cs = len(currentDockerState)
			var rs = len(Reference)
			if cs == rs {
				for i < rs {
					i++
					if Reference[i].contName == currentDockerState[i].contName {
						fmt.Println("Vse is gut")
					} else {
						fmt.Println("Vot i v vozduhe chemto zapahlo...")

					}
				}

			}

		}

	}
}

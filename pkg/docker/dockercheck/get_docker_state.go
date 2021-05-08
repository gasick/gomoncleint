package dockercheck

import (
	"context"
	"gomonclient/pkg/docker"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getDockerState() []docker.RunningContainer {
	var RunContainersState []docker.RunningContainer

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
		//fmt.Println("docker from main")
		//fmt.Print(container.Names)
		//fmt.Print(" -- ")
		//fmt.Print(container.Image)
		//fmt.Print(" -- ")
		//fmt.Print(container.State)
		//fmt.Print(" -- ")
		//fmt.Println(container.Ports)
		RunContainersState = append(RunContainersState, docker.RunningContainer{
			ContId:    container.ID,
			ContName:  container.Names,
			ContState: container.State,
			ContPorts: container.Ports,
		})

	}
	return RunContainersState
}

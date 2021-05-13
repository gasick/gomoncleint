package dockercheck

import (
	"fmt"
	"gomonclient/pkg/docker"
	"reflect"
	"time"
)

func CheckDockerState(delay time.Duration) {
	var Reference []docker.RunningContainer
	Reference = getDockerState()

	for {
		fmt.Println("\n\n-------------------------------")
		fmt.Println("\n\nChecking docker state:")
		if Reference != nil {
			var currentDockerState []docker.RunningContainer
			currentDockerState = getDockerState()
			dockerState := docker.StructContains(Reference, currentDockerState[0])
			fmt.Println("\n\n^^^^^^^^^^")
			fmt.Println(dockerState)
			fmt.Println("^^^^^^^^^^")
			if reflect.DeepEqual(Reference, currentDockerState) {
				for i, v := range Reference {
					if v.ContId != currentDockerState[i].ContId || v.ContState != currentDockerState[i].ContState {
						fmt.Println("\n\tSomewhat izmenenia in state")
					} else {
						fmt.Println("\n\tVse is gut")
						for _, value := range currentDockerState {
							fmt.Println(value)
						}
						fmt.Println("***")
					}
				}
			} else {
				fmt.Println("\n\tSomewhat izmenenia in amount")
				for _, value := range currentDockerState {
					fmt.Println(value)
				}
				fmt.Println("***")
			}

		} else {
			fmt.Println("\n\tThe first one izmenenia")
		}
		Reference = getDockerState()
		fmt.Println("\n\n-------------------------------")
		time.Sleep(delay)
	}
}

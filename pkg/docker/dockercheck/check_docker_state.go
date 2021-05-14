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
			fmt.Println("\n\nloop for for ")
			for _, item := range currentDockerState {
				val, dockerMessage := docker.StructContains(Reference, item)
				fmt.Println("\n\n^^^^^^^^^^")
				fmt.Printf("Difference was: %v -- %s\n", val, dockerMessage)
				fmt.Println("^^^^^^^^^^")
			}
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
			fmt.Println("\n\tThe first one izmenenia or net docker rabotaushih")
		}
		Reference = getDockerState()
		fmt.Println("\n\n-------------------------------")
		time.Sleep(delay)
	}
}

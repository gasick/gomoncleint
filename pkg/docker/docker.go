package docker

import (
	"github.com/docker/docker/api/types"
)

type RunContainers struct {
	ContId    string
	ContName  []string
	ContState string
	ContPorts []types.Port
}

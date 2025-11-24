package valueobjects

import (
	"fmt"
)

type Port struct {
	hostPort      int
	containerPort int
}

func (p Port) HostPort() int {
	return p.hostPort
}
func (p Port) ContainerPort() int {
	return p.containerPort
}

func (p Port) ToYaml() string {
	return fmt.Sprintf("%d:%d", p.HostPort(), p.ContainerPort())
}

func NewPort(hostPort int, containerPort int) Port {
	return Port{hostPort, containerPort}
}

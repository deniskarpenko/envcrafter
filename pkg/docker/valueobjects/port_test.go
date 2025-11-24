package valueobjects

import (
	"fmt"
	"testing"
)

func TestNewPort(t *testing.T) {
	hostPort := 8000
	containerPort := 80

	port := NewPort(hostPort, containerPort)

	if port.HostPort() != hostPort {
		t.Errorf("expected hostPort %d, got %d", hostPort, port.HostPort())
	}

	if port.ContainerPort() != containerPort {
		t.Errorf("expected containerPort %d, got %d", containerPort, port.ContainerPort())
	}
}

func TestPortToYaml(t *testing.T) {
	hostPort := 8000
	containerPort := 80

	port := NewPort(hostPort, containerPort)

	yaml := port.ToYaml()

	if yaml != fmt.Sprintf("%d:%d", hostPort, containerPort) {
		t.Errorf("expected yaml %s, got %s", fmt.Sprintf("%d:%d", hostPort, containerPort), yaml)
	}

}

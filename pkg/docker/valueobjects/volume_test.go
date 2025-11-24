package valueobjects

import (
	"fmt"
	"testing"
)

func TestNewVolume(t *testing.T) {
	hostPath := "/hostPath"
	containerPath := "/containerPath"

	volume := NewVolume(hostPath, containerPath)

	if volume.HostPath() != hostPath {
		t.Errorf("expected hostPath %s, got %s", hostPath, volume.HostPath())
	}

	if volume.ContainerPath() != containerPath {
		t.Errorf("expected containerPath %s, got %s", containerPath, volume.ContainerPath())
	}

	yaml := volume.toYaml()

	expectedYaml := fmt.Sprintf("%s:%s", hostPath, containerPath)

	if yaml != expectedYaml {
		t.Errorf("expected %s, got %s", expectedYaml, yaml)
	}
}

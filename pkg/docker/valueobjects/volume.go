package valueobjects

import "fmt"

type Volume struct {
	hostPath      string
	containerPath string
}

func (v *Volume) HostPath() string {
	return v.hostPath
}

func (v *Volume) ContainerPath() string {
	return v.containerPath
}

func (v *Volume) toYaml() string {
	return fmt.Sprintf("%s:%s", v.hostPath, v.containerPath)
}

func NewVolume(hostPath string, containerPath string) Volume {
	return Volume{
		hostPath:      hostPath,
		containerPath: containerPath,
	}
}

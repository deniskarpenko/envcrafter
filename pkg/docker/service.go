package docker

type ServiceBuilder interface {
	WithVolumes([]string)
	WithPorts([]string)
}

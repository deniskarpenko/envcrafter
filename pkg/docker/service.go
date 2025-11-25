package docker

import "env-crafter/pkg/docker/valueobjects"

type ServiceBuilder interface {
	WithImage(image string, tag string) ServiceBuilder
	WithBuild(context string, dockerfile string) ServiceBuilder
	WithVolumes(hostPath string, containerHost string) ServiceBuilder
	WithPorts(hostPort int, containerPort int) ServiceBuilder
	WithEnv(variable string, value string) ServiceBuilder
}

type Service struct {
	image   *valueobjects.Image
	build   *valueobjects.Build
	volumes []*valueobjects.Volume
	ports   []*valueobjects.Port
	envs    []*valueobjects.Env
	errors  []error
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) WithImage(image string, tag string) ServiceBuilder {
	imageVo, err := valueobjects.NewImage(image, tag)

	if err != nil {
		s.errors = append(s.errors, err)
	}

	s.image = &imageVo

	return s
}

func (s *Service) WithBuild(context string, dockerfile string) ServiceBuilder {
	b := valueobjects.NewBuild(context, dockerfile)

	s.build = &b

	return s
}

func (s *Service) WithVolumes(hostPath string, containerHost string) ServiceBuilder {
	v := valueobjects.NewVolume(hostPath, containerHost)

	s.volumes = append(s.volumes, &v)
	return s
}

func (s *Service) WithPorts(hostPort int, containerPort int) ServiceBuilder {
	p := valueobjects.NewPort(hostPort, containerPort)

	s.ports = append(s.ports, &p)
	return s
}

func (s *Service) WithEnv(variable string, value string) ServiceBuilder {
	e := valueobjects.NewEnv(variable, value)
	s.envs = append(s.envs, &e)
	return s
}

func (s *Service) GetImage() *valueobjects.Image {
	return s.image
}

func (s *Service) GetBuild() *valueobjects.Build {
	return s.build
}

func (s *Service) GetVolumes() []*valueobjects.Volume {
	return s.volumes
}

func (s *Service) GetPorts() []*valueobjects.Port {
	return s.ports
}

func (s *Service) GetEnvs() []*valueobjects.Env {
	return s.envs
}

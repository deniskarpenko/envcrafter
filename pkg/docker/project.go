package docker

type Project struct {
	services []*Service
}

func (p *Project) AddService(s Service) {
	p.services = append(p.services, &s)
}

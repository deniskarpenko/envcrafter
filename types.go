package main

type ImageWithTag struct {
	Image string `json:"name"`
	Tag   string `json:"tag"`
}

type ContainerConfig struct {
	Ports    []string `json:"ports"`
	Volumes  []string `json:"volumes"`
	EnvFiles [][]byte `json:"envFiles"`
	Envs     []string `json:"envs"`
}

type Service struct {
	Image  *ImageWithTag    `json:"image"`
	Config *ContainerConfig `json:"config"`
}

type Project struct {
	Backend *Service `json:"backend"`
	SQL     *Service `json:"sql"`
	Nosql   *Service `json:"nosql"`
	Web     *Service `json:"web"`
}

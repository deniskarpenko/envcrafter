package models

import "env-crafter/pkg/db/types"

type Image struct {
	ImageId      string          `gorm:"column:image_id;primary_key;autoIncrement" json:"image_id"`
	Image        string          `gorm:"column:image" json:"image"`
	DockerImage  string          `gorm:"column:docker_image" json:"docker_image"`
	ImageType    types.ImageType `gorm:"column:image_type" json:"image_type"`
	IsDockerfile bool            `gorm:"column:is_dockerfile" json:"is_dockerfile"`
}

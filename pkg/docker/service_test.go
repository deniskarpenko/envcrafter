package docker

import (
	"testing"
)

func TestServiceWithImage(t *testing.T) {
	service := NewService()

	image := "php"

	tag := "latest"

	service.WithImage(image, tag)

	imageVo := service.GetImage()

	if imageVo == nil {
		t.Errorf("Image cannot be nil")
		return
	}

	if imageVo.ImageName() != image {
		t.Errorf("Image name should be %s, but got %s", image, imageVo.ImageName())
	}

	if imageVo.Tag() != tag {
		t.Errorf("Image tag should be %s, but got %s", tag, imageVo.Tag())
	}

}

func TestService_WithBuild(t *testing.T) {
	service := NewService()

	context := "."

	dockerfile := "dockerfiles/php.Dockerfile"

	service.WithBuild(context, dockerfile)

	buildVo := service.GetBuild()

	if buildVo == nil {
		t.Errorf("Build cannot be nil")
		return
	}

	if buildVo.Context() != context {
		t.Errorf("Context should be %s, but got %s", context, buildVo.Context())
	}

	if buildVo.Dockerfile() != dockerfile {
		t.Errorf("Dockerfile should be %s, but got %s", dockerfile, buildVo.Dockerfile())
	}
}

func TestServiceWithVolumes(t *testing.T) {
	tests := []struct {
		name          string
		hostsPath     []string
		containerPath []string
	}{
		{
			"One Volume",
			[]string{"./"},
			[]string{"./var/www"},
		},
		{
			"Couple Volumes",
			[]string{"./", "./settings/php/php.ini"},
			[]string{"./var/www", "/usr/local/etc/php/php.ini"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewService()

			for index, localPath := range tt.hostsPath {
				service.WithVolumes(localPath, tt.containerPath[index])
			}

			volumes := service.GetVolumes()

			if len(volumes) != len(tt.containerPath) {
				t.Errorf("Volumes length should be %d, but got %d", len(tt.containerPath), len(volumes))
			}
		})
	}
}

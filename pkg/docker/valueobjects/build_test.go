package valueobjects

import (
	"testing"

	"gopkg.in/yaml.v3"
)

func TestNewBuild(t *testing.T) {
	context := "./app"
	dockerfile := "Dockerfile.prod"

	build := NewBuild(context, dockerfile)

	if build.Context() != context {
		t.Errorf("expected context %q, got %q", context, build.Context())
	}

	if build.Dockerfile() != dockerfile {
		t.Errorf("expected dockerfile %q, got %q", dockerfile, build.Dockerfile())
	}
}

func TestBuild_Context(t *testing.T) {
	tests := []struct {
		name    string
		context string
	}{
		{
			name:    "simple path",
			context: "./app",
		},
		{
			name:    "absolute path",
			context: "/home/user/project",
		},
		{
			name:    "empty context",
			context: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := NewBuild(tt.context, "Dockerfile")
			if got := build.Context(); got != tt.context {
				t.Errorf("Context() = %q, want %q", got, tt.context)
			}
		})
	}
}

func TestBuild_Dockerfile(t *testing.T) {
	tests := []struct {
		name       string
		dockerfile string
	}{
		{
			name:       "default dockerfile",
			dockerfile: "Dockerfile",
		},
		{
			name:       "custom dockerfile",
			dockerfile: "Dockerfile.dev",
		},
		{
			name:       "empty dockerfile",
			dockerfile: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := NewBuild(".", tt.dockerfile)
			if got := build.Dockerfile(); got != tt.dockerfile {
				t.Errorf("Dockerfile() = %q, want %q", got, tt.dockerfile)
			}
		})
	}
}

func TestBuild_ToYamlData(t *testing.T) {
	tests := []struct {
		name       string
		context    string
		dockerfile string
	}{
		{
			name:       "standard build",
			context:    "./app",
			dockerfile: "Dockerfile",
		},
		{
			name:       "custom paths",
			context:    "/custom/path",
			dockerfile: "Dockerfile.prod",
		},
		{
			name:       "empty values",
			context:    "",
			dockerfile: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			build := NewBuild(tt.context, tt.dockerfile)

			data, err := build.ToYamlData()
			if err != nil {
				t.Fatalf("ToYamlData() returned error: %v", err)
			}

			// Marshal to YAML to verify structure
			yamlBytes, err := yaml.Marshal(data)
			if err != nil {
				t.Fatalf("failed to marshal to YAML: %v", err)
			}

			// Unmarshal back to verify values
			var result struct {
				Context    string `yaml:"context"`
				Dockerfile string `yaml:"dockerfile"`
			}

			err = yaml.Unmarshal(yamlBytes, &result)
			if err != nil {
				t.Fatalf("failed to unmarshal YAML: %v", err)
			}

			if result.Context != tt.context {
				t.Errorf("Context in YAML = %q, want %q", result.Context, tt.context)
			}

			if result.Dockerfile != tt.dockerfile {
				t.Errorf("Dockerfile in YAML = %q, want %q", result.Dockerfile, tt.dockerfile)
			}
		})
	}
}

func TestBuild_ToYamlData_Structure(t *testing.T) {
	build := NewBuild("./src", "Dockerfile.test")

	data, err := build.ToYamlData()
	if err != nil {
		t.Fatalf("ToYamlData() returned error: %v", err)
	}

	// Type assertion to verify structure
	structData, ok := data.(struct {
		Context    string `yaml:"context"`
		Dockerfile string `yaml:"dockerfile"`
	})

	if !ok {
		t.Fatal("ToYamlData() returned unexpected type")
	}

	if structData.Context != "./src" {
		t.Errorf("Context = %q, want %q", structData.Context, "./src")
	}

	if structData.Dockerfile != "Dockerfile.test" {
		t.Errorf("Dockerfile = %q, want %q", structData.Dockerfile, "Dockerfile.test")
	}
}

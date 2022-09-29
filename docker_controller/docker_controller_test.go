package docker_controller

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestBuildImageWithRelativePath(t *testing.T) {
	controller, err := NewController()
	if err != nil {
		t.Error(err)
	}

	err = controller.BuildLocalImage("./dockerfile", "default")
	if err != nil {
		t.Error(err)
	}
}

func TestBuildImageWithAbsolutePath(t *testing.T) {
	controller, err := NewController()
	if err != nil {
		t.Error(err)
	}

	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	pathElements := []string{path, "dockerfile"}
	concatenatedPath := strings.Join(pathElements, "/")
	fmt.Println(concatenatedPath)

	err = controller.BuildLocalImage(concatenatedPath, "default")
}

func TestPullImage(t *testing.T) {
	controller, err := NewController()
	if err != nil {
		t.Error(err)
	}

	err = controller.PullImage("alpine")
	if err != nil {
		t.Error(err)
	}
}

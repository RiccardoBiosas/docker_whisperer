package docker_controller

import (
	"context"
	pb "docker_whisperer/dockersinventory/pb"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Controller struct {
	cli *client.Client
}

func NewController() (c *Controller, err error) {
	c = new(Controller)

	c.cli, err = client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Controller) ListImages() ([]*pb.DockerImage, error) {
	list, err := c.cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		return nil, err
	}
	var imagesTitleList []*pb.DockerImage
	for _, image := range list {
		for _, repoTag := range image.RepoTags {
			imagesTitleList = append(imagesTitleList, &pb.DockerImage{
				RepoTag: repoTag,
			})
		}

	}
	return imagesTitleList, nil
}

func (c *Controller) PullImage(image string) (err error) {
	reader, err := c.cli.ImagePull(context.Background(), image, types.ImagePullOptions{})

	if err != nil {
		return err
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	return nil
}

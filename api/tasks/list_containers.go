package tasks

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListContainers(c *client.Client) []types.Container {

	containers, err := c.ContainerList(context.Background(), types.ContainerListOptions{All: true})

	if err != nil {
		panic(err)
	}

	return containers

}

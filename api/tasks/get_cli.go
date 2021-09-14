package tasks

import "github.com/docker/docker/client"

func NewCLI(host string) (*client.Client, error) {

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}

	return cli, nil

}

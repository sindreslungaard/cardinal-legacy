package tasks

import "github.com/docker/docker/client"

func NewCLI(host string) (*client.Client, error) {

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}

	cli.Ping()

	return cli, nil

}

package tasks

import (
	"context"
	"sync"

	"github.com/docker/docker/client"
)

var cliCache = map[string]*client.Client{}
var cliCacheMu sync.Mutex

func NewCLI(host string) (*client.Client, error) {

	cliCacheMu.Lock()
	defer cliCacheMu.Unlock()

	cli, ok := cliCache[host]

	if ok {
		_, err := cli.Ping(context.Background())

		if err != nil {
			cli.Close()
			delete(cliCache, host)
		} else {
			return cli, nil
		}
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		return nil, err
	}

	cliCache[host] = cli

	return cli, nil

}

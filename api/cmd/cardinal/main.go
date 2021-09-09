package main

import (
	"cardinal/network"
	"cardinal/tasks"
	"fmt"

	"github.com/docker/docker/client"
)

func main() {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers := tasks.ListContainers(cli)

	println(fmt.Sprintf("containers: %v", containers))

	runMaster()

}

func runMaster() {
	network.ListenAndServe()
}

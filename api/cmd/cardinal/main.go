package main

import (
	"cardinal/logger"
	"cardinal/network"
)

func main() {

	logger.Info("Test")

	/* cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers := tasks.ListContainers(cli)

	println(fmt.Sprintf("containers: %v", containers)) */

	runMaster()

}

func runMaster() {
	network.ListenAndServe()
}

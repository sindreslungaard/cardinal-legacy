package main

import (
	"cardinal/network"
	"cardinal/system"
)

func main() {

	/* cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers := tasks.ListContainers(cli)

	println(fmt.Sprintf("containers: %v", containers)) */

	runMaster()

}

func runMaster() {
	go system.Process(system.NewSystem())
	network.ListenAndServe()
}

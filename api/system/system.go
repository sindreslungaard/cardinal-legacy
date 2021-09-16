package system

import (
	"cardinal/data"
	"cardinal/logger"
	"cardinal/tasks"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const (
	UpdateInterval = time.Second * 5
)

type System struct {
	State State
	Exit  chan bool
}

func NewSystem() System {
	return System{
		State: NewState(),
		Exit:  make(chan bool),
	}
}

func Process(sys System) {

	ticker := time.NewTicker(UpdateInterval)
	defer ticker.Stop()

	for {

		select {

		case <-sys.Exit:
			return

		case <-ticker.C:

			logger.Debug("Tick")

			s := NewState()
			s.Servers = UpdateServerState()
			s.Containers = UpdateContainerState(sys.State)

			sys.State = s

		}

	}

}

func UpdateServerState() map[string]ServerState {

	servers := map[string]ServerState{}

	d := data.Copy()

	for _, s := range d.Servers {

		server := ServerState{
			Uid:    s.Uid,
			Status: ServerStatusOK,
			Host:   s.Host,
		}

		servers[server.Uid] = server

	}

	// todo: ping servers

	return servers

}

func UpdateContainerState(prevState State) map[string]ContainerState {

	containers := map[string]ContainerState{}

	d := data.Copy()

	for _, c := range d.Containers {

		state := ContainerState{
			Uid:    c.Uid,
			Host:   c.Host,
			Image:  c.Image,
			Status: ContainerStatusRunning,
		}

		cli, err := tasks.NewCLI(c.Host)

		if err != nil {
			logger.Error("Unable to connect to host %s (%s)", c.Uid, c.Host)
			state.Status = ContainerStatusConflict
		} else {
			status, err := ensureContainerRuns(c, cli)
			state.Status = status

			logger.Error(err.Error())
		}

		containers[c.Uid] = state

	}

	return containers

}

func ensureContainerRuns(c data.Container, cli *client.Client) (ContainerStatus, error) {

	info, err := cli.ContainerInspect(context.Background(), c.Uid)

	if err != nil {

		// container is not currently running
		err2 := newContainer(c, cli)
		if err2 != nil {
			return "", err2
		}

		return ContainerStatusRestarting, nil

	}

	if info.State.Running {
		return ContainerStatusRunning, nil
	} else {
		return ContainerStatusConflict, fmt.Errorf("container %s is not running properly", c.Uid)
	}

}

func newContainer(c data.Container, cli *client.Client) error {

	logger.Info("Starting container %s", c.Uid)

	reader, err := cli.ImagePull(
		context.Background(),
		c.Image,
		types.ImagePullOptions{},
	)

	if err != nil {
		return fmt.Errorf("failed to pull image %s for container %s", c.Image, c.Image)
	}

	logger.Info("Pulling image %s for container %s", c.Image, c.Uid)

	ioutil.ReadAll(reader)
	reader.Close()

	config := &container.Config{
		Image:        c.Image,
		Env:          c.Env,
		ExposedPorts: nat.PortSet{},
	}

	for _, cPort := range c.Ports {
		config.ExposedPorts[nat.Port(cPort)] = struct{}{}
	}

	host := &container.HostConfig{
		PortBindings: nat.PortMap{},
	}

	for hPort, cPort := range c.Ports {
		host.PortBindings[nat.Port(cPort)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hPort,
			},
		}
	}

	resp, err := cli.ContainerCreate(
		context.Background(),
		config,
		host,
		nil,
		nil,
		c.Uid,
	)

	if err != nil {
		return fmt.Errorf("failed to create container %s with error: %s", c.Uid, err.Error())
	}

	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})

	if err != nil {
		return fmt.Errorf("failed to start container %s with error: %s", c.Uid, err.Error())
	}

	return nil

}

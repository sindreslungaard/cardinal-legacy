package system

import (
	"cardinal/data"
	"cardinal/tasks"
	"time"
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

			s := NewState()
			s.Servers = UpdateServerState()
			s.Containers = UpdateContainerState()

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

func UpdateContainerState() map[string]ContainerState {

	d := data.Copy()

	

	cli, err := tasks.NewCLI("")

}

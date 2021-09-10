package system

type ServerState struct {
	Uid    string `json:"uid"`
	Status string `json:"status"`
	Host   string `json:"host"`
}

type ContainerState struct {
	Uid    string `json:"uid"`
	Host   string `json:"host"`
	Image  string `json:"image"`
	Status string `json:"status"`
}

type State struct {
	Servers    map[string]ServerState    `json:"servers"`
	Containers map[string]ContainerState `json:"containers"`
}

func NewState() State {
	return State{
		Servers:    make(map[string]ServerState),
		Containers: make(map[string]ContainerState),
	}
}

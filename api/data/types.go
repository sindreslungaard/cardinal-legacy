package data

type ServerConfig struct {
	Uid  string `json:"uid"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type ContainerConfig struct {
	Uid      string   `json:"uid"`
	Replicas int      `json:"replicas"`
	Hosts    []string `json:"hosts"`

	// docker container settings
	Image   string            `json:"image"`
	Restart string            `json:"restart"`
	Name    string            `json:"name"`
	Env     map[string]string `json:"env"`
	Ports   map[string]string `json:"ports"`
}

type User struct {
	Uid  string `json:"uid"`
	Hash string `json:"hash"`
}

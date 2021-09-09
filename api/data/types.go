package data

type ServerConfig struct {
	uid  string `json:"uid"`
	host string `json:"host"`
	port string `json:"port"`
}

type ContainerConfig struct {
	uid      string   `json:"uid"`
	replicas int      `json:"replicas"`
	hosts    []string `json:"hosts"`

	// docker container settings
	image   string            `json:"image"`
	restart string            `json:"restart"`
	name    string            `json:"name"`
	env     map[string]string `json:"env"`
	ports   map[string]string `json:"ports"`
}

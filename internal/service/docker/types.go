package docker

type DockerPsResult struct {
	ID      string `json:"ID"`
	Image   string `json:"Image"`
	Command string `json:"Command"`
	Created string `json:"CreatedAt"`
	Status  string `json:"Status"`
	Ports   string `json:"Ports"`
	Names   string `json:"Names"`
}

type DockerStatsResult struct {
	ID       string `json:"ID"`
	CPUPerc  string `json:"CPUPerc"`
	MemUsage string `json:"MemUsage"`
	NetIO    string `json:"NetIO"`
	BlockIO  string `json:"BlockIO"`
}

type ContainerInfo struct {
	ID          string
	Image       string
	Command     string
	Created     string
	Status      string
	Ports       string
	Names       string
	CPUUsage    string
	MemoryUsage string
	NetworkIO   string
	BlockIO     string
}

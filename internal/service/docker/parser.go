package docker

func GetDockerContainerInfo(ps DockerPsResult, stat DockerStatsResult) ContainerInfo {
	return ContainerInfo{
		ID:          ps.ID,
		Image:       ps.Image,
		Command:     ps.Command,
		Created:     ps.Created,
		Status:      ps.Status,
		Ports:       ps.Ports,
		Names:       ps.Names,
		CPUUsage:    stat.CPUPerc,
		MemoryUsage: stat.MemUsage,
		NetworkIO:   stat.NetIO,
		BlockIO:     stat.BlockIO,
	}
}

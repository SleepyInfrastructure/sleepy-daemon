package main

import (
	"sync"
	"time"
)

type HandlerSnapshot struct {
	Timestamp       time.Time
	RawCPUUsage     CPUUsageRaw
	RawDiskUsages   []DiskUsageRaw
	NetworkUsage    NetworkUsage
	ContainerUsages []ContainerUsage
}

type HandlerCache struct {
	DockerInfo        DockerInfo
	Containers        []Container
	ContainerProjects []ContainerProject
}

func InitSnapshot(handler *Handler) {
	handler.LastSnapshot.Timestamp = time.Now()
	handler.LogManager.Containers = make(map[string]DaemonLogItem)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.LastSnapshot.RawCPUUsage = GetCPUUsage()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.LastSnapshot.RawDiskUsages = GetDiskUsages()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.LastSnapshot.NetworkUsage = GetNetworkUsage()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.LastCache.Containers, handler.LastCache.ContainerProjects = GetContainers(handler)
		handler.LastSnapshot.ContainerUsages = GetContainerUsages(handler)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.LastCache.DockerInfo = GetDockerInfo(handler)
	}()
	wg.Wait()

	elapsed := time.Since(handler.LastSnapshot.Timestamp)
	SleepyLogLn("Built initial snapshot! (took %v ms)", elapsed.Milliseconds())
}

package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Backend struct {
	proxy       *httputil.ReverseProxy
	containerID string
}

type ServiceRegistry struct {
	BackendsStore atomic.Value
}

func (s *ServiceRegistry) Init() {
	s.BackendsStore.Store([]Backend{})
}

func (s *ServiceRegistry) Add(containerID, addr string) {
	URL, _ := url.Parse(addr)
	s.BackendsStore.Swap(append(s.GetBackends(), Backend{
		proxy:       httputil.NewSingleHostReverseProxy(URL),
		containerID: containerID,
	}))
}

func (s *ServiceRegistry) GetByContainerID(containerId string) (Backend, bool) {
	for _, b := range s.GetBackends() {
		if b.containerID == containerId {
			return b, true
		}
	}
	return Backend{}, false
}

func (s *ServiceRegistry) GetByIndex(index int) Backend {
	return s.GetBackends()[index]
}

func (s *ServiceRegistry) RemoveByContainerID(containerID string) {
	var backends []Backend
	for _, b := range s.GetBackends() {
		if b.containerID == containerID {
			continue
		}
		backends = append(backends, b)
	}
	s.BackendsStore.Store(backends)
}

func (s *ServiceRegistry) RemoveALl() {
	s.BackendsStore.Store([]Backend{})
}

func (s *ServiceRegistry) len() int {
	return len(s.GetBackends())
}

func (s *ServiceRegistry) List() {
	backends := s.GetBackends()
	for i := range backends {
		fmt.Println(backends[i].containerID)
	}
}

func (s *ServiceRegistry) GetBackends() []Backend {
	return s.BackendsStore.Load().([]Backend)
}

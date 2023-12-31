package cache

import (
	"k8s/object"
)

// ----------------POD CACHE---------------------

type PodCache struct {
	PodStorage    object.PodStorage
	ContainerMeta map[string]*ContainerMeta
}

type ContainerMeta struct {
	Name        string
	ContainerID string
	InitialName string
	Limit       object.ContainerResources
}

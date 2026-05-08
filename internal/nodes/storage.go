package nodes

import (
	"github.com/HavelCTF/ProxmoxSDK/internal/nodes/storage"
)

// Storage returns a service scoped to a particular storage on this node.
func (s *NodeService) Storage(name string) *storage.StorageService {
	return storage.New(
		storage.StorageContext{
			Client:  s.c,
			Node:    s.node,
			Storage: name,
		},
	)
}

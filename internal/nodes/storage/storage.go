// Package storage provides functions for the Proxmox API
// /nodes/{node}/storage/{storage} endpoints (content listing, template upload).
package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/client"
	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

type StorageService struct {
	c       *client.Client
	node    string
	storage string
}

type StorageContext struct {
	Client  *client.Client
	Node    string
	Storage string
}

func New(ctx StorageContext) *StorageService {
	return &StorageService{
		c:       ctx.Client,
		node:    ctx.Node,
		storage: ctx.Storage,
	}
}

func (s *StorageService) Node() string {
	return s.node
}

func (s *StorageService) Storage() string {
	return s.storage
}

// GetContent lists volumes (templates, ISOs, backups, ...) on the storage,
// optionally filtered by content type ("vztmpl", "iso", "backup", "rootdir", ...).
// Pass an empty string to skip the filter and return all content types.
func (s *StorageService) GetContent(contentType string) (*types.StorageContentResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	endpoint := fmt.Sprintf("/nodes/%s/storage/%s/content", s.node, s.storage)
	if contentType != "" {
		endpoint = fmt.Sprintf("%s?content=%s", endpoint, contentType)
	}

	return http.DoRequest[types.StorageContentResponse](ctx, s.c,
		http.RequestContent{
			Method:   "GET",
			Endpoint: endpoint,
		},
	)
}

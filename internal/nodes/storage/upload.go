package storage

import (
	"context"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/HavelCTF/ProxmoxSDK/internal/http"
	"github.com/HavelCTF/ProxmoxSDK/types"
)

// UploadTemplate uploads an LXC template (.tar.* / .tar.zst) to the storage
// via POST /nodes/{node}/storage/{storage}/upload as multipart/form-data with
// content="vztmpl" and the given filename. Pass size as a hint (>= 0) when known;
// it is forwarded as the Proxmox "size" field but does not constrain the actual
// payload length.
//
// Note: retries are disabled for this upload — go-retryablehttp cannot replay
// the streaming body, so a partial failure surfaces directly to the caller.
func (s *StorageService) UploadTemplate(filename string, body io.Reader, size int64) (*types.UploadResponse, error) {
	// Long timeout: Proxmox writes the template to disk before returning the UPID.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
	defer cancel()

	fields := map[string]string{
		"content":  "vztmpl",
		"filename": filename,
	}
	if size > 0 {
		fields["size"] = strconv.FormatInt(size, 10)
	}

	endpoint := fmt.Sprintf("/nodes/%s/storage/%s/upload", s.node, s.storage)
	return http.DoMultipartRequest[types.UploadResponse](ctx, s.c, endpoint, fields, "filename", filename, body)
}

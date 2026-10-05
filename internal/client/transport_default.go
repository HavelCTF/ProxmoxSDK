//go:build !(js && wasm)

// SPDX-License-Identifier: MIT
// Copyright (c) 2026 HavelCTF

package client

import (
	"crypto/tls"
	"net/http"
	"time"
)

// newPlatformHTTPClient returns an HTTP client configured for native platforms
// with custom TLS settings (InsecureSkipVerify for self-signed Proxmox certs).
func newPlatformHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 60 * time.Second,
	}
}

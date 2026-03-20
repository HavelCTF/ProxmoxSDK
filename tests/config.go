package tests

import (
	"net/http"

	"github.com/HavelCTF/ProxmoxSDK/pkg/proxmox"
	"github.com/h2non/gock"
)

const (
	testURI   string = "http://test.localhost"
	testToken string = "PVEAPIToken=root@pam!testtoken=testtoken"
	testUUID  string = "test-0"
)

func initTestClient() *proxmox.Client {
	httpClient := http.DefaultClient
	gock.InterceptClient(httpClient)

	mockClient := proxmox.NewClient(testURI, testToken, testUUID,
		proxmox.WithHTTPClient(httpClient))
	return mockClient
}

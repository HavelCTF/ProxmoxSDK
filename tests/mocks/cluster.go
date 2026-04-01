package mocks

import "github.com/h2non/gock"

func ClusterNextId(baseUrl string) {
	gock.New(baseUrl).Get("/cluster/nextid").Reply(200).JSON(`{"data": "100"}`)
}

func ClusterTasks(baseUrl string) {
	gock.New(baseUrl).
		Get("/cluster/tasks").
		Reply(200).
		JSON(`{"data":[{"upid":"UPID:node1:00000001:00000001:00000001:test:running:root@pam:"}]}`)
}

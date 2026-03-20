package mocks

import (
	"github.com/h2non/gock"
)

func Tasks(baseUrl string) {
	gock.New(baseUrl).
		Get("/nodes/node1/tasks").
		Reply(200).
		JSON(`
{
	"data": [
		{
			"upid": "UPID:node1:00000001:00000001:00000001:aptupdate::root@pam:",
			"starttime": 1773973186,
			"id": "",
			"status": "OK",
			"user": "root@pam",
			"node": "node1",
			"endtime": 1773973190,
			"pid": 3256843,
			"type": "aptupdate",
			"pstart": 246674507
		},
		{
			"status": "OK",
			"node": "node1",
			"user": "root@pam",
			"upid": "UPID:node1:00000002:00000002:00000002:aptupdate::root@pam:",
			"id": "",
			"starttime": 1773888046,
			"pid": 2988769,
			"endtime": 1773888051,
			"pstart": 238160552,
			"type": "aptupdate"
		}
	]
}`)
}

func TasksStatus(baseUrl string) {
	gock.New(baseUrl).
		Persist().
		Get("/nodes/node1/tasks/UPID:node1:00000001:00000001:00000001:test:running:root@pam:/status").
		Reply(200).
		JSON(`
{
    "data": {
        "status": "running",
        "upid": "UPID:node1:00000001:00000001:00000001:test:running:root@pam:",
        "type": "test",
        "id": "running",
        "user": "root@pam",
        "node": "node1",
        "pid": 1,
        "pstart": 1,
        "starttime": 1693252591
    }
}`)

	gock.New(baseUrl).
		Persist().
		Get("/nodes/node1/tasks/UPID:node1:00000002:00000002:00000002:test:completed:root@pam:/status").
		Reply(200).
		JSON(`
{
	"data": {
		"status": "stopped",
		"exitstatus": "OK",
		"upid": "UPID:node1:00000002:00000002:00000002:test:completed:root@pam:",
		"type": "test",
		"id": "completed",
		"user": "root@pam",
		"node": "node1",
		"pid": 2,
		"pstart": 2,
		"starttime": 1693252591,
		"endtime": 1693252600
	}
}`)

	gock.New(baseUrl).
		Persist().
		Get("/nodes/node1/tasks/UPID:node1:00000003:00000003:00000003:test:failed:root@pam:/status").
		Reply(200).
		JSON(`
{
    "data": {
        "status": "stopped",
        "exitstatus": "some error occurred",
        "upid": "UPID:node1:00000003:00000003:00000003:test:failed:root@pam:",
        "type": "test",
        "id": "failed",
        "user": "root@pam",
        "node": "node1",
        "pid": 3,
        "pstart": 3,
        "starttime": 1693252591,
        "endtime": 1693252600
    }
}`)
}

func StopTask(baseUrl string) {
	gock.New(baseUrl).
		Persist().
		Delete("/nodes/node1/tasks/UPID:node1:00000001:00000001:00000001:test:running:root@pam:").
		Reply(200).
		JSON(`
{
	"data": null
}`)

	gock.New(baseUrl).
		Persist().
		Delete("/nodes/node1/tasks/UPID:node1:00000001:00000001:00000003:test::root@pam:").
		Reply(400).
		JSON(`
{
	"errors": {
		"upid": "no such task"
	},
	"message": "Parameter verification failed.\n",
	"data": null
}`)

	gock.New(baseUrl).
		Persist().
		Delete("/nodes/node1/tasks/UPID:node1:00000001:0000001:0000000Z:test::root@pam:").
		Reply(400).
		JSON(`
{
    "errors": {
        "upid": "unable to parse worker upid"
    },
    "data": null,
    "message": "Parameter verification failed.\n"
}`)
}

package mocks

import "github.com/h2non/gock"

func LXC(baseUrl string) {
	gock.New(baseUrl).
		Get("/nodes/node1/lxc").
		Reply(200).
		JSON(`
{
    "data": [
        {
            "vmid": 100,
            "status": "running",
            "name": "ct-test-1",
            "cpus": 2,
            "maxmem": 2147483648,
            "maxdisk": 10737418240,
            "maxswap": 536870912,
            "uptime": 12345,
            "tags": "prod;web"
        },
        {
            "vmid": 101,
            "status": "stopped",
            "name": "ct-test-2",
            "cpus": 1,
            "maxmem": 1073741824,
            "maxdisk": 8589934592,
            "maxswap": 268435456,
            "uptime": 0,
            "tags": "tag1;tag2"
        },
        {
            "vmid": 102,
            "status": "running",
            "name": "ct-test-3",
            "cpus": 4,
            "maxmem": 4294967296,
            "maxdisk": 21474836480,
            "maxswap": 1073741824,
            "uptime": 54321,
            "tags": ""
        }
    ]
}`)
}

func PostLXC(baseUrl string) {
	gock.New(baseUrl).
		Post("/nodes/node1/lxc").
		Reply(200).
		JSON(`{"data": "UPID:node1:00001234:00005678:5A3B7C8D:vzcreate:101:root@pam:"}`)
}

func LXCStart(baseUrl string) {
	gock.New(baseUrl).
		Post("/nodes/node1/lxc/101/status/start").
		Reply(200).
		JSON(`{"data": "UPID:node1:00001234:00005678:5A3B7C8D:vzstart:101:root@pam:"}`)
}

func LXCStop(baseUrl string) {
	gock.New(baseUrl).
		Post("/nodes/node1/lxc/101/status/stop").
		Reply(200).
		JSON(`{"data": "UPID:node1:00001234:00005678:5A3B7C8D:vzstop:101:root@pam:"}`)
}

func DeleteLXC(baseUrl string) {
	gock.New(baseUrl).
		Delete("/nodes/node1/lxc/101").
		Reply(200).
		JSON(`{"data": "UPID:node1:00001234:00005678:5A3B7C8D:vzdestroy:101:root@pam:"}`)
}

func CloneLXC(baseUrl string) {
	gock.New(baseUrl).
		Post("/nodes/node1/lxc/101/clone").
		Reply(200).
		JSON(`{"data": "UPID:node1:00001234:00005678:5A3B7C8D:vzmigrate:101:root@pam:"}`)
}

func LXCStatusCurrent(baseUrl string) {
	gock.New(baseUrl).
		Get("/nodes/node1/lxc/100/status/current").
		Reply(200).
		JSON(`
{
    "data": {
        "name": "ct-test-1",
        "status": "running",
        "vmid": 100,
        "uptime": 12345,
        "cpus": 2,
        "cpu": 0.05,
        "mem": 268435456,
        "maxmem": 2147483648,
        "disk": 1073741824,
        "maxdisk": 10737418240,
        "swap": 0,
        "maxswap": 536870912,
        "netin": 1024,
        "netout": 2048,
        "diskread": 0,
        "diskwrite": 0,
        "tags": "prod;web",
        "type": "lxc"
    }
}`)
}

func LXCStatusCurrentStopped(baseUrl string) {
	gock.New(baseUrl).
		Get("/nodes/node1/lxc/101/status/current").
		Reply(200).
		JSON(`
{
    "data": {
        "name": "ct-test-2",
        "status": "stopped",
        "vmid": 101,
        "uptime": 0,
        "cpus": 1,
        "maxmem": 1073741824,
        "maxdisk": 8589934592,
        "type": "lxc"
    }
}`)
}

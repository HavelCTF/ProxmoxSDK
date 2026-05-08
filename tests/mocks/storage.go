package mocks

import "github.com/h2non/gock"

func StorageContent(baseUrl string) {
	gock.New(baseUrl).
		Get("/nodes/node1/storage/local/content").
		MatchParam("content", "vztmpl").
		Reply(200).
		JSON(`
{
    "data": [
        {
            "volid": "local:vztmpl/debian-12-standard_12.7-1_amd64.tar.zst",
            "content": "vztmpl",
            "format": "tzst",
            "size": 134217728,
            "ctime": 1696000000
        },
        {
            "volid": "local:vztmpl/ubuntu-22.04-standard_22.04-1_amd64.tar.zst",
            "content": "vztmpl",
            "format": "tzst",
            "size": 167772160,
            "ctime": 1696500000
        }
    ]
}`)
}

func StorageUpload(baseUrl string) {
	gock.New(baseUrl).
		Post("/nodes/node1/storage/local/upload").
		Reply(200).
		JSON(`{"data": "UPID:node1:00001234:00005678:5A3B7C8D:imgcopy::root@pam:"}`)
}

package mocks

import "github.com/h2non/gock"

const versionJSON string = `
{
    "data": {
        "repoid": "9a1b2c3d",
        "release": "9.1",
        "version": "9.1-1"
    }
}`

func Version(baseUrl string) {
	gock.New(baseUrl).Get("/version").Reply(200).JSON(versionJSON)
}

func VersionWithServiceUnavailable(baseUrl string) {
	gock.New(baseUrl).Get("/version").Reply(500)
	gock.New(baseUrl).Get("/version").Reply(500)
	gock.New(baseUrl).Get("/version").Reply(200).JSON(versionJSON)
}

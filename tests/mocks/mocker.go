// SPDX-License-Identifier: MIT
// Copyright (c) 2026 HavelCTF

package mocks

import (
	"embed"
	"fmt"
	"strings"

	"github.com/h2non/gock"
)

//go:embed testdata/*/*.json
var testdataFS embed.FS

type Mocker struct {
	BaseUrl string
	Method  string
	Route   string
	Status  int
	Fixture string
}

func (m Mocker) StartMocker() error {
	if m.BaseUrl == "" || m.Method == "" || m.Route == "" ||
		(m.Status < 100 || m.Status > 599) || m.Fixture == "" {
		return fmt.Errorf("mocks: invalid mocker configuration: %+v", m)
	}

	jsonData, err := testdataFS.ReadFile("testdata/" + m.Fixture)
	if err != nil {
		return fmt.Errorf("mocks: unable to read fixture %q: %v", m.Fixture, err)
	}

	request := gock.New(m.BaseUrl)
	request.Method = strings.ToUpper(m.Method)
	request.Path(m.Route).
		Reply(m.Status).
		JSON(jsonData)
	return nil
}

func InitMocker(baseUrl, method, route string, status int, fixture string) error {
	mocker := Mocker{
		BaseUrl: baseUrl,
		Method:  method,
		Route:   route,
		Status:  status,
		Fixture: fixture,
	}
	return mocker.StartMocker()
}

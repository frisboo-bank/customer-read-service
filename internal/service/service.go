package service

import (
	"frisboo-bank/pkg/application"
	"frisboo-bank/pkg/environment"
)

var (
	Name    string
	Version string
	Build   string
	Commit  string
)

func Start() error {
	_, err := application.NewApplication(environment.Development)
	if err != nil {
		return err
	}

	return nil
}

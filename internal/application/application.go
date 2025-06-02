package application

import "frisboo-bank/pkg/application"

var (
	Name    string
	Version string
	Build   string
	Commit  string
)

func Run() {
	application.NewApplication(environment.Environment)
}

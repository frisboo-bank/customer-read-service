package customers

import (
	"frisboo-bank/pkg/container/dependencies/module"
)

func ModuleFunc() module.Module {
	m := module.ModuleFunc("customers")

	return m
}

package module

import (
	"context"
	applicationContracts "frisboo-bank/pkg/application/contracts"
)

type Module struct{}

func Boostrap(ctx context.Context, app applicationContracts.Application) (err error) {
	container := di.NewContainer()
}

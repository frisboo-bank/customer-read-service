package main

import (
	"errors"
	"fmt"
	"os"

	migrationCommand "frisboo-bank/pkg/db/migration/command"
	migrationConfig "frisboo-bank/pkg/db/migration/config"
)

func main() {
	err := migrationCommand.NewMigrationCommand().Execute()
	if err != nil {
		fmt.Print(errors.Join(migrationConfig.ErrMigrationFailed, err))
		os.Exit(1)
	}
}

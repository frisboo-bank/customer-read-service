package main

import (
	"fmt"

	"frisboo-bank/customers-service/internal/shared/app"
	"frisboo-bank/pkg/application/command"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

// @contact.name frisboo
// @contact.email contact@frisboo-bank.com
func main() {
	cli := command.NewApplicationCli(&command.ApplicationCliConfig{
		Name: "customers-service",
		Header: func() {
			fmt.Println("")
			_ = pterm.DefaultBigText.WithLetters(
				putils.LettersFromStringWithStyle("Customers", pterm.FgMagenta.ToStyle()),
				putils.LettersFromStringWithStyle(" Service", pterm.FgLightBlue.ToStyle()),
			).Render()
		},
		Description: "services to get customers",
	})

	if err := cli.Execute(func(cmd *cobra.Command, args []string) {
		if err := app.NewBootstrap().Run(); err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}
}

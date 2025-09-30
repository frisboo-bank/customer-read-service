package main

import (
	"fmt"

	"frisboo-bank/customers-service/internal/shared/app"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "customers-service",
	Short:            "services to get or manipulate customers",
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := app.NewBootstrap().Run(); err != nil {
			panic(err)
		}
	},
}

// @contact.name frisboo
// @contact.email contact@frisboo.com
func main() {
	fmt.Println("")
	_ = pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("Customers", pterm.FgMagenta.ToStyle()),
		putils.LettersFromStringWithStyle(" Service", pterm.FgLightBlue.ToStyle()),
	).Render()

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

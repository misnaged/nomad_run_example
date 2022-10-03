package serve

import (
	"microservice/internal"

	"github.com/spf13/cobra"
)

func Cmd(app *internal.App) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Run Application",
		RunE: func(cmd *cobra.Command, args []string) error {

			return app.InitServer()
		},
	}
}

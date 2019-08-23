package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/trivago/grafana-datasource-to-yaml/pkg/services"
)

// rootCmd represents the base command when called without any subcommands
func newRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grafana-datasource-to-yaml",
		Short: "Exporter of grafana datasources to YAML",
		Long:  `A tool for export datasources from the existing Grafana DB into a YAML provisioning file`,
	}

	out := cmd.OutOrStdout()

	cmd.AddCommand(
		newExportCmd(out, services.NewGrafanaExporter()),
	)

	cmd.PersistentFlags().Parse(args)

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := newRootCmd(os.Args[1:]).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

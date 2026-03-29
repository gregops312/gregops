package cmd

import (
	"fmt"
	output "gregops/pkg"

	"github.com/spf13/cobra"
)

var Version = "0.1.0"

const versionCmdName = "version"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   versionCmdName,
	Short: fmt.Sprintf("Print the version number of %s", CliName),
	Long:  fmt.Sprintf("All software has versions. This is %s's version.", CliName),
	Run: func(cmd *cobra.Command, args []string) {
		formatStr, _ := cmd.Flags().GetString("output")

		format, err := output.ParseFormat(formatStr)
		if err != nil {
			fmt.Printf("Invalid format: %v\n", err)
			return
		}

		formatter := output.NewWithWriter(format, cmd.OutOrStdout())

		data := map[string]interface{}{
			"Name":    CliName,
			"Version": Version,
		}

		if err := formatter.PrintKeyValue(data); err != nil {
			if perr := formatter.PrintError(err); perr != nil {
				cmd.PrintErrln(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().StringP("output", "o", "text", "Output format (text, json)")
}

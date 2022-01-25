package cmd

import (
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Render a go text template",
}

func init() {
	rootCmd.AddCommand(textCmd)

	textCmd.PersistentFlags().StringVarP(&inputTemplate, "template", "t", "", "A golang template to be rendered")
	textCmd.PersistentFlags().StringVarP(&inputData, "input", "i", "", "A structured file or string that represents the data values to be used in rendering the template")
	textCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "A file to which the rendered template should written; defaults to stdout")
}

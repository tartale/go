package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tartale/go/tools/render/pkg"
)

var fromJSONCmd = &cobra.Command{
	Use:   "from-json",
	Short: "Renders a template using a JSON document as the data value inputs",
	Example: `
Render text from-json -t myGoTemplate.gotmpl -i myInput.json
Render text from-json -t myGoTemplate.gotmpl -i '{"foo": "bar"}'
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return pkg.RenderTextFromJSON(inputTemplate, inputData, output)
	},
}

func init() {
	textCmd.AddCommand(fromJSONCmd)
}

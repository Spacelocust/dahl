package template

import (
	"github.com/spf13/cobra"

	"github.com/Spacelocust/dahl/cmd/template/run"
)

// templateCmd represents the template command
var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Template commands",
}

func init() {
	TemplateCmd.AddCommand(run.RunCmd)
}

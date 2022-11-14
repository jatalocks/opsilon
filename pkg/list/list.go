package list

import (
	"html/template"
	"os"

	"github.com/jatalocks/opsilon/internal/logger"
	"github.com/jatalocks/opsilon/internal/utils"
)

func List() {
	actions := utils.ConfigPopulateWorkflows()

	tmpl := `{{range .}}
	--------- {{.Name}} ----------
Description: 
{{.Workflow.Description}}
Input: {{range .Workflow.Input}}
- {{.Name}} {{ if .Default}}({{.Default}}){{end}}{{end}}
{{end}}`

	t := template.Must(template.New("tmpl").Parse(tmpl))
	err := t.Execute(os.Stdout, actions)
	logger.HandleErr(err)

}
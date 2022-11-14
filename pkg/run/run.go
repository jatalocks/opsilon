package run

import (
	"fmt"
	"html/template"
	"os"

	"github.com/fatih/color"
	"github.com/jatalocks/opsilon/internal/config"
	"github.com/jatalocks/opsilon/internal/engine"
	"github.com/jatalocks/opsilon/internal/logger"
	"github.com/jatalocks/opsilon/internal/utils"
	"github.com/manifoldco/promptui"
)

func Select() {
	actions := utils.ConfigPopulateWorkflows()
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\u25B6\uFE0F {{ .Name | cyan }} ({{ .Workflow.Description | green }})",
		Inactive: "  {{ .Name | cyan }} ({{ .Workflow.Description | yellow }})",
		Selected: "\u25B6\uFE0F {{ .Name | cyan }}",
	}

	prompt := &promptui.Select{
		Label:     "Select Action",
		Items:     actions,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	logger.HandleErr(err)
	chosenAct := actions[i]
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Printf("You Chose: %s\n", cyan(chosenAct.Name))
	PromptArguments(&chosenAct)
	toRun, err := utils.Confirm(chosenAct)
	logger.HandleErr(err)
	if toRun {
		// engine.Engine(chosenAct.Workflow)
		engine.ToGraph(chosenAct.Workflow)
	} else {
		fmt.Println("Run Canceled")
	}
}

func PromptArguments(act *config.Action) {

	argsWithValues := act.Workflow.Input
	// Each template displays the data received from the prompt with some formatting.
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ .Name }} ({{ .Default | faint }}): ",
		Valid:   "{{ .Name | green }} ({{ .Default | faint }}): ",
		Invalid: "{{ .Name | red }} ({{ .Default | faint }}): ",
		Success: "{{ .Name | bold }} ({{ .Default | faint }}): ",
	}

	for i, v := range argsWithValues {
		// The validate function follows the required validator signature.
		validate := func(input string) error {
			if input == "" && !v.Optional && v.Default == "" {
				return fmt.Errorf("This argument is mandatory")
			}
			return nil
		}

		prompt := promptui.Prompt{
			Label:     v,
			Templates: templates,
			Validate:  validate,
		}

		result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if result == "" {
			result = v.Default
		}

		// The result of the prompt, if valid, is displayed in a formatted message.
		argsWithValues[i].Default = result
		fmt.Printf("%s\n", result)
	}
	tmpl := `--------- Running "{{.Name}}" with: ----------
{{range .Workflow.Input}}
{{ .Name }}: {{ .Default }}
{{end}}
	`

	t := template.Must(template.New("tmpl").Parse(tmpl))

	err := t.Execute(os.Stdout, act)

	logger.HandleErr(err)
}
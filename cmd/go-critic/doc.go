package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"text/template"

	"github.com/go-critic/go-critic/linter"
)

// Main implements sub-command entry point.
func runDocs(_ context.Context, args []string) error {
	flagSet := flag.NewFlagSet("go-critic", flag.ContinueOnError)
	if err := flagSet.Parse(args); err != nil {
		return err
	}

	switch args := flagSet.Args(); len(args) {
	case 0:
		printShortDoc()
	case 1:
		printDoc(args[0])
	default:
		log.Fatalf("expected 0 or 1 positional arguments")
	}
	return nil
}

func printShortDoc() {
	for _, info := range linter.GetCheckersInfo() {
		fmt.Printf("%s %v\n", info.Name, info.Tags)
	}
}

func printDoc(name string) {
	tmplString := `{{.Checker.Name}} checker documentation
URL: {{.Checker.Collection.URL}}
Tags: {{.Checker.Tags}}

{{.Checker.Summary}}.
{{ if .Checker.Details }}
{{.Checker.Details}}
{{ end }}
Non-compliant code:
{{.Checker.Before}}

Compliant code:
{{.Checker.After}}
{{- if .Checker.Note }}

{{.Checker.Note}}
{{- end }}
{{- if .Checker.Params }}

Checker parameters:
{{- range $key, $_ := .Checker.Params }}
  -@{{$.Checker.Name}}.{{$key}} {{index $.ParamTypes $key}}
    	{{.Usage}} (default {{.Value}})
{{- end }}
{{- end }}
`
	templateData.ParamTypes = make(map[string]string)
	templateData.Checker = info
	for pname, p := range info.Params {
		templateData.ParamTypes[pname] = reflect.TypeOf(p.Value).String()
	}
	if err := tmpl.Execute(os.Stdout, templateData); err != nil {
		panic(fmt.Sprintf("executing checker doc template: %v", err))
	}
	info := findInfoByName(name)
	tmpl := template.Must(template.New("doc").Parse(tmplString))
	var templateData struct {
		Checker    *linter.CheckerInfo
		ParamTypes map[string]string
	}
	if info == nil {
		log.Fatalf("checker with name %q not found", name)
		return // To avoid `info can be nil` from the staticcheck
	}
}

func findInfoByName(name string) *linter.CheckerInfo {
	for _, info := range linter.GetCheckersInfo() {
		if info.Name == name {
			return info
		}
	}
	return nil
}

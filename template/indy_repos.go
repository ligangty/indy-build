package template

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

// IndyGroupVars ...
type IndyGroupVars struct {
	Name         string
	Constituents []string
}

// IndyGroupTemplate ...
func IndyGroupTemplate(indyGroupVars *IndyGroupVars) string {
	groupTemplate := `
{
  "type" : "group",
  "key" : "maven:group:{{.Name}}",
  "metadata" : {
    "changelog" : "init group {{.Name}}"
  },
  "disabled" : false,
  "constituents" : [{{range $index,$con := .Constituents}}"{{$con}}"{{if isNotLast $index $.Constituents}},{{end}}{{end}}],
  "packageType" : "maven",
  "name" : "{{.Name}}",
  "type" : "group",
  "disable_timeout" : 0,
  "path_style" : "plain",
  "authoritative_index" : false,
  "prepend_constituent" : false
}
  `
	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"isNotLast": func(index int, array []string) bool {
			return index < len(array)-1
		},
	}
	t := template.Must(template.New("settings").Funcs(funcMap).Parse(groupTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, indyGroupVars)
	if err != nil {
		log.Fatal("executing template:", err)
		os.Exit(1)
	}

	return buf.String()
}

// IndyHostedVars ...
type IndyHostedVars struct {
	Name string
}

// IndyHostedTemplate ...
func IndyHostedTemplate(indyHostedVars *IndyHostedVars) string {
	hostedTemplate := `
{
  "key" : "maven:hosted:{{.Name}}",
  "description" : "{{.Name}}",
  "metadata" : {
    "changelog" : "init hosted {{.Name}}"
  },
  "disabled" : false,
  "snapshotTimeoutSeconds" : 0,
  "readonly" : false,
  "packageType" : "maven",
  "name" : "{{.Name}}",
  "type" : "hosted",
  "disable_timeout" : 0,
  "path_style" : "plain",
  "authoritative_index" : true,
  "allow_snapshots" : true,
  "allow_releases" : true
}
	`

	t := template.Must(template.New("settings").Parse(hostedTemplate))
	var buf bytes.Buffer
	err := t.Execute(&buf, indyHostedVars)
	if err != nil {
		log.Fatal("executing template:", err)
		os.Exit(1)
	}

	return buf.String()
}

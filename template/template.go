package template

import "github.com/common-nighthawk/go-figure"

// template for help
var Help = figure.NewFigure("./dahl", "ogre", true).String() + `
{{with .Long}}{{. | trimTrailingWhitespaces}}{{end -}}
{{if ne .Use "dahl"}}Description:
  {{with .Short}}{{. | trimTrailingWhitespaces}}{{end}}{{end}}

Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

Available Commands:{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{else}}{{range $group := .Groups}}

{{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

Additional Commands:{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

// template for hello world template file
var HelloWorld = `_{ @message }__{ @emoji.smile }_

This is a template file example, named: "_{ filename }__{ extension }_"

More info at https://github.com/Spacelocust/dahl/blob/main/README.md _{ @emoji.wink }_
`

// templte configuration file for hello world template file
var ConfigFile = `## Example template 

templates:
 example:
  filename: "hello-world"
  from: "./.dahl/example/hello-world.dahl"
  to: "."
  extension: ".txt"
  props:
   message: "Hello and welcome "
   emoji:
    wink: ";)"
    smile: ":D"
`
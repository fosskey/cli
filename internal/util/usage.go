package util

func UsageTemplate() string {
	return `
Usage:{{if .Runnable}} {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

Commands:{{range $cmds}}{{if .IsAvailableCommand}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}{{end}}{{if .HasAvailableSubCommands}}

Run "{{.CommandPath}} [command] --help" for more information about a command.
{{end}}
`
}

func HelpTemplate() string {
	return `{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}{{if .Runnable}}
{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}{{end}}
{{end}}`
}

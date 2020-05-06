package consts

var PWNBUDDY_VERSION = "1.0"

var PWNBUDDY_AUTHORS = []map[string]interface{}{
	{"Name": "Muhammet Raşit AYDIN", "Email": "rasity.aydin@gmail.com"},
	{"Name": "Mert UĞUR", "Email": "ugur.mertugur@gmail.com"},
	{"Name": "Anıl BAŞ", "Email": "anilbas..."},
	{"Name": "Özgür Burak ...", "Email": "..."},
}

var USAGE_TEMPLATE = `Usage:{{if .Runnable}}
   {{.UseLine}}
{{end}}{{if .HasAvailableSubCommands}}{{if not .Parent}}
   [module] [module command] [flags]
{{else}}
   {{.CommandPath}} [command]
{{end}}{{end}}{{if gt (len .Aliases) 0}}
Aliases:
   {{.NameAndAliases}}
{{end}}{{if .HasExample}}
Examples:
   {{.Example}}
{{end}}{{if .HasAvailableSubCommands}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
   {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
{{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

var HELP_TEMPLATE = `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}
`

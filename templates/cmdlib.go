{{ if .CMD.Parent }}
package {{ .CMD.Parent.Name }}
{{ else }}
package commands
{{ end }}

func {{ .CMD.CmdName }}() {

}

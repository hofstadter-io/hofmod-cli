{{ if .CMD.Args -}}
{{ range $i, $A := .CMD.Args -}}
{{ if gt $i 0 }}, {{end }}{{ $A.argName }} {{ $A.Type -}}
{{ end -}}
{{ else -}}
args []string{{ end -}}

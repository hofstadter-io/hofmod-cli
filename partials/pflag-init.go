{{ define "pflag-init" }}
{{ if or $.Pflags $.Flags}}
func init () {
  {{ range $i, $F := $.Pflags }}
  {{ template "pflag-bind" $F }}
  {{- end }}
}
{{ end }}
{{ end }}

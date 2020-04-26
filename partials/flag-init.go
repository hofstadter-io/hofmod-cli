{{ define "flag-init" }}
{{ if $.Flags}}
func init () {
  {{ range $i, $F := $.Flags }}
  {{ template "flag-bind" . }}
  {{- end }}
}
{{ end }}
{{ end }}

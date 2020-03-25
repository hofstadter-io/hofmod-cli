package partials

FlagInit :: RealFlagInit

RealFlagInit :: """
{{ define "flag-init" }}
{{ if or $.Pflags $.Flags}}
func init () {
  {{ range $i, $F := $.Pflags }}
  {{ template "pflag-bind" $F }}
  {{- end }}
  {{ range $i, $F := $.Flags }}
  {{ template "flag-bind" . }}
  {{- end }}
}
{{ end }}
{{ end }}
"""

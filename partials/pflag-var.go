{{ define "pflag-vars" }}
{{ if $.Pflags}}
{{ $Prefix := "Root" }}
{{ if $.CMD }}{{ $Prefix = $.CMD.CmdName }}{{ end }}
var (
  {{ range $i, $P := $.Pflags }}
  {{ $Prefix }}{{ $P.FlagName }}Pflag {{ $P.Type }}
  {{- end }}
)
{{ end }}
{{ end }}

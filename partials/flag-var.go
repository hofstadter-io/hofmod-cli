{{ define "flag-vars" }}
{{ if $.Flags}}
{{ $Prefix := "Root" }}
{{ if $.CMD }}{{ $Prefix = $.CMD.CmdName }}{{ end }}
var (
  {{ range $i, $F := $.Flags }}
  {{ $Prefix }}{{ $F.FlagName }}Flag {{ $F.Type }}
  {{- end }}
)
{{ end }}
{{ end }}

package partials

FlagVar : RealFlagVar

RealFlagVar : """
{{ define "flag-vars" }}
{{ if or $.Pflags $.Flags}}
{{ $Prefix := "Root" }}
{{ if $.CMD }}{{ $Prefix = $.CMD.CmdName }}{{ end }}
var (
  {{ range $i, $P := $.Pflags }}
  {{ $Prefix }}{{ $P.FlagName }}Pflag {{ $P.Type }}
  {{- end }}
  {{ range $i, $F := $.Flags }}
  {{ $Prefix }}{{ $F.FlagName }}Flag {{ $F.Type }}
  {{- end }}
)
{{ end }}
{{ end }}


"""

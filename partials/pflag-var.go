{{ define "pflag-vars" }}
{{ if $.Pflags}}

{{ $Prefix := "Root" }}
{{ if $.Parent.Parent.Parent.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Parent.Parent.Parent.Name "__" $.CMD.Parent.Parent.Parent.Name "__" $.CMD.Parent.Parent.Name "__" $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if $.Parent.Parent.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Parent.Parent.Name "__" $.CMD.Parent.Parent.Name "__" $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if $.Parent.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Parent.Name "__" $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if $.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if $.CmdName }}
  {{ $Prefix = $.CmdName }}
{{ end }}

type {{ $Prefix }}Pflagpole struct {
  {{ range $i, $F := $.Pflags }}
  {{ $F.FlagName }} {{ $F.Type }}
  {{- end }}
}

var {{ $Prefix }}Pflags {{ $Prefix }}Pflagpole
{{ end }}
{{ end }}

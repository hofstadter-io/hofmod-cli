{{ define "flag-vars" }}
{{ if $.Flags}}

{{ $Prefix := "Root" }}
{{ if $.Parent.Parent.Parent.Parent }}
  {{ $Prefix = (print $.Parent.Parent.Parent.Parent.Name "__" $.Parent.Parent.Parent.Name "__" $.Parent.Parent.Name "__" $.Parent.Name "__" $.CmdName) }}
{{ else if $.Parent.Parent.Parent }}
  {{ $Prefix = (print $.Parent.Parent.Parent.Name "__" $.Parent.Parent.Name "__" $.Parent.Name "__" $.CmdName) }}
{{ else if $.Parent.Parent }}
  {{ $Prefix = (print $.Parent.Parent.Name "__" $.Parent.Name "__" $.CmdName) }}
{{ else if $.Parent }}
  {{ $Prefix = (print $.Parent.Name "__" $.CmdName) }}
{{ else if $.CmdName }}
  {{ $Prefix = $.CmdName }}
{{ end }}
{{ $Prefix = ( title $Prefix ) }}

type {{ $Prefix }}Flagpole struct {
  {{ range $i, $F := $.Flags }}
  {{ $F.FlagName }} {{ $F.Type }}
  {{- end }}
}

var {{ $Prefix }}Flags {{ $Prefix }}Flagpole

{{ end }}
{{ end }}

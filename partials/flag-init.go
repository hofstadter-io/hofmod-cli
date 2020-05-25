{{ define "flag-init" }}
{{ if $.Flags}}

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

func init () {
  {{ range $i, $F := $.Flags }}
	{{ $Prefix }}Cmd.Flags().{{- template "cobra-type" $F.Type -}}VarP(&(flags.{{ $Prefix }}Flags.{{ $F.FlagName }}), "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
  {{- end }}
}
{{ end }}
{{ end }}

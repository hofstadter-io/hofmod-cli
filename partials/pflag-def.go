{{- define "pflag-bind" }}
{{ $F := . }}

{{ $Prefix := "Root" }}
{{ if .CMD.Parent.Parent.Parent.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Parent.Parent.Parent.Name "__" $.CMD.Parent.Parent.Parent.Name "__" $.CMD.Parent.Parent.Name "__" $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if .CMD.Parent.Parent.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Parent.Parent.Name "__" $.CMD.Parent.Parent.Name "__" $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if .CMD.Parent.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Parent.Name "__" $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if .CMD.Parent }}
  {{ $Prefix = (print $.CMD.Parent.Name "__" $.CMD.CmdName) }}
{{ else if $.CMD }}
  {{ $Prefix = $.CMD.CmdName }}
{{ end }}

{{ $Prefix }}Cmd.PersistentFlags().{{- template "cobra-type" $F.Type -}}VarP(&pflags.{{ $Prefix }}{{ $F.FlagName }}Pflag, "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
{{ end -}}

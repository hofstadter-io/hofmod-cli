{{ define "pflag-init" }}
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

func init () {
  {{ range $i, $F := $.Pflags }}
	{{ $Prefix }}Cmd.PersistentFlags().{{- template "cobra-type" $F.Type -}}VarP(&(flags.{{ $Prefix }}Pflags.{{ $F.FlagName }}), "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
  {{- end }}
}
{{ end }}
{{ end }}

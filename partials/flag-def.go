{{- define "pflag-bind" }}
{{ $F := . }}
{{ $Prefix := "Root" }}
{{ if $.CMD }}{{ $Prefix = $.CMD.CmdName }}{{ end }}
{{ $Prefix }}Cmd.PersistentFlags().{{- template "cobra-type" $F.Type -}}VarP(&{{ $Prefix }}{{ $F.FlagName }}Pflag, "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
{{ end -}}

{{- define "flag-bind" }}
{{ $F := . }}
{{ $Prefix := "Root" }}
{{ if $.CMD }}{{ $Prefix = $.CMD.CmdName }}{{ end }}
{{ $Prefix }}Cmd.Flags().{{- template "cobra-type" $F.Type -}}VarP(&{{ $Prefix }}{{ $F.FlagName }}Flag, "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
{{ end -}}


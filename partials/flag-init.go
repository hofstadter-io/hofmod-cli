{{ define "flag-init" }}
{{ if $.Flags}}
{{ $Prefix := "Root" }}
{{ if $.CmdName }}{{ $Prefix = $.CmdName }}{{ end }}
func init () {
  {{ range $i, $F := $.Flags }}
	{{ $Prefix }}Cmd.Flags().{{- template "cobra-type" $F.Type -}}VarP(&{{ $Prefix }}{{ $F.FlagName }}Flag, "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
  {{- end }}
}
{{ end }}
{{ end }}

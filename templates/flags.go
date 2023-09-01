package flags

import (
	"github.com/spf13/pflag"
)

var _ *pflag.FlagSet

{{ if .CMD }}
{{ template "flag-setup" .CMD }}
{{ else }}
{{ template "flag-setup" .CLI }}
{{ end }}


{{ define "flag-setup" }}

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

{{ if (or $.Flags $.Pflags)}}
var {{ $Prefix }}FlagSet *pflag.FlagSet

{{ if $.Pflags }}
type {{ $Prefix }}Pflagpole struct {
  {{ range $i, $F := $.Pflags }}
  {{ $F.FlagName }} {{ $F.Type }}
  {{- end }}
}

func Setup{{ $Prefix }}Pflags(fset *pflag.FlagSet, fpole *{{ $Prefix }}Pflagpole) {
	// pflags
  {{ range $i, $F := $.Pflags }}
	fset.{{- template "cobra-type" $F.Type -}}VarP(&(fpole.{{ $F.FlagName }}), "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
  {{- end }}
}

var {{ $Prefix }}Pflags {{ $Prefix }}Pflagpole
{{ end }}

{{ if $.Flags }}
type {{ $Prefix }}Flagpole struct {
  {{ range $i, $F := $.Flags }}
  {{ $F.FlagName }} {{ $F.Type }}
  {{- end }}
}

var {{ $Prefix }}Flags {{ $Prefix }}Flagpole

func Setup{{ $Prefix }}Flags(fset *pflag.FlagSet, fpole *{{ $Prefix }}Flagpole) {
	// flags
  {{ range $i, $F := $.Flags }}
	fset.{{- template "cobra-type" $F.Type -}}VarP(&(fpole.{{ $F.FlagName }}), "{{ $F.Long }}", "{{ $F.Short }}", {{ if $F.Default}}{{$F.Default}}{{else}}{{template "go-default" $F.Type }}{{end}}, "{{ $F.Help }}")
  {{- end }}
}
{{ end }}

func init() {
	{{ $Prefix }}FlagSet = pflag.NewFlagSet("{{$Prefix}}", pflag.ContinueOnError)
{{ if $.Pflags }}
	Setup{{ $Prefix }}Pflags({{ $Prefix }}FlagSet, &{{ $Prefix }}Pflags)
{{ end }}
{{ if $.Flags }}
	Setup{{ $Prefix }}Flags({{ $Prefix }}FlagSet, &{{ $Prefix }}Flags)
{{ end }}	
}


{{ end }}
{{ end }}

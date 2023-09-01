{{ define "pflag-init" }}
{{ if $.Pflags}}

{{ $CmdName := "Root" }}
{{ if $.CmdName }}
{{ $CmdName = $.CmdName }}
{{ end }}

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

func init () {
  {{- range $i, $F := $.Pflags }}
	{{ $CmdName }}Cmd.PersistentFlags().AddFlag(flags.{{ $Prefix }}FlagSet.Lookup("{{ $F.Long }}"))
  {{- end }}
}
{{ end }}
{{ end }}

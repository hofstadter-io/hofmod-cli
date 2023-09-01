{{ define "flag-init" }}
{{ if (or $.Pflags $.Flags)}}

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
{{ if $.Pflags }}
	flags.Setup{{ $Prefix }}Pflags({{ $CmdName }}Cmd.PersistentFlags(), &(flags.{{ $Prefix }}Pflags))
{{ end }}
{{ if $.Flags }}
	flags.Setup{{ $Prefix }}Flags({{ $CmdName }}Cmd.Flags(), &(flags.{{ $Prefix }}Flags))
{{ end }}	
}
{{ end }}
{{ end }}

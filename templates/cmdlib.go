{{ if .CMD.Parent }}
package libcmd{{ .CMD.Parent.Name }}
{{ else }}
package libcmd
{{ end }}

{{ if .CMD.PersistentPrerun }}
func {{ .CMD.CmdName }}PersistentPreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CMD.PersistentPrerunBody }}
	{{ .CMD.PersistentPrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if .CMD.Prerun }}
func {{ .CMD.CmdName }}PreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CMD.PrerunBody }}
	{{ .CMD.PrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if not .CMD.OmitRun}}
func {{ .CMD.CmdName }}Run({{ template "lib-args.go" . -}}) (err error) {

	{{ if .CMD.Body}}
	{{ .CMD.Body}}
	{{ end }}

	return err
}
{{ end }}

{{ if .CMD.PersistentPostrun}}
func {{ .CMD.CmdName }}PersistentPostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CMD.PersistentPostrunBody}}
	{{ .CMD.PersistentPostrunBody}}
	{{ end }}

	return err
}
{{ end }}

{{ if .CMD.Postrun}}
func {{ .CMD.CmdName }}PostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CMD.PostrunBody }}
	{{ .CMD.PostrunBody }}
	{{ end }}

	return err
}
{{ end }}

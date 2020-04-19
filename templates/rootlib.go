package libcmd

{{ if .CLI.PersistentPrerun }}
func RootPersistentPreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CLI.PersistentPrerunBody }}
	{{ .CLI.PersistentPrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if .CLI.Prerun }}
func RootPreRun({{- template "lib-args.go" . -}}) (err error) {
	{{ if .CLI.PrerunBody }}
	{{ .CLI.PrerunBody }}
	{{ end }}

	return err
}
{{ end }}

{{ if not .CLI.OmitRun}}
func RootRun({{ template "lib-args.go" . -}}) (err error) {

	{{ if .CLI.Body}}
	{{ .CLI.Body}}
	{{ end }}

	return err
}
{{ end }}

{{ if .CLI.PersistentPostrun}}
func RootPersistentPostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CLI.PersistentPostrunBody}}
	{{ .CLI.PersistentPostrunBody}}
	{{ end }}

	return err
}
{{ end }}

{{ if .CLI.Postrun}}
func RootPostRun({{- template "lib-args.go" . -}}) (err error) {

	{{ if .CLI.PostrunBody }}
	{{ .CLI.PostrunBody }}
	{{ end }}

	return err
}
{{ end }}

package flags

{{ if .CMD }}
{{ template "pflag-vars" .CMD }}
{{ template "flag-vars" .CMD }}
{{ else }}
{{ template "pflag-vars" .CLI }}
{{ template "flag-vars" .CLI }}
{{ end }}

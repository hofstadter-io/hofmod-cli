package pflags

{{ if .CMD }}
{{ template "pflag-vars" .CMD }}
{{ else }}
{{ template "pflag-vars" .CLI }}
{{ end }}

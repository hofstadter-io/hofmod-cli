{{- define "cobra-type" -}}
{{- if eq . "string"}}String
{{- else if eq . "[]string" }}StringSlice
{{- else if eq . "int" }}Int
{{- else if eq . "[]int" }}IntSlice
{{- else if eq . "float64" }}Float64
{{- else if eq . "[]float64" }}Float64Slice
{{- else if eq . "bool" }}Bool
{{- else }}<unknown type '{{.}}'>
{{- end -}}
{{- end -}}

{{- define "go-default" -}}
{{- if eq . "string"}}""
{{- else if eq . "[]string" }}[]string{}
{{- else if eq . "int" }}0
{{- else if eq . "[]int" }}[]int{}
{{- else if eq . "float64" }}0.0
{{- else if eq . "[]float64" }}[]float64P{}
{{- else if eq . "bool" }}false
{{- else }}<unknown type '{{.}}'>
{{- end -}}
{{- end -}}

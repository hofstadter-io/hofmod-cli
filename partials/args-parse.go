{{ define "args-parse" }}
{{ $ARGS := . }}
// Argument Parsing
{{ range $i, $A := $ARGS }}
{{ if $A.Required }}
if {{ $i }} >= len(args) {
  fmt.Println("missing required argument: '{{$A.argName}}'")
  cmd.Usage()
  os.Exit(1)
}
{{ end }}
var {{ $A.argName }} {{$A.Type}}
{{ if $A.Default }}{{ $A.argName }} = {{ $A.Default }}{{end}}

if {{ $i }} < len(args) {
  {{ if $A.Rest }}
  {{ $A.argName }} = args[{{ $i }}:]

  {{ else if eq $A.Type "string" }}
  {{ $A.argName }} = args[{{ $i }}]

  {{ else if eq $A.Type "int" }}
  {{ $A.argName}}Str := args[{{ $i }}]
  var {{ $A.argName }}Err error
  {{ $A.argName }}Type, {{ $A.argName }}Err := strconv.ParseInt({{ $A.argName }}Str, 10, 64)
  if {{ $A.argName }}Err != nil {
    fmt.Printf("argument of wrong type. expected: '{{ $A.Type}}' got error: %v", {{ $A.argName }}Err )
    cmd.Usage()
    os.Exit(1)
  }
  {{ $A.argName }} = int({{ $A.argName }}Type)

  {{ end }}
}
{{ end }}
{{ end }}

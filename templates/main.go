package main

import (
	"fmt"
	"os"

  {{ if .CLI.EnablePProf }}
  "log"
	"runtime/pprof"
  {{end}}

	"{{ .CLI.Package }}/cmd"
)

func main() {
	{{ if .CLI.EnablePProf }}
	f, err := os.Create("hof-cpu.prof")
	if err != nil {
			log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	{{ end }}

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

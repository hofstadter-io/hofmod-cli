package templates

MainTemplate :: RealMainTemplate

RealMainTemplate :: """
package main

import (
	"fmt"
	"os"

  {{ if .CLI.EnablePProf }}
  // Debug & PProf
	"log"
	"net/http"
	_ "net/http/pprof"
  {{end}}

	"github.com/spf13/viper"

	"{{ .CLI.Package }}/commands"
)

func main() {
  {{ if .CLI.EnablePProf }}
  // TODO, turn this into a flag and run if enabled, in root command persistent-pre-run
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
  {{ end }}

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
"""

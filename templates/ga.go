package ga

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/hofstadter-io/yagu"

	"{{ .CLI.Package }}/verinfo"
)

func SendCommandPath(cmd string) {
	cs := strings.Fields(cmd)
	c := strings.Join(cs[1:], "/")
	SendGaEvent(c, "", 0)
}

func SendGaEvent(action, label string, value int) {
	if os.Getenv("{{ .CLI.CLI_NAME }}_TELEMETRY_DISABLED") != "" {
		return
	}

	cid, err := readGaId()
	if err != nil {
		cid = "unknown"
	}

	ua := fmt.Sprintf(
		"%s/%s %s/%s",
		"{{ .CLI.cliName }}", verinfo.Version,
		verinfo.BuildOS, verinfo.BuildArch,
	)

	cfg := yagu.GaConfig{
		TID: "{{ .CLI.Telemetry }}",
		CID: cid,
		UA: ua,
		CN: "{{ .CLI.cliName }}",
		CS: "{{ .CLI.cliName }}/" + verinfo.Version,
		CM: verinfo.Version,
	}

	evt := yagu.GaEvent{
		Source: cfg.UA,
		Category: "{{ .CLI.cliName }}",
		Action: action,
		Label: label,
	}

	if value >= 0 {
		evt.Value = value
	}

	yagu.SendGaEvent(cfg, evt)
}

func readGaId() (string, error) {
	// ucd := yagu.UserHomeDir()
	ucd, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(ucd, "{{ .CLI.TelemetryIdDir }}")
	fn := filepath.Join(dir, ".uuid")

	_, err = os.Lstat(fn)
	if err != nil {
		// file does not exist, probably...
		return writeGaId()
	}

	content, err := ioutil.ReadFile(fn)
	if err != nil {
		return writeGaId()
	}

	return string(content), nil
}

func writeGaId() (string, error) {

	ucd, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(ucd, "{{ .CLI.TelemetryIdDir }}")
	err = yagu.Mkdir(dir)
	if err != nil {
		return "", err
	}

	fn := filepath.Join(dir, ".uuid")

	id, err := uuid.NewUUID()
	if err != nil {
		return id.String(), err
	}

	err = ioutil.WriteFile(fn, []byte(id.String()), 0644)
	if err != nil {
		return id.String(), err
	}

	return id.String(), nil
}

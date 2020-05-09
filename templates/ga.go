package ga

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hofstadter-io/yagu"

	"{{ .CLI.Package }}/verinfo"
)

func SendGaEvent(action, label string, value int) {
	// Do something here to lookup / create
	// cid := "13b3ad64-9154-11ea-9eba-47f617ab74f7"
	cid, err := readGaId()
	if err != nil {
		fmt.Println("Error", err)
		cid = "unknown"
	}

	fmt.Println("CID:", cid)


	cfg := yagu.GaConfig{
		TID: "{{ .CLI.Telemetry }}",
		CID: cid,
		UA: "{{ .CLI.cliName }}/" + verinfo.Version,
		CN: "{{ .CLI.cliName }}",
		CS: "{{ .CLI.cliName }}",
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
	ucd := yagu.UserHomeDir()
	dir := filepath.Join(ucd, "{{ .CLI.TelemetryIdDir }}")
	fn := filepath.Join(dir, ".uuid")

	_, err := os.Lstat(fn)
	if err != nil {
		// make sure we check err for something actually bad
		if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
			// return "", err
		}
		// file does not exist
		return writeGaId()
	}

	content, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println("ERROR:", err)
		return writeGaId()
		// return "", err
	}

	return string(content), nil
}

func writeGaId() (string, error) {
	fmt.Println("writeGaId")
	ucd := yagu.UserHomeDir()

	dir := filepath.Join(ucd, "{{ .CLI.TelemetryIdDir }}")
	err := yagu.Mkdir(dir)
	if err != nil {
		return "", err
	}

	fmt.Println("Mkdir:", dir)

	fn := filepath.Join(dir, ".uuid")

	id, err := uuid.NewUUID()
	if err != nil {
		return id.String(), err
	}

	fmt.Println("writeGaId: ", id.String())

	err = ioutil.WriteFile(fn, []byte(id.String()), 0644)
	if err != nil {
		return id.String(), err
	}

	return id.String(), nil
}

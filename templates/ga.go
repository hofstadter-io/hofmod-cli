package ga

import (
	"github.com/hofstadter-io/yagu"

	"{{ .CLI.Package }}/verinfo"
)

func SendGaEvent(action, label string, value int) {
	// Do something here to lookup / create
	cid := "13b3ad64-9154-11ea-9eba-47f617ab74f7"

	cfg := yagu.GaConfig{
		TID: "{{ .CLI.Telemetry }}",
		CID: cid,
		UA: "{{ .CLI.cliName }}/" + verinfo.Version,
		CS: "{{ .CLI.cliName }}",
		CN: verinfo.Version,
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

# todo, turn this into a flow on the generator?
# either post or manually run

.PHONY: first
first:
	rm Makefile
	hof mod vendor cue
	hof gen
	go mod tidy
	go build -o {{ .name }} ./cmd/{{ .name }}
	./{{ .name }} -h

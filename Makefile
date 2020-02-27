# export GO111MODULE=on

.PHONY: gen
gen:
	go generate ./...

.PHONY: mod
mod:
	go get -u
	go mod tidy

.PHONY: run
run:
	go run ./main.go

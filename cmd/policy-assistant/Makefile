test:
	go test ./pkg/...

fmt:
	go fmt ./cmd/... ./pkg/...

vet:
	go vet ./cmd/... ./pkg/...

policy-assistant:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./cmd/policy-assistant/policy-assistant ./cmd/policy-assistant
	docker build -t docker.io/mfenwick100/policy-assistant:latest ./cmd/policy-assistant

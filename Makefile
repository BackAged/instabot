lint:
	golint --set_exit_status ./...
	go vet $(go list ./... | grep -v /examples/)


test:
	go test -race -coverprofile=coverage.txt -covermode=atomic
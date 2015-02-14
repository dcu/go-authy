test:
	go test -cover -coverprofile=coverage.out -short
	go tool cover -func=coverage.out
	@rm *.out


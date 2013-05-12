test:
	@@GOPATH=${PWD} go test authy/tests

install_tests:
	@@GOPATH=${PWD} go test -i authy/tests

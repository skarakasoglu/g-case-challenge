build:
	echo "Building GetirCaseChallenge"
	go build -o bin/GetirCaseChallenge main.go

run:
	echo "Running GetirCaseChallenge"
	bin/GetirCaseChallenge

test:
	echo "Running all tests for GetirCaseChallenge"
	go test github.com/skarakasoglu/g-case-challenge
build:
	protoc -I. --go_out=. --micro_out=. \
		proto/configuration/configuration.proto

runlocal:
	DB_HOST=localhost DB_USER=postgres DB_PASSWORD=docker \
		DB_NAME=inact_mini go run *.go
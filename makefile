build:
	protoc -I. --go_out=. --micro_out=. \
		proto/configuration/configuration.proto
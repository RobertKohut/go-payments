BINARY_NAME=payments
SERVICE_NAME=payments

build:
	go build -ldflags "-s -w" -o /usr/local/bin/${BINARY_NAME} cmd/main/main.go

run:
	/usr/local/bin/${BINARY_NAME}

restart:
	systemctl restart ${SERVICE_NAME}

debug:
	go build -gcflags="all=-N -l" -o /usr/local/bin/${BINARY_NAME} cmd/api/main.go
	systemctl stop ${SERVICE_NAME}
	dlv --listen=:2345 --headless=true --api-version=2 exec /usr/local/bin/${BINARY_NAME}

build_and_run: build run
build_and_restart: build restart

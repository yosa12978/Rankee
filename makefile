GOFILE = ./cmd/api/main.go

build:
	go mod tidy
	go build -o ${GOBIN} ${GOFILE}

run:
	go mod tidy
	go run ${GOFILE}

deps:
	go mod tidy
	go get gorm.io/gorm
	go get gorm.io/driver/postgres
	go get github.com/google/uuid

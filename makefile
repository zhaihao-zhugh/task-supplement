export CGO_ENABLED=0
export GOOS=linux
export GOARCH=arm64


NAME=task_supplement_linux

all: build

build:
	go build -ldflags "-s -w" -o ${NAME}  .

clean:
	rm  ${NAME}
# GOOS：darwin、freebsd、linux、windows
# GOARCH：386、amd64、arm
#linux
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux/srun main.go
#windows
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/win/srun.exe main.go
#macos
#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/mac/srun main.go

all: darwin

.PHONY:darwin
darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun-mac cmd/main.go

.PHONY:linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun-linux cmd/main.go

.PHONY:windows
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun-windows cmd/main.go

.PHONY:clean
clean:
	rm -rf ./bin

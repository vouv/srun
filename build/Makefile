# GOOS：darwin、freebsd、linux、windows
# GOARCH：386、amd64、arm

all: darwin

publish:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun-mac cmd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun-linux cmd/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun.exe cmd/main.go

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun cmd/main.go

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun cmd/main.go

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun.exe cmd/main.go

clean:
	rm -rf ./bin

.PHONY:publish
.PHONY:darwin
.PHONY:linux
.PHONY:windows
.PHONY:clean
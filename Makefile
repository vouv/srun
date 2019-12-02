# GOOS：darwin、freebsd、linux、windows
# GOARCH：386、amd64、arm、s390x

all: darwin

publish:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun-mac ./cmd/srun
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun-linux ./cmd/srun
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun.exe ./cmd/srun

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun ./cmd/srun

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun ./cmd/srun

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun.exe ./cmd/srun

clean:
	rm -rf ./bin

.PHONY:publish
.PHONY:darwin
.PHONY:linux
.PHONY:windows
.PHONY:clean
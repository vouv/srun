# GOOS：darwin、freebsd、linux、windows
# GOARCH：386、amd64、arm

all: darwin

publish:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun-mac
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun-linux
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun.exe

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun.exe

clean:
	rm -rf ./bin

.PHONY:publish
.PHONY:darwin
.PHONY:linux
.PHONY:windows
.PHONY:clean
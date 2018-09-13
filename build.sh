#!/bin/bash
# rm
rm -r bin/

# GOOS：darwin、freebsd、linux、windows
# GOARCH：386、amd64、arm
#macos
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/srun-linux64 main.go
#windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/srun-windows.exe main.go
#macos
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/srun-darwin main.go
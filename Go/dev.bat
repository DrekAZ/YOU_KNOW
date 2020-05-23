@echo off
set GOOS=linux
go build storage.go
set GOOS=windows
docker build -t scratch-go .
docker container run --rm scratch-go
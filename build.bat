@echo off
setlocal EnableDelayedExpansion

set archs=amd64 arm64 ppc64le ppc64 s390x

for %%a in (%archs%) do (
    set GOOS=linux
    set GOARCH=%%a
    go build -ldflags "-s -w" -o ./bin/user_manager_%%a
)
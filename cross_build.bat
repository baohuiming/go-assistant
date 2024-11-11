SET CGO_ENABLED=0

SET GOOS=linux
SET GOARCH=amd64
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o dist/go_assistant_linux_amd64

SET GOARCH=386
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o dist/go_assistant_linux_386

SET GOARCH=arm64
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o dist/go_assistant_linux_aarch64

SET GOARCH=arm
SET GOARM=7
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o dist/go_assistant_linux_armv7

SET GOOS=windows
SET GOARCH=amd64
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o dist/go_assistant_windows_x64.exe

SET GOARCH=386
go build -v -a -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -o dist/go_assistant_windows_x86.exe
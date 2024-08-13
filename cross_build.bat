SET GOOS=linux
SET GOARCH=amd64
go build -o dist/go_assistant_linux_amd64

SET GOOS=linux
SET GOARCH=386
go build -o dist/go_assistant_linux_386

SET GOOS=linux
SET GOARCH=arm64
go build -o dist/go_assistant_linux_aarch64

SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -o dist/go_assistant_linux_armv7

SET GOOS=windows
SET GOARCH=amd64
go build -o dist/go_assistant_windows_x64.exe

SET GOOS=windows
SET GOARCH=386
go build -o dist/go_assistant_windows_x86.exe
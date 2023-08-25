## 交叉编译

### windows


go build -o .\bin\client\client-windows.exe .\client\client.go
go build -o .\bin\server\server-windows.exe .\server\server.go

```shell
$OS_ARCH="windows"
$env:CGO_ENABLED=0
$env:GOOS=${OS_ARCH}
$env:GOARCH="amd64"

go build -o .\bin\client\client-${OS_ARCH}.exe .\client\client.go
go build -o .\bin\server\server-${OS_ARCH}.exe .\server\server.go
```

### MacOS

```shell
$OS_ARCH="darwin"
$env:CGO_ENABLED=1
$env:GOOS=${OS_ARCH}
$env:GOARCH="amd64"
go build -o .\bin\client\client-${OS_ARCH} .\client\client.go
go build -o .\bin\server\server-${OS_ARCH} .\server\server.go
```

### Linux

```shell
$OS_ARCH="linux"
$env:CGO_ENABLED=1
$env:GOOS=${OS_ARCH}
$env:GOARCH="amd64"
go build -o .\bin\client\client-${OS_ARCH} .\client\client.go
go build -o .\bin\server\server-${OS_ARCH} .\server\server.go
```
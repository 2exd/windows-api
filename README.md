## 编译

### windows

```shell
$OS_ARCH="windows"
$env:GOOS=${OS_ARCH}
$env:GOARCH="amd64"

go build -ldflags "-s -w" -o .\bin\client\client-${OS_ARCH}.exe .\client\client.go
go build -ldflags "-s -w" -o .\bin\server\server-${OS_ARCH}.exe .\server\server.go
```

## 配置

###  server

```yaml
ip: 10.17.237.33
port: 8081
file: ./temp.txt
```

### client

指定 server `ip:port`

```yaml
addr: 10.17.237.33:8081
```


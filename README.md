# barcode

QR/PDF417二维码生成器

![MacOS](screenshots/macos.png)
![Windows](screenshots/windows.png)

## Windows交叉编译

```shell
brew install mingw-w64
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -ldflags '-w -s -H=windowsgui'
```

#win 64bit
GOOS=windows GOARCH=amd64 go build -o bin/wechatbot-amd64.exe main.go

#win 32-bit
GOOS=windows GOARCH=386 go build -o bin/wechatbot-386.exe main.go

#linux 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/wechatbot-amd64-linux main.go

#linux 32-bit
GOOS=linux GOARCH=386 go build -o bin/wechatbot-386-linux main.go

#mac 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/wechatbot-amd64-darwin main.go

#mac 32-bit
GOOS=darwin GOARCH=386 go build -o bin/wechatbot-386-darwin main.go
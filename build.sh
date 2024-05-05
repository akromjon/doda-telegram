GOOS=darwin     GOARCH=amd64    go build -ldflags '-s' -o bin/doda-telegram-darwin-amd64       main.go
GOOS=darwin     GOARCH=arm64    go build -ldflags '-s' -o bin/doda-telegram-darwin-arm64       main.go
GOOS=linux      GOARCH=386      go build -ldflags '-s' -o bin/doda-telegram-linux-386          main.go
GOOS=linux      GOARCH=amd64    go build -ldflags '-s' -o bin/doda-telegram-linux-amd64        main.go
GOOS=linux      GOARCH=arm      go build -ldflags '-s' -o bin/doda-telegram-linux-arm          main.go
GOOS=linux      GOARCH=arm64    go build -ldflags '-s' -o bin/doda-telegram-linux-arm64        main.go
GOOS=windows    GOARCH=386      go build -ldflags '-s' -o bin/doda-telegram-windows-386.exe    main.go
GOOS=windows    GOARCH=amd64    go build -ldflags '-s' -o bin/doda-telegram-windows-amd64.exe  main.go
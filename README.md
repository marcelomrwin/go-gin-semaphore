
## Create the app and add module
```bash
go mod init article-go-gin-semaphore
go get -u github.com/gin-gonic/gin
```

## Build and run (without left binary)
```bash
go run .
```

## Build the app
```bash
go build -o app
```

## Run the app
```bash
./app
```

## Test the app
```bash
go test ./test -v
```
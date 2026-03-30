# Go Layer7 HTTP Stresser

Dirty and fast Layer7 flooder in Golang.

## Directory Structure
golang-stresser/
├── main.go
├── go.mod
├── README.md
├── proxies.txt          # ← put your proxies here
├── stresser             # Linux binary
└── stresser.exe         # Windows binary
text## Build

```bash
go mod tidy
go build -o stresser main.go
Run
Bash./stresser https://example.com proxies.txt 1000
Proxy Format (proxies.txt)
text1.2.3.4:8080
user:pass@5.6.7.8:3128
Compile for other OS
Bash# Linux
GOOS=linux GOARCH=amd64 go build -o stresser main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o stresser.exe main.go

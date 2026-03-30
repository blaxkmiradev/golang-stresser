# Go Layer7 HTTP Stresser

Dirty fast Layer7 flooder written in Golang. Uses proxy list, random UA, random methods. Good for stress testing.

## Directory Structure
golang-stresser/
├── main.go
├── go.mod
├── README.md
├── proxies.txt      # put your proxies here
├── stresser         # compiled binary (linux)
└── stresser.exe     # compiled binary (windows)
text## How to use

### 1. Build

```bash
go mod tidy
go build -o stresser main.go
2. Run
Bash./stresser <target> <proxylist> <threads>

# Example
./stresser https://example.com proxies.txt 1200
Proxy List Format (proxies.txt)
text1.2.3.4:8080
user:pass@5.6.7.8:3128
45.67.89.10:80
Features

Random GET/POST/HEAD/PUT
Random User-Agents
Proxy rotation (http)
TLS insecure skip
Fast as fuck

Compile for different OS
Bash# Linux
GOOS=linux GOARCH=amd64 go build -o stresser main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o stresser.exe main.go
Tips

Fresh proxies only. Dead ones waste threads.
More threads = more power until your machine or bandwidth chokes.
Run on high bandwidth VPS.

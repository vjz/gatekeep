# GateKeep

**GateKeep** is a lightweight, password-protected CLI tool that blocks and unblocks websites by modifying the hosts file. Ideal for parental controls or focused productivity.

## Build

To build GateKeep for different platforms:

### Windows
```bash
GOOS=windows GOARCH=amd64 go build -o gatekeep.exe main.go
```

### macOS
```bash
GOOS=darwin GOARCH=amd64 go build -o gatekeep main.go
```
## 🔧 Usage

```bash
gatekeep block youtube.com reddit.com
gatekeep unblock tiktok.com
```
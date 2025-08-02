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
sudo gatekeep block youtube.com reddit.com
sudo gatekeep unblock tiktok.com
```

### Important: Flush DNS Cache

After blocking or unblocking sites, you may need to flush DNS caches for changes to take effect:

#### macOS
```bash
# Flush DNS cache
sudo dscacheutil -flushcache
sudo killall -HUP mDNSResponder

# Or use the newer command (macOS 10.12+)
sudo killall -HUP mDNSResponder
```

#### Windows
```bash
# Flush DNS cache
ipconfig /flushdns
```

#### Linux
```bash
# Flush DNS cache (systemd)
sudo systemctl restart systemd-resolved

# Or for older systems
sudo service network-manager restart
```

#### Browser Cache
- **Chrome/Edge**: Press `Ctrl+Shift+Delete` (Windows/Linux) or `Cmd+Shift+Delete` (macOS)
- **Firefox**: Press `Ctrl+Shift+Delete` (Windows/Linux) or `Cmd+Shift+Delete` (macOS)
- **Safari**: Press `Cmd+Option+E` to empty cache

**Note**: Some browsers may still cache DNS results. If a site still loads after blocking, try:
1. Hard refresh: `Ctrl+F5` (Windows/Linux) or `Cmd+Shift+R` (macOS)
2. Open in incognito/private mode
3. Restart the browser completely
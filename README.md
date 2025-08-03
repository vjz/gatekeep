# GateKeep

**GateKeep** is a lightweight CLI tool that blocks and unblocks websites by modifying the hosts file. Ideal for parental controls or focused productivity.

## 🔒 Security Model

GateKeep uses **administrative privileges** as its security mechanism. Since modifying the hosts file requires sudo/admin access, the real security is the admin password itself.

### For Parents
- **Don't share admin passwords** with children
- **Keep the admin account separate** from children's user accounts
- **Use standard user accounts** for children (no sudo access)
- The admin password becomes the "key" to modify website blocks

This approach is more secure than application-level passwords because:
- ✅ **System-level protection** - requires admin privileges
- ✅ **No password to remember** - uses existing admin authentication
- ✅ **Cannot be bypassed** by simply deleting the app or changing settings
- ✅ **Cross-application** - blocks sites system-wide, not just in browsers

## Build

To build GateKeep for different platforms:

### Windows
```bash
mkdir -p bin
GOOS=windows GOARCH=amd64 go build -o bin/gatekeep.exe main.go
```

### macOS
```bash
mkdir -p bin
GOOS=darwin GOARCH=amd64 go build -o bin/gatekeep main.go
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
package blocker

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

var redirect = "127.0.0.1"

func getHostsPath() string {
	switch runtime.GOOS {
	case "windows":
		return "C:\\Windows\\System32\\drivers\\etc\\hosts"
	case "darwin":
		return "/etc/hosts"
	default:
		return "/etc/hosts"
	}
}

func Block(sites []string) {
	path := getHostsPath()
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read hosts file:", err)
		return
	}

	data := string(content)
	changed := false

	for _, site := range sites {
		entry := fmt.Sprintf("%s %s", redirect, site)
		if !strings.Contains(data, entry) {
			data += "\n" + entry
			changed = true
		}
	}

	if changed {
		err = os.WriteFile(path, []byte(data), 0644)
		if err != nil {
			fmt.Println("Failed to write hosts file:", err)
			return
		}
		fmt.Println("Sites blocked:", sites)
	} else {
		fmt.Println("Sites already blocked.")
	}
}

func Unblock(sites []string) {
	path := getHostsPath()
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Failed to read hosts file:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string

	for _, line := range lines {
		shouldSkip := false
		for _, site := range sites {
			if strings.Contains(line, site) && strings.Contains(line, redirect) {
				shouldSkip = true
				break
			}
		}
		if !shouldSkip {
			newLines = append(newLines, line)
		}
	}

	err = os.WriteFile(path, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Println("Failed to write hosts file:", err)
		return
	}

	fmt.Println("Sites unblocked:", sites)
}

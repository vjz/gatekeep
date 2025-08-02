fmt.Println(`
  _____       _       _  __               
 / ____|     | |     (_)/ _|              
| |  __  __ _| |_ ___ _| |_ ___  ___  ___ 
| | |_ |/ _` + "`" + ` | __/ _ \ |  _/ _ \/ __|/ _ \
| |__| | (_| | ||  __/ | || (_) \__ \  __/
 \_____|\__,_|\__\___|_|_| \___/|___/\___|
               GateKeep - Block the Noise
`)
package main

import (
	"fmt"
	"os"
	"strings"

	"myblocker/auth"
	"myblocker/blocker"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  blocker block <site1> <site2> ...")
		fmt.Println("  blocker unblock <site1> <site2> ...")
		return
	}

	action := os.Args[1]
	sites := os.Args[2:]

	fmt.Print("Enter password: ")
	var pass string
	fmt.Scanln(&pass)

	if !auth.CheckPassword(pass) {
		fmt.Println("Incorrect password.")
		return
	}

	switch strings.ToLower(action) {
	case "block":
		blocker.Block(sites)
	case "unblock":
		blocker.Unblock(sites)
	default:
		fmt.Println("Invalid command. Use 'block' or 'unblock'.")
	}
}

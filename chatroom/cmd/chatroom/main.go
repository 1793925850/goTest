package main

import "fmt"

var (
	addr   = ":2022"
	banner = `
    ____              _____
   |    |    |   /\     |
   |    |____|  /  \    | 
   |    |    | /----\   |
   |____|    |/      \  |

	ChatRoom, start on: %s
`
)

func main() {
	fmt.Printf(banner, addr)

}

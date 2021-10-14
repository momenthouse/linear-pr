package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello %s", os.Args[1])
	fmt.Printf("GITHUB_REPOSITORY %s", os.Getenv("GITHUB_REPOSITORY"))
	fmt.Printf("GITHUB_EVENT_PATH  %s", os.Getenv("GITHUB_EVENT_PATH"))
	fmt.Printf("GITHUB_REPOSITORY_OWNER %s", os.Getenv("GITHUB_REPOSITORY_OWNER"))
}

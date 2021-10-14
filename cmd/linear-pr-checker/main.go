package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Printf("Hello %s", os.Args[1])
	fmt.Printf("GITHUB_REPOSITORY %s", os.Getenv("GITHUB_REPOSITORY"))
	fmt.Printf("GITHUB_EVENT_PATH  %s", os.Getenv("GITHUB_EVENT_PATH"))
	fmt.Printf("GITHUB_REPOSITORY_OWNER %s", os.Getenv("GITHUB_REPOSITORY_OWNER"))

	// Open our jsonFile
	jsonFile, err := os.Open(os.Getenv("GITHUB_EVENT_PATH"))
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["users"])
}

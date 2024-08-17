package main

import "log"

func main() {
	err := LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

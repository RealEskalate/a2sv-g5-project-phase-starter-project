package main

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"fmt"
	"log"
)

func main() {
	content := `This article is about how numbers are written as words in various languages.  It also has a tiny bit about functions, specifically fixed points of functions, which I use to describe how numbers work in the different languages. I speak only one of the languages below as a native, so I’ve probably made lots of mistakes – please leave a comment if you find one.`
	resp, err := infrastructure.RecommendTitle(content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

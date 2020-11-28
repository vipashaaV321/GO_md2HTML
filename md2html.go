package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/gomarkdown/markdown"
)

func main() {

	file := "readme.md"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("%s file not found", file)
	}
	
	html :=markdown.ToHTML(content,nil,nil)
	
	fileOut:= "result.html"
	writeErr:=ioutil.WriteFile(fileOut,html,0644)

	if writeErr != nil{
		log.Fatalf("can't write in %s file", fileOut)
	}

	fmt.Println(string(html))
	fmt.Println("MD to HTML converted!")
}
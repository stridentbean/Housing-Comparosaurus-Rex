package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//res, err := http.Get("http://www.google.com/robots.txt")
	res, err := http.Get("https://sfbay.craigslist.org/apa/")
	
	if err != nil {
		log.Fatal(err)
	}
	rawText, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%s", rawText)
	
	// write whole the body
    err = ioutil.WriteFile("output.txt", rawText, 0644)
    if err != nil { panic(err) }
		
}
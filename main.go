package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "os"
    "strings"

)


func main() {
    // Create HTTP GET request
    response, err := http.Get("https://www.google.com/")
    if err != nil {
        log.Fatal(err)
    }
    defer response.Body.Close()

    // Get the response body as a string
    bodyInBytes, err := ioutil.ReadAll(response.Body)
    pageContent := string(bodyInBytes)

    // substring 
    titleStartingIndex := strings.Index(pageContent, "<title>")
    if titleStartingIndex == -1 {
        fmt.Println("No title found")
        os.Exit(0)
    }
 
    titleStartingIndex += 7

    titleEndingIndex := strings.Index(pageContent, "</title>")
    if titleEndingIndex == -1 {
        fmt.Println("No tag found for title")
        os.Exit(0)
    }

    pageTitle := []byte(pageContent[titleStartingIndex:titleEndingIndex])

    fmt.Printf("Page title: %s\n", pageTitle)
}

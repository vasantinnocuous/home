package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "log"
    "os"
	"strings"
	
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func goLogin() {
    // create a new collector
    c := colly.NewCollector()

    // authenticate
    err := c.Post("http://github.com/login", map[string]string{"username": "admin123", "password": "admin456"})
    if err != nil {
        log.Fatal(err)
    }

    // attach callbacks after login
    c.OnResponse(func(r *colly.Response) {
        log.Println("response received", r.StatusCode)
    })

    // start scraping
    c.Visit("https://github.com/")
}

func goGet() {
	var headings, row []string
	var rows [][]string

	data := `<html><body>
	<table>
		<tr><th>Marvel Heading One</th><th>Marvel Heading Two</th></tr>
		<tr><td>Marvel 1</td><td>Marvel 2</td></tr>
		<tr><td>Marvel 3</td><td>Marvel 4</td></tr>
		<tr><td>Marvel 5</td><td>Marvel 6</td></tr>
		<tr><td>marvel 7</td><td>Marvel 8</td></tr>
	</table>
	<p>Lets begin</p>
	<table>
		<tr><th>DC Heading One</th><th>DC Heading two</th></tr>
		<tr><td>DC 1</td><td>DC 2</td></tr>
		<tr><td>DC 3</td><td>DC 4</td></tr>
		<tr><td>DC 5</td><td><span></span><span><a href="">DC 6</a></span></td></tr>
		<tr><td>DC 7</td><td>DC 8</td></tr>
	</table>
	</body>
	</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))
	if err != nil {
		fmt.Println("No url found")
		log.Fatal(err)
	}

	// Find each table
	doc.Find("table").Each(func(index int, tablehtml *goquery.Selection) {
		tablehtml.Find("tr").Each(func(indextr int, rowhtml *goquery.Selection) {
			rowhtml.Find("th").Each(func(indexth int, tableheading *goquery.Selection) {
				headings = append(headings, tableheading.Text())
			})
			rowhtml.Find("td").Each(func(indexth int, tablecell *goquery.Selection) {
				row = append(row, tablecell.Text())
			})
			rows = append(rows, row)
			row = nil
		})
	})
	fmt.Println("Headings = ", len(headings), headings)
	fmt.Println("Rows = ", len(rows), rows)
}


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

	
	goGet()
	goLogin()
}


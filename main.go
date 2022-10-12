package main

import (
	"fmt"
	"os"

	link "github.com/mouad-eh/html-link-parser/link"
)

// var HTMLText = `
// <html>

// <body>
//     <h1>Hello!</h1>
//     <a href="/other-page">A link to another page</a>
// </body>

// </html>
// `

func main() {
	// a way create a reader out of a string
	// r := strings.NewReader(HTMLText)
	r, err := os.Open("ex4.html")
	if err != nil {
		panic(err)
	}
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}

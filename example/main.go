package main

import (
	"fmt"
	"metabolize"
	"net/http"
	"net/url"
)

type MetaData struct {
	Title       string  `meta:"og:title,title"`
	Description string  `meta:"og:description,description"`
	Type        string  `meta:"og:type"`
	URL         url.URL `meta:"og:url"`
	VideoWidth  int64   `meta:"og:video:width"`
	VideoHeight int64   `meta:"og:video:height"`
}

func main() {
	res, _ := http.Get("https://zozo.vn")

	data := new(MetaData)

	err := metabolize.Metabolize(res.Body, data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Title: %s\n", data.Title)
	fmt.Printf("Description: %s\n", data.Description)
	fmt.Printf("Type: %s\n", data.Type)
	fmt.Printf("URL: %s\n", data.URL.String())
	fmt.Printf("VideoWidth: %d\n", data.VideoWidth)
	fmt.Printf("VideoHeight: %d\n", data.VideoHeight)
}

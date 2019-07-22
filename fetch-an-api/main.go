package main

import (
	"fmt"

	"github.com/thenguyenit/testing/fetch-an-api/fetch"
)

func main() {
	//Fetch Astros
	astroAPI := &fetch.AstroAPI{
		URL: "http://api.open-notify.org/astros.json",
	}
	number, err := astroAPI.Get()
	fmt.Println(number)
	fmt.Println(err)
}

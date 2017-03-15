package main

import (
	"fmt"
	"github.com/cseeger-epages/i-doit-go-api"
)

func main() {
	// create api object
	a, _ := goidoit.Newapi("https://example.com/src/jsonrpc.php", "yourapikey")

	// search for sht like test
	s, _ := a.Search("test")

	/* show output
	the SearchResponse type is nearly the same like Response type
	except that it does handle the type assertions for you
	*/
	for i, v := range s.Result {
		fmt.Println("----", i, "----")
		for k, d := range v {
			fmt.Println(k, d)
		}
	}

	// or a specific value
	fmt.Println("----", "Single Output", "----")
	fmt.Println(s.Result[0]["value"])
}

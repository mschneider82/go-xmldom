package main

import (
	"fmt"
	"github.com/subchen/go-xmldom"
)

func main() {
	dom, err := xmldom.ParseFile("test/ovfenv-1.xml")
	if err != nil {
		panic(err)
	}

	fmt.Println(dom.XMLPretty())
}

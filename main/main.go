package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"

	"github.com/antchfx/xmlquery"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"

	// - "github.com/gogs/gogs"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/owncast/owncast/logging"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
	fmt.Println("Hello world")
	wadl, err := xmlquery.LoadURL("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	attr := xmlquery.FindOne(wadl, "//application/@xmlns")
	fmt.Println(attr.InnerText())
}

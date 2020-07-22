package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	var b bytes.Buffer

	if _, err := io.Copy(&b, r); err != nil {
		log.Fatal(err)
	}

	fmt.Println(b.String())
}

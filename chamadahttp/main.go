package main

import (
	"io"
	"net/http"
)

// chamadas http
func main() {

	req, err := http.Get("http://criptoclube.com")
	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	println(string(res))
	req.Body.Close()

}

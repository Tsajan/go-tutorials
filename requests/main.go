package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://storage.googleapis.com/cdh-common-config/tokens.json")
	if err != nil {
		log.Fatal(err)
	}

	// we read the response body on the line below
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// convert the body to type string
	sb := string(body)
	log.Printf(sb)

	postBody, _ := json.Marshal(map[string]string{
		"name":  "Sajan Maharjan",
		"email": "sajan.maharjan294@gmail.com",
	})

	resp, err = http.Post("https://postman-echo.com/post", "application/json", bytes.NewBuffer(postBody))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb2 := string(body)
	log.Printf(sb2)
}

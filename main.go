package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	showHeader = false
	showBody   = true
	https      = false
	fileName   string
)

func init() {
	flag.BoolVar(&showHeader, "head", false, "show headers")
	flag.BoolVar(&showBody, "body", true, "show body")
	flag.BoolVar(&https, "https", false, "if default use https")

	flag.Parse()
	fileName = flag.Arg(0)
}

func main() {

	strByte, err := getStr(fileName)
	if err != nil {
		log.Fatal(err)
	}
	str := string(strByte)
	if !strings.Contains(str, "\n\n") {
		str = str + "\n\n"
	}
	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(str)))
	if err != nil {
		log.Fatal(err)
	}

	if req.URL.Scheme == "" {
		if https {
			req.URL.Scheme = "https"
		} else {
			req.URL.Scheme = "http"
		}
	}
	req.URL.Path = req.RequestURI
	req.URL.Host = req.Host

	req.RequestURI = ""
	client := &http.Client{}
	res, err := client.Do(req)
	if err!= nil     {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Errorf("status code %d\n", res.StatusCode)
	}

	if showHeader {
		for it, value := range res.Header {
			fmt.Println(it + ": " + strings.Join(value, " "))
		}

		if showBody {
			fmt.Println()
		}
	}

	if showBody {
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}

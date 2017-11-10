package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/k4s/phantomgo"
)

func main() {
	p := phantomgo.NewPhantom()
	jsBytes, err := ioutil.ReadFile("parse.js")
	if err != nil {
		fmt.Println(err)
	}
	js := string(jsBytes)
	js = strings.Replace(js, "<<URL>>", "http://www.intermod.ro", -1)

	res, _ := p.Exec(js)
	output, _ := ioutil.ReadAll(res)
	fmt.Println(string(output))
}

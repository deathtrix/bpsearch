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
	js = strings.Replace(js, "<<SIZE_WEIGHT>>", "12", -1)
	js = strings.Replace(js, "<<BOLD_WEIGHT>>", "4/3", -1)
	js = strings.Replace(js, "<<H1_WEIGHT>>", "2", -1)
	js = strings.Replace(js, "<<H2_WEIGHT>>", "5/3", -1)
	js = strings.Replace(js, "<<H3_WEIGHT>>", "4/3", -1)
	js = strings.Replace(js, "<<H4_WEIGHT>>", "4/3", -1)
	js = strings.Replace(js, "<<NRP_WEIGHT>>", "2", -1)

	res, _ := p.Exec(js)
	output, _ := ioutil.ReadAll(res)
	fmt.Println(string(output))
}

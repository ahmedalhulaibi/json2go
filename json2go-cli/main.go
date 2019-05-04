package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gitlab.com/ahmedalhulaibi/json2go"
)

const usage = `json2go-cli - Outputs go struct based on given JSON input to stdin

Usage: json2go-cli [options]
	--typename 		The typename for the output struct

Example:
	curl http://api.open-notify.org/astros.json | json2go-cli --typename Astros`

func main() {

	help := flag.Bool("help", false, "help")
	typename := flag.String("typename", "", "typename for the output struct")
	flag.Parse()
	if *help {
		fmt.Println(usage)
		os.Exit(0)
	}

	if json, err := ioutil.ReadAll(os.Stdin); err == nil {
		if gostruct, err := json2go.JsonToGo(json, *typename); err == nil {
			fmt.Fprintf(os.Stdout, "%s", string(gostruct))
		} else {
			log.Println("json2go error: ", err)
			os.Exit(1)
		}
	} else {
		log.Println("json2go error: ", err)
		os.Exit(1)
	}

	os.Exit(0)
}

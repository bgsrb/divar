package main

import (
	"encoding/json"
	"log"
	"time"

	"flag"

	"github.com/bgsrb/divar"
)

var oldPostList = make(map[string]divar.Post)
var newPostList = make(map[string]divar.Post)

const usage = `
Usage:
	divar -req= -interval=
Examples:
	divar -req='{"jsonrpc":"2.0","id":1,"method":"getPostList","params":[[["place2",0,["1"]],["cat3",0,[26]],["cat2",0,[24]],["cat1",0,[1]],["v10",0,["2"]]],0]}' -interval=10
`

var reqStr = flag.String("req", `{"jsonrpc":"2.0","id":1,"method":"getPostList","params":[[["place2",0,["1"]],["cat3",0,[26]],["cat2",0,[24]],["cat1",0,[1]],["v10",0,["2"]]],0]}`, "params")
var interval = flag.Uint64("interval", 60, "interval")

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Usage = func() {
		panic(usage)
	}
	flag.Parse()
}

func main() {

	client := divar.NewClient()
	var req = divar.Request{}
	err := json.Unmarshal([]byte(*reqStr), &req)
	if err != nil {
		panic(err)
	} else {
		postList, err := client.GetPostList(req)
		if err != nil {
			panic(err)
		}
		for _, p := range postList {
			oldPostList[p.Token] = p
		}
		var ticker = time.NewTicker(time.Duration(*interval) * time.Second)
		quit := make(chan struct{})
		for {
			select {
			case <-ticker.C:
				postList, err := client.GetPostList(req)
				if err != nil {
					panic(err)
				} else {
					for _, p := range postList {
						if _, ok := oldPostList[p.Token]; !ok {
							newPostList[p.Token] = p
							go notify(p)
						}
						oldPostList[p.Token] = p
					}
				}
				break
			case <-quit:
				ticker.Stop()
			}
		}
	}
}

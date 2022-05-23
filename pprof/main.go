package main

import (
	"github.com/yang201396/GoExamples/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/yang201396"))
		}
	}()

	err := http.ListenAndServe("0.0.0.0:6060", nil)
	if err != nil {
		return
	}
}

package main

import (
	"log"
	"sync"

	"github.com/756445638/go-example/example"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	//go example.RunClient(&wg)

	err := example.RunServer(&wg, ":8080", "root:123@tcp(localhost:3306)/test?charset=utf8")
	log.Fatalf("server is down,err:%v", err)

	wg.Wait()
}

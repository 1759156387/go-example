package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"unsafe"

	"github.com/756445638/go-example/controller"
)

var (
	wg = sync.WaitGroup{}
)

var (
	conf config
)

type config struct {
	Listen string
	Mysql  string
}

func readConf() error {
	filename := filepath.Abs(os.Args[0]) + "../conf/conf.json"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &conf)
}

func main() {

	err := readConf()
	if err != nil {
		panic(error)
	}

	err := example.RunServer(&wg, conf.Listen, conf.Mysql)
	log.Fatalf("server is down,err:%v", err)

	wg.Wait()
}

func dll() {
	var mod = syscall.NewLazyDLL("user32.dll")
	var proc = mod.NewProc("MessageBoxW")
	var MB_YESNOCANCEL = 0x00000003

	ret, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("calling ddl"))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Done Title"))),
		uintptr(MB_YESNOCANCEL))
	fmt.Printf("Return: %d\n", ret)
}

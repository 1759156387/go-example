package main

import (
	"fmt"
	"log"
	"sync"
	"syscall"
	"unsafe"

	"github.com/756445638/go-example/example"
)

var (
	wg = sync.WaitGroup{}
)

func main() {
	//go example.RunClient(&wg)
	go dll()
	err := example.RunServer(&wg, ":8080", "root:123@tcp(localhost:3306)/test?charset=utf8")
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

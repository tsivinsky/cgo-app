package main

/*
#include "hello.h"
*/
import "C"

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Printf("What's your name: ")
	name, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name = strings.TrimSpace(name)

	cs := C.CString(name)
	C.SayHello(cs)
	C.free(unsafe.Pointer(cs))
}

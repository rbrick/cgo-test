package main

//#include "hello.h"
import "C"
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

//export SayHello
func SayHello(x *C.char) {
	fmt.Println(C.GoString(x))
}

//export Test
func Test() (int, int) {
	return 1337, 1337
}

//export Fib
func Fib(i int) int {
	if i == 1 || i == 0 {
		return 1
	}
	return Fib(i-1) + Fib(i-2)
}

//export MakeHTTPRequestC
func MakeHTTPRequestC(url string) *C.char {
	return C.CString(MakeHTTPRequest(url))
}

// Makes an HTTP request, and turns the response into a string.
func MakeHTTPRequest(url string) string {
	fmt.Println("Making a HTTP request for:", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}
	d, _ := ioutil.ReadAll(resp.Body)
	return string(d)
}

func main() {
	C.printHello()
	time.Sleep(10 * time.Second)
}

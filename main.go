package main

//#include "hello.h"
import "C"
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"syscall"
	"time"
	"unsafe"
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

func abort(funcname string, err error) {
	panic(fmt.Sprintf("%s failed: %v", funcname, err))
}

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	getModuleHandle, _ = syscall.GetProcAddress(kernel32, "GetModuleHandleW")

	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300

	MB_CANCEL_ANSWER    = 0x02
	MB_TRY_AGAIN_ANSWER = 0x0A
	MB_CONTINUE_ANSWER  = 0x0B
)

func MessageBox(caption, text string, style uintptr) (result int) {
	var nargs uintptr = 4
	ret, _, callErr := syscall.Syscall9(uintptr(messageBox),
		nargs,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		style,
		0,
		0,
		0,
		0,
		0)
	if callErr != 0 {
		abort("Call MessageBox", callErr)
	}
	result = int(ret)
	return
}

func GetModuleHandle() (handle uintptr) {
	var nargs uintptr = 0
	if ret, _, callErr := syscall.Syscall(uintptr(getModuleHandle), nargs, 0, 0, 0); callErr != 0 {
		abort("Call GetModuleHandle", callErr)
	} else {
		handle = ret
	}
	return
}

func main() {
	C.printHello()
	time.Sleep(10 * time.Second)

	stringChan := make(chan string)
	printChan := make(chan string)

	for x := 0; x < 10; x++ {
		go func(id int) {
			for i := 0; i < 10000000; i++ {
				x := i * i
				x = (i * i)
				printChan <- fmt.Sprint("%d", x)
			}
			stringChan <- fmt.Sprintf("Thread #%d completed", id+1)
		}(x)
	}

	for m := range printChan {
		fmt.Println(m)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-stringChan)
	}

}

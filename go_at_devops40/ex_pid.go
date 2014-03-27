package main

/*
#include <windows.h>
*/
import "C"

import "fmt"

func GetCurrentProcessID() int {
	return int(C.GetCurrentProcessId())
}

func main() {
	fmt.Println("My PID is", GetCurrentProcessID())
} 
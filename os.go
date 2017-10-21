//Fonte para descobrir os dados do os
package main

import (
	"fmt"
)

type StOperatingSystem struct {
	KernelName	string //Ex: Linux
	KernelVersion string 
	Architecture  string
	Name string
	Version string
}

func discovery() *StOperatingSystem {
	//discovery os information
	osInfo := StOperatingSystem{}
//	cat /etc/*-release
	fmt.Println(commandLine("cat", "/etc/*-release"))

//	lsb_release -a
	
	// discover kernel version

//	$ uname -mrs
//	cat /proc/version
	return &osInfo
}


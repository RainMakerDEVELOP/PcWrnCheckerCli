package main

import (
	cliproc "PcWrnChecker/PcWrnCheckerCli/pwcpkg"
	"fmt"
)

func main() {
	fmt.Println("-------------------------------------------------")
	fmt.Println("---------- Pc Wrn Checker Client Start ----------")
	fmt.Println("-------------------------------------------------")
	fmt.Println("")

	// 실제 모든 데이터 처리를 PcWrnCheckerCliProc 으로 명명된 cliproc 패키지에서 하도록 한다.
	if !cliproc.Run() {
		fmt.Println("FAIL : argument parse error")
	}

	fmt.Println("")
	fmt.Println("-------------------------------------------------")
	fmt.Println("---------- Pc Wrn Checker Client End   ----------")
	fmt.Println("-------------------------------------------------")
}

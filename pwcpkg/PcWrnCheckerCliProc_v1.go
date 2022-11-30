// 해당 패키지에 큰 변동이 있을 경우,
// 1. 신규 go 파일을 버전명으로 하나 더 생성
// 2. 기존 go 파일의 확장자에 .bak 를 추가하여 미사용 처리

package cliproc

import (
	"fmt"
	"os"
	"strings"
)

func Run() bool {
	fmt.Println("CliProc Run Function Start!")

	// 실행인자 체크 루틴 실행
	if !readArg() {
		return
	}
}

func readArg() (bRet bool) {
	bRet = false
	var argCnt int
	var arg string
	for _, v := range os.Args {
		if strings.Index(v, "-type") == 0 {
			arg = strings.Replace(v, "-type=", "", -1)
			argCnt++
		}
	}

	fmt.Println(arg, argCnt)

	if argCnt >= 2 {
		bRet = true
	}

	return
}

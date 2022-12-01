// 해당 패키지에 큰 변동이 있을 경우,
// 1. 신규 go 파일을 버전명으로 하나 더 생성
// 2. 기존 go 파일의 확장자에 .bak 를 추가하여 미사용 처리

package cliproc

import (
	"fmt"
	"os"
	"strings"
)

const (
	USEDCPU = "USEDCPU"
)

func Run() bool {
	fmt.Println("CliProc Run Function Start!")

	// 실행인자 체크 루틴 실행
	var strType string
	var bRet bool

	if bRet, strType = readArg(); !bRet {
		return false
	}

	fmt.Printf("readArg() ret strType = '%v'\n", strType)
	// 서버 전송용 JSON 설정

	// 서버 전송

	// 첫번째 전송 성공하면 goroutine 실행하여 주기적으로 전송

	// 주기적 전송 루틴이 goroutine 으로 실행되므로, goroutine 종료전에 리턴되면 안되므로 체크하여 기다리는 루틴 삽입

	// 모든 하위 프로세스가 종료되면 리턴한다.

	fmt.Println("CliProc Run Function End!")
	return true
}

func readArg() (bool, string) {
	var arg []string
	for _, v := range os.Args {
		// 상태 체크할 사항 실행인자값이 반드시 있어야 함.
		if strings.Index(v, "-type=") == 0 {
			arg = append(arg, strings.Replace(v, "-type=", "", -1))
		}
	}

	// 필수 실행인자가 없었으면 실패
	if len(arg) <= 0 {
		fmt.Printf("not enough argument (cnt = '%v')\n", len(arg))
		return false, ""
	}
	fmt.Printf("argument cnt = '%v'\n", len(arg))

	// 필수 실행인자의 값이 미리 정해진 상태 체크 리스트에 존재하는지 체크
	if strings.Compare(USEDCPU, arg[0]) != 0 {
		fmt.Printf("'%v' not in type list\n", arg[0])
		return false, ""
	}
	fmt.Printf("'%v' in type list\n", arg[0])

	return true, arg[0]
}

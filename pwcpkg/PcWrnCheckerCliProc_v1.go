// 해당 패키지에 큰 변동이 있을 경우,
// 1. 신규 go 파일을 버전명으로 하나 더 생성
// 2. 기존 go 파일의 확장자에 .bak 를 추가하여 미사용 처리

package cliproc

import "fmt"

func Run() {
	fmt.Println("This is CliProc Run Function!")
}

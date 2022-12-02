// 해당 패키지에 큰 변동이 있을 경우,
// 1. 신규 go 파일을 버전명으로 하나 더 생성
// 2. 기존 go 파일의 확장자에 .bak 를 추가하여 미사용 처리

package cliproc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// 모니터링 수집 대상 명칭 목록
const (
	USEDCPU = "USEDCPU" // CPU 사용량
	USEDMEM = "USEDMEM" // Memory 사용량
)

type RestData struct {
	ItemName string `json:"itemname"`
	Value    string `json:"value"`
}

func Run() bool {
	fmt.Println("CliProc Run Function Start!")

	// [PROC] 실행인자 체크 루틴 실행
	var strType string
	var bRet bool

	if bRet, strType = readArg(); !bRet {
		return false
	}
	fmt.Printf("readArg() ret strType = '%v'\n", strType)

	// [PROC] 서버 전송
	resp, bReqRet := reqClient(strType, "application/json")
	if !bReqRet {
		fmt.Println("http request FAILED")
		return false
	}

	defer resp.Body.Close()

	var respData RestData

	// [PROC] 첫번째 전송 성공인지 응답 체크
	if resp.StatusCode == http.StatusOK {
		// 응답 데이터 헤더 Content-Type 조사
		if strings.Compare(resp.Header.Get("Content-Type"), "Content-Type") == 0 {
			// [PROC] goroutine 실행하여 주기적으로 전송

			// 응답 데이터 파싱 테스트용 코드 START
			body, errIoRead := io.ReadAll(resp.Body)
			if errIoRead != nil {
				log.Panicln(errIoRead.Error())
			} else {
				errUnmarshal := json.Unmarshal(body, &respData)

				if errUnmarshal != nil {
					log.Panicln(errUnmarshal.Error())
				} else {
					fmt.Printf("Response Data Parse : '%v'\n", respData)
				}
			}
			// 응답 데이터 파싱 테스트용 코드 END
		}
	}

	// [PROC] 주기적 전송 루틴이 goroutine 으로 실행되므로, goroutine 종료전에 리턴되면 안되므로 체크하여 기다리는 루틴 삽입

	// [PROC] 모든 하위 프로세스가 종료되면 리턴한다.

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

	// 필수 실행인자가 없으면 실패
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

func makeJsonData(data string) ([]byte, error) {
	stData := RestData{data, "START"}
	return json.Marshal(stData)
}

func reqClient(strSendData string, strContentType string) (*http.Response, bool) {
	// 서버 전송용 JSON 설정
	btSendData, marshalErr := makeJsonData(strSendData)
	if marshalErr != nil {
		log.Panicln(marshalErr.Error())
		return nil, false
	}

	if len(btSendData) <= 0 {
		fmt.Printf("marshaling data return len = '%v'\n", len(btSendData))
		return nil, false
	}
	fmt.Println(btSendData)
	fmt.Println(string(btSendData))

	// HTTP 로 요청 전송
	// -- NOTE : 2022.12.02 http.Client{} > http.NewRequest > http.Client{}.do 로 하면 응답 헤더가 제대로 넘어오지 않아서 http.Post 로 변경함
	parsedUrl, _ := url.Parse("http://localhost:1234/" + strSendData)
	res, errRequest := http.Post(parsedUrl.String(), strContentType, bytes.NewBuffer(btSendData))
	if errRequest != nil {
		log.Panicln(errRequest.Error())
		return nil, false
	}
	fmt.Printf("http response status = '%v'\n", res.Status)

	return res, true
}

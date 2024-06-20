// script/execution.go
package script

/**
1. author : 유용원
2. 주요 기능 (순서대로)
2-1). 입력 받은 스크립트 실행전, 어떤 OS에서 동작해야 할지 판단
	2-1-1). Windows 인지 linux 인지에 따라 각각 powershell or /bin/sh 로 나뉘어 입력받은 스크립트를 실행
2-2). 실행 결과는 총 3가지로 구분함 (정상, 오류, timeout(3분))
	2-2-1). 정상일시
		SaveExecutionLog 함수 호출하여 실행 결과 저장 (단 result는 최대 255자 이므로 임의로 254바이트 까지 끊어서 저장시킴)
	2-2-2). 비정상일시 (오류 + timeout)
		SaveExecutionLog 함수 호출하여 실행 결과 저장 (단 result는 최대 255자 이므로 임의로 254바이트 까지 끊어서 저장시킴)
		http://localhost:8080/mailsend/{scriptId}?q={comment} 로 메일서버를 GET 방식으로 http 통해 호출
**/
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"datacatAgent/ent"
)

func SaveExecutionLog(client *ent.Client, scriptID int, status int, result string) error {
	ctx := context.Background()
	_, err := client.ExecutionLog.Create().
		SetScriptId(scriptID).
		SetStatus(status).
		SetResult(result).
		SetExecutedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed to save execution log: %v", err)
	}
	log.Printf("Execution log saved for script ID %d", scriptID)
	return nil
}

// 주어진 바이트 슬라이스를 최대 maxBytes 바이트로 자릅니다.
func trimToMaxBytes(b []byte, maxBytes int) []byte {
	if len(b) <= maxBytes {
		return b
	}

	// maxBytes 위치에서 잘라내고, 해당 위치에서부터 역방향으로 유효한 UTF-8 문자를 찾습니다.
	cut := b[:maxBytes]
	for len(cut) > 0 && !utf8.Valid(cut) {
		cut = cut[:len(cut)-1]
	}
	return cut
}

func ExecuteScript(client *ent.Client, wg *sync.WaitGroup, scriptId int, command string, comment string) {
	defer wg.Done() // 작업이 완료되면 WaitGroup의 카운터를 감소 (defer 문은 프로그램 종료시 무조건 실행되는 구문임)

	// 쉘 명령어를 실행하기 위해 exec.CommandContext를 사용하여 컨텍스트 전달
	var cmd *exec.Cmd
	// 타임아웃 컨텍스트 생성 (180초)
	ctx, cancel := context.WithTimeout(context.Background(), 180*time.Second)
	defer cancel()

	switch runtime.GOOS {
	case "windows":
		cmd = exec.CommandContext(ctx, "powershell.exe", "-c", command)
	case "linux":
		cmd = exec.CommandContext(ctx, "/bin/sh", "-c", command)
	default:
		cmd = exec.CommandContext(ctx, "/bin/sh", "-c", command)
	}

	output, err := cmd.CombinedOutput()
	status := 0
	if ctx.Err() == context.DeadlineExceeded {
		status = 1
		output = []byte("script execution timed out after 180 seconds")
	} else if err != nil {
		status = 2
	}

	// 실행 결과를 executionlog 테이블에 저장

	limitedOutput := trimToMaxBytes(output, 254)
	if saveErr := SaveExecutionLog(client, scriptId, status, string(limitedOutput)); saveErr != nil {
		log.Printf("Error saving execution log for script ID %d: %v", scriptId, saveErr)

	}

	if status != 0 {
		log.Printf("Failed to execute script ID %d: %s", scriptId, string(output))
		log.Printf("Failed to execute script ID(Error log) %d: %s", scriptId, err)
		comment := strings.Replace(comment, " ", "_", -1)
		baseURL := "http://localhost:8080/mailsend/" + strconv.Itoa(scriptId)

		// URL 인코딩
		params := url.Values{}
		params.Add("q", comment)
		urlWithParams := fmt.Sprintf("%s?%s", baseURL, params.Encode())

		// GET 요청 보내기
		resp, err := http.Get(urlWithParams)
		if err != nil {
			fmt.Printf("Error making GET request: %v\n", err)
			return
		}
		defer resp.Body.Close()

		return
	}

	log.Printf("Script ID %d output: %s", scriptId, string(output))
}

//http://localhost:8080/mailsend/1?q=test

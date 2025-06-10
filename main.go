// main.go
package main

/**
1. author : 유용원
2. 주요 기능 (순서대로)
2-1) DB커넥션 후 script 테이블에서 usable 1(사용중)인 script 전체 조회
2-2) 최종 실행 시간 체크를 위해 script ID를key로 ExecutionLog 테이블에서 실행이력 조회
	2-2-1). 이력 존재시 최종 실행 시간(lastExecuteAt) 변수를 ExecutionLog 테이블의 최종 실행 시간으로 저장
		2-2-1-1). 최종 실행 시간(lastExecuteAt) 변수의 시간 부분과 시작시간~종료시간 을 비교 하여 실행 여부 판단
	2-2-2). 이력 미 존재시 최종 실행 시간(lastExecuteAt) 변수를 1900.01.01 로 저장
2-3) 실행 주기를 판단하여 실행 여부 최종 확인
2-4) scriptExecutor 호출하여 스크립트 실행
**/

import (
	"context"
	"datacatAgent/ent"
	"datacatAgent/ent/executionlog"
	"datacatAgent/ent/script"
	scriptpkg "datacatAgent/scriptExecutor"
	db "datacatAgent/util"
	"log"
	"sync"
	"time"

	_ "datacatAgent/util"
)

/*
*
1. Author : 유용원
2. 주요기능
2-1). 웹포털에서 입력한 시작시간 ~ 종료시간에 현재 시간이 포함되는지 확인
*
*/
func IsTimeWithinNMinutes(t time.Time, minutes time.Duration) bool {
	now := time.Now()                               // 현재 시간을 가져옵니다.
	adjustedTime := now.Add(-minutes * time.Minute) // 현재 시간에서 minutes 분을 뺀 시간을 계산합니다.
	return t.After(adjustedTime) && t.Before(now)   // t가 adjustedTime 이후이면서 now 이전인지 확인합니다.
}

/*
*
1. Author : 유용원
2. 주요기능
2-1). 실행 주기 체크하기 위함, 최종 실행 시간을 start에 입력받고 실행주기가 더해진 시간을 end에 입력받으면
현재시간이 실행시간 보다 이후이고, 실행주기가 더해진 시간이 현재보다 이전일 경우를 판단하여 리턴
*
*/
func IsTimeWithinStartNEndTime(start string, end string) bool {
	now := time.Now() // 현재 시간을 가져옵니다.
	layout := "20060102"
	formattedDate := now.Format(layout)
	layout = "2006010215:04:05"
	comparableStartDate := formattedDate + start
	comparableEndDate := formattedDate + end
	location, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		log.Fatalf("Failed to load location: %v", err)
	}
	startTime, startDateErr := time.ParseInLocation(layout, comparableStartDate, location)
	if startDateErr != nil {
		log.Fatalf("Failed to Parse startDate: %v", startDateErr)
	}
	endTime, endDateErr := time.ParseInLocation(layout, comparableEndDate, location)
	if endDateErr != nil {
		log.Fatalf("Failed to Parse endDateErr: %v", endDateErr)
	}

	return now.After(startTime) && now.Before(endTime) //
}

func main() {
	client, err := db.OpenDB()
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	for true { //5분마다 실행
		// usable이 1인 스크립트만 쿼리
		scripts, err := client.Script.Query().
			Where(script.Usable(1)).
			All(context.Background())
		if err != nil {
			log.Fatalf("failed getting scripts: %v", err)
		}

		var wg sync.WaitGroup // 스크립트 전체 조회한 대상에 대해 waitGroup 생성

		for _, s := range scripts {
			logs, logErr := client.ExecutionLog.Query().
				Where(executionlog.ScriptId(int(s.ID))).
				Order(ent.Desc(executionlog.FieldExecutedAt)). //가장 최근(Desc) 실행 시간(ExecutedAt)
				Limit(1).                                      // 하나만
				All(ctx)

			if logErr != nil {
				log.Fatalf("failed querying execution logs: %v", logErr)
			}
			var lastExecuteAt time.Time //최종 실행 시간 변수 선언 (time형)
			// 실행 이력 없는경우 불러오는 값이 없으므로 에러발생함 (index out of bound)
			if len(logs) != 0 { // 이력이 있으면 DB의 최종 실행 시간을 변수에 입력
				location, err := time.LoadLocation("Asia/Seoul")
				if err != nil {
					log.Fatalf("Failed to load location: %v", err)
				}
				latestLog := logs[0]
				lastExecuteAt = latestLog.ExecutedAt.In(location) //KST로 변환(한국시간)

				var IsAble bool = IsTimeWithinStartNEndTime(s.StartTime, s.EndTime) // script의 시작시간, 종료시간사이에 현재 시간이 있으면 true(실행) 아니면 false(다음스크립트로)
				if !IsAble {
					continue // 지정된 시간대 체크 해서 제외면 다음스크립트로
				}

			} else {
				dateStr := "19000101"
				layout := "20060102"
				// KST (UTC+9) 타임존 설정
				location, err := time.LoadLocation("Asia/Seoul")
				if err != nil {
					log.Fatalf("Failed to load location: %v", err)
				}
				parseDate, dateErr := time.ParseInLocation(layout, dateStr, location)
				if dateErr != nil {
					log.Fatalf("Date parsing Failed: %v", dateErr)
				}
				lastExecuteAt = parseDate //최종 실행 시간이 없으면 실행 주기 체크할때 null 값을 비교하므로,
				//에러가 발생하기 때문에 19000101을 최근 실행시간으로 세팅하여 무조건 실행되게함
			}

			var IsPast bool = IsTimeWithinNMinutes(lastExecuteAt, time.Duration(s.RepeatInterval)) //최종 실행 시간과 실행 주기를 계산해서 실행 해도 되는지 체크
			if !IsPast {                                                                           // 반복 주기를 지났을때만 실행 ( 현재시간-주기(분) ~ 현재시간 사이에 최종 시작일시 제외)
				wg.Add(1)                                                                // 새로운 고루틴을 추가할 때마다 WaitGroup의 카운터를 증가
				go scriptpkg.ExecuteScript(client, &wg, int(s.ID), s.Command, s.Comment) // 스크립트 실행 및 에러시 메일 서버 호출
			}
		}

		wg.Wait() // 모든 고루틴이 완료될 때까지 대기 (한사이클 돌때까지)
		log.Println("All scripts executed")
		time.Sleep(5 * time.Minute) // 5분에 한번씩 체크
	}
}

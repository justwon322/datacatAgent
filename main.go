// main.go
package main

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

func IsTimeWithinNMinutes(t time.Time, minutes time.Duration) bool {
	now := time.Now()                               // 현재 시간을 가져옵니다.
	adjustedTime := now.Add(-minutes * time.Minute) // 현재 시간에서 minutes 분을 뺀 시간을 계산합니다.
	return t.After(adjustedTime) && t.Before(now)   // t가 adjustedTime 이후이면서 now 이전인지 확인합니다.
}

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
		log.Fatalf("Failed to Parse startDate: %v", err)
	}
	endTime, endDateErr := time.ParseInLocation(layout, comparableEndDate, location)
	if endDateErr != nil {
		log.Fatalf("Failed to Parse endDateErr: %v", err)
	}

	return now.After(startTime) && now.Before(endTime) // t가 adjustedTime 이후이면서 now 이전인지 확인합니다.
}

func main() {
	client, err := db.OpenDB()
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	for true {
		// usable이 1인 스크립트만 쿼리
		scripts, err := client.Script.Query().
			Where(script.Usable(1)).
			All(context.Background())
		if err != nil {
			log.Fatalf("failed getting scripts: %v", err)
		}

		var wg sync.WaitGroup

		for _, s := range scripts {
			logs, logErr := client.ExecutionLog.Query().
				Where(executionlog.ScriptId(int(s.ID))).
				Order(ent.Desc(executionlog.FieldExecutedAt)).
				Limit(1).
				All(ctx)

			if logErr != nil {
				log.Fatalf("failed querying execution logs: %v", err)
			}
			var lastExecuteAt time.Time
			// 실행 이력 없는경우 에러발생함 (index out of bound)
			if len(logs) != 0 {
				location, err := time.LoadLocation("Asia/Seoul")
				if err != nil {
					log.Fatalf("Failed to load location: %v", err)
				}
				latestLog := logs[0]
				lastExecuteAt = latestLog.ExecutedAt.In(location)

				var IsAble bool = IsTimeWithinStartNEndTime(s.StartTime, s.EndTime)
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
				lastExecuteAt = parseDate
			}

			var IsPast bool = IsTimeWithinNMinutes(lastExecuteAt, time.Duration(s.RepeatInterval))
			if !IsPast { // 반복 주기를 지났을때만 실행 ( 현재시간-주기(분) ~ 현재시간 사이에 최종 시작일시 제외)
				wg.Add(1) // 새로운 고루틴을 추가할 때마다 WaitGroup의 카운터를 증가
				go scriptpkg.ExecuteScript(client, &wg, int(s.ID), s.Command, s.Comment)
			}
		}

		wg.Wait() // 모든 고루틴이 완료될 때까지 대기
		log.Println("All scripts executed")
		time.Sleep(5 * time.Minute) // 5분에 한번씩 체크
	}
}

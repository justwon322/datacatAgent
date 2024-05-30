// main.go
package main

import (
	"context"
	"datacatAgent/ent"
	"datacatAgent/ent/executionlog"
	"datacatAgent/ent/script"
	scriptpkg "datacatAgent/scriptExecutor"
	"log"
	"sync"
	"time"

	"datacatAgent/util"
)

func IsTimeWithinNMinutes(t time.Time, minutes time.Duration) bool {
	now := time.Now()
	adjustedTime := now.Add(minutes * time.Minute)
	return t.After(adjustedTime) && t.Before(now)
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
				latestLog := logs[0]
				lastExecuteAt = latestLog.ExecutedAt
			} else {
				dateStr := "19000101"
				layout := "20060102"
				parseDate, dateErr := time.Parse(layout, dateStr)
				if dateErr != nil {
					log.Fatalf("Date parsing Failed: %v", dateErr)
				}
				lastExecuteAt = parseDate
			}

			var IsPast bool = IsTimeWithinNMinutes(lastExecuteAt, time.Duration(s.RepeatInterval))
			if !IsPast { // 반복 주기를 지났을때만 실행
				wg.Add(1) // 새로운 고루틴을 추가할 때마다 WaitGroup의 카운터를 증가
				go scriptpkg.ExecuteScript(client, &wg, int(s.ID), s.Command, s.Comment)
			}
		}

		wg.Wait() // 모든 고루틴이 완료될 때까지 대기
		log.Println("All scripts executed")
		time.Sleep(5 * time.Minute) // 5분에 한번씩 체크
	}
}

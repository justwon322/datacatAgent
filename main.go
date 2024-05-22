// main.go
package main

import (
	"context"
	"datacatAgent/ent/script"
	"log"
	"sync"
	"time"

	scriptpkg "datacatAgent/scriptExecutor"
	"datacatAgent/util"
)

func main() {
	client, err := db.OpenDB()
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}
	defer client.Close()

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
			wg.Add(1) // 새로운 고루틴을 추가할 때마다 WaitGroup의 카운터를 증가
			go scriptpkg.ExecuteScript(client, &wg, int(s.ID), s.Command, s.Comment)
		}

		wg.Wait() // 모든 고루틴이 완료될 때까지 대기
		log.Println("All scripts executed")
		time.Sleep(5 * time.Minute) // 5분에 한번씩 체크
	}
}

# Datacat Agent
## 주요기능

1. shell script 실행 및 정상 유무 판단 가능
2. 비정상 실행시 담당자 알림 기능(NotificationSender와 연계)

## 실행 방법
.env 파일에
<br/>
DB_HOST=?
<br/>
DB_PORT=?
<br/>
DB_USER=?
<br/>
DB_PASS=?
<br/>
DB_NAME=?
<br/>
입력 및 루트 폴더 아래에 저장 후
main.go 실행 
dependency는 go.mod 참조
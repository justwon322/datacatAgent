# Datacat Agent
## 주요기능

1. shell script 실행 및 정상 유무 판단 가능
2. 비정상 실행시 담당자 알림 기능(NotificationSender와 연계)

## 필요사양
go 1.22.X 버전 설치 필요

## 설정 방법
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
입력 및 루트 폴더 아래에 저장
dependency는 go.mod 참조,

## 설치 방법
build.sh 실행하여 binary 파일 생성 후 설치 하면 됨(OS및 CPU 아키텍쳐 확인 필수)

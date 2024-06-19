# Datacat Agent
## 주요기능

1. shell script 실행 및 정상 유무 판단 가능
2. 비정상 실행시 담당자 알림 기능(NotificationSender와 연계)

## 필요사양
go 1.22.X 버전 설치 필요

## 설정 방법
.env 파일에
```dotenv
DB_HOST=?
DB_PORT=?
DB_USER=?
DB_PASS=?
DB_NAME=?
```
입력 및 루트 폴더 아래에 .env 라고 생성 후 저장 
dependency는 go.mod 참조,

## 설치(빌드) 방법
build.sh 실행시 루트폴더 하위에 datacatAgent.exe 생성됨(윈도우일경우)
아니면 루트폴더에서 아래 명령어 실행해도 동일하게 바이너리 파일 생성됨
```shell
go build
```

## 실행 방법
datacatAgent.exe를 더블클릭하여 실행해도 되고 </br>
datacatAgent.vbs를 더블클릭하여 백그라운드 실행도 가능함

## schema 변경
ent > schema 폴더에서 스키마에 대한 내용이 변경이 있을경우 변경후 아래 스크립트 실행 필요함
```shell
go generate ./ent #루트폴더에서 해당 명령어 실행
```

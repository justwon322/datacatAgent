package main

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"log"
	"time"
)

func thesis() {
	// Rod 라이브러리를 사용해 브라우저 연결
	url := launcher.New().MustLaunch()
	browser := rod.New().ControlURL(url).MustConnect()

	// 새 탭을 열고 웹페이지 로딩
	pageURL := "<CHANGE WHAT YOU WANT>" // 원하는 URL로 변경하세요.
	page := browser.MustPage(pageURL)

	id := "#user_id"                  // 비밀번호 입력 필드의 선택자
	loginButton := "#login"           // 로그인 버튼의 선택자
	searchKeyword := "#input-keyword" // 검색키워드 필드 선택자

	page.MustElement(id).MustInput("<ID>")

	// 로그인 버튼 클릭
	page.MustElement(loginButton).MustClick()
	page.MustElement(searchKeyword).MustInput("pro").MustKeyActions().Press(input.Enter).MustDo()

	//page.MustElement(firstPaperLink).MustScrollIntoView().MustClick()

	// 웹페이지의 스크린샷 생성
	now := time.Now()
	date := now.Format("20060102")
	//page.MustWaitLoad().MustScreenshotFullPage("<CHANGE WHAT YOU WANT>")

	page = browser.MustPage("<CHANGE WHAT YOU WANT>")
	page.MustWaitLoad().MustScreenshotFullPage("<CHANGE WHAT YOU WANT>")
	log.Println("Succeeded")
	// 리소스 정리
	browser.MustClose()

}

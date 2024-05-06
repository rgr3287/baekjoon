package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 현재 시간과 시작 시간을 저장할 변수
	var now, start int

	// 현재 시간 입력 받기
	nowInput, _ := reader.ReadString('\n')
	nowInput = strings.TrimSpace(nowInput)
	nowParts := strings.Split(nowInput, ":")
	h, _ := strconv.Atoi(nowParts[0])
	m, _ := strconv.Atoi(nowParts[1])
	s, _ := strconv.Atoi(nowParts[2])
	now = h*3600 + m*60 + s

	// 시작 시간 입력 받기
	startInput, _ := reader.ReadString('\n')
	startInput = strings.TrimSpace(startInput)
	startParts := strings.Split(startInput, ":")
	h, _ = strconv.Atoi(startParts[0])
	m, _ = strconv.Atoi(startParts[1])
	s, _ = strconv.Atoi(startParts[2])
	start = h*3600 + m*60 + s

	// 앞으로 남은 시간 계산
	var time int
	if start > now {
		time = start - now
	} else {
		time = 24*3600 - (now - start)
	}

	// 결과 출력
	fmt.Printf("%02d:%02d:%02d", time/3600, (time/60)%60, time%60)
}

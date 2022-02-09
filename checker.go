package main

import (
	"net/http"
	"time"
)

func checker(url string) (int, error) {
	//url := "http://xml.kishou.go.jp/revise.html"
	resp, err := http.Get(url)
	if err != nil {
		return 2, err
	}
	defer resp.Body.Close()

	// 過去24時間の間に更新があったか確認
	// 実行中に更新されてしまうケースを加味し、Endは2分程度未来時間とする
	now := time.Now()
	start := now.AddDate(0, 0, -1)
	end := now.Add(time.Minute * 2)

	// ページの更新日を取得
	modified := resp.Header["Last-Modified"][0]
	modified_date, err := time.Parse(time.RFC1123, modified)
	if err != nil {
		return 2, err
	}

	// 更新日が過去24時間に収まっている場合は更新ありと判断
	if start.Before(modified_date) && end.After(modified_date) {
		return 0, nil
	} else {
		return 1, nil
	}
}

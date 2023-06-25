package main

import (
	"github.com/daan0220/oneday-demo-server/cmd"
)

func main() {
	router := cmd.NewRouter()

	// サーバーを開始
	if err := router.Start(":1324"); err != nil {
		// エラーハンドリング
		panic(err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"golang/internal/widget"
	"os"
	"strconv"
)

func main() {
	// 呼び出し元を変更する。
	var x int
	x, err := strconv.Atoi(Input("実行するメソッドの番号を入力して下さい。"))

	if err != nil {
		panic(err.Error())
	}

	switch x {
	case 1:
		widget.Hello()
	case 2:
		widget.Click()
	case 3:
		widget.Entry()
	}
}

// 入力メソッド
func Input(m string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(m + ": ")
	scanner.Scan()
	return scanner.Text()
}

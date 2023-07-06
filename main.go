package main

import (
	//"bufio"
	"fmt"
	"os"
	"github.com/gdamore/tcell/v2"
	"log"
)

func saveTextToFile(filename string, input string) {
	// filenameテキストファイルにinputの内容を保存する。
	f, err := os.Create(filename)

	data := []byte(input + "\n")
	count, err := f.Write(data)
	_ = count

	if err != nil {
		fmt.Println(err)
	}
}


func main() {
	// ファイル名指定がなければ終了
	if len(os.Args) != 2 {
			return
	}

	var filename string = os.Args[1]

	
	// tcell初期化
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	screen.Clear()

	var text string = ""


	// 終了処理
	quit := func() {
		maybePanic := recover()
		screen.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}

	// main関数終了時に実行
	defer quit()

	// var text string = ""
	var cursorX int = 0
	var cursorY int = 0


	// mainループ
	for {
		screen.Show()

		ev := screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				// 保存せず終了
				return
			} else if ev.Key() == tcell.KeyEnter {
				// 改行処理
				cursorY++
				cursorX = 0
				text += "\n"
			} else if ev.Key() == tcell.KeyCtrlS {
				// 保存処理
				saveTextToFile(filename, text)
				return
				// TODO: バックスペース処理
			} else {
				// 入力処理
				var c rune = ev.Rune()
				screen.SetContent(cursorX, cursorY, c, nil, tcell.StyleDefault)
				text += string(c)
				cursorX++
			}
		}
	}
}

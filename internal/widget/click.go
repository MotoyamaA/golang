package widget

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Click() {
	c := 0
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	w.SetContent(
		// VBOXは縦に並べる、HBOXは横に並べる
		widget.NewVBox(
			l,
			// clickmeボタンを作成
			widget.NewButton("Click! me!", func() {
				c++
				// ボタンを押すとカウントアップしてラベルを書き換えます
				l.SetText("count : " + strconv.Itoa(c))
			}),
		),
	)

	w.ShowAndRun()
}

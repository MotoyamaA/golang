package widget

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func Entry() {
	a := app.New()
	w := a.NewWindow("入力した数まで1から足していきます。")
	l := widget.NewLabel("Hello Fyne!")
	e := widget.NewEntry()
	e.SetText("0")
	w.SetContent(
		// VBOXは縦に並べる、HBOXは横に並べる
		widget.NewVBox(
			l, e,
			// clickmeボタンを作成
			widget.NewButton("Click! me!", func() {
				n, _ := strconv.Atoi(e.Text)
				// ボタンを押すとカウントアップしてラベルを書き換えます
				l.SetText("count : " + strconv.Itoa(total(n)))
			}),
		),
	)
	a.Settings().SetTheme(theme.LightTheme())
	w.ShowAndRun()
}

func total(n int) int {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	return t
}

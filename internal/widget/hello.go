package widget

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func Hello() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		// VBOXは縦に並べる、HBOXは横に並べる
		widget.NewVBox(
			widget.NewLabel("Hello Fyne!"),
			widget.NewLabel("This is sample application!"),
		),
	)

	w.ShowAndRun()
}

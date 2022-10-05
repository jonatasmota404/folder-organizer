package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("Clock")
	w.Resize(fyne.Size{Width: 250, Height: 100})
	w.SetMaster()

	clock := widget.NewLabel("")
	updateTime(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.SetContent(clock)
	w.Show()

	w2 := a.NewWindow("TESTE")
	w2.SetContent(widget.NewLabel("More content"))
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Show()
	}))
	w2.Show()

	a.Run()
	//w.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}

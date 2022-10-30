package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")
	myWindow.Resize(fyne.Size{Width: 500, Height: 500})

	//box := container.NewHBox(combo)

	split := makeSplitTab(myWindow)

	myWindow.SetContent(split)
	myWindow.Show()

	myApp.Run()
}


func makeMyContainer(window fyne.Window) fyne.CanvasObject {
	dataList := binding.BindFloatList(&[]float64{0.1, 0.2, 0.3})

	button := widget.NewButton("Append", func() {
		dataList.Append(float64(dataList.Length()+1) / 10)
	})

	list := widget.NewListWithData(dataList,
		func() fyne.CanvasObject {
			return container.NewBorder(nil, nil, nil, widget.NewButton("+", nil),
				widget.NewLabel("item x.y"))
		},
		func(item binding.DataItem, obj fyne.CanvasObject) {
			f := item.(binding.Float)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			text.Bind(binding.FloatToStringWithFormat(f, "item %0.1f"))

			btn := obj.(*fyne.Container).Objects[1].(*widget.Button)
			btn.OnTapped = func() {
				val, _ := f.Get()
				_ = f.Set(val + 1)
			}
		})
	
	listPanel := container.NewBorder(nil, button, nil, nil, list)

	return listPanel
}

func makeSplitTab(window fyne.Window) fyne.CanvasObject {
	left := makeMyContainer(window)//widget.NewMultiLineEntry()
	//left.Wrapping = fyne.TextWrapWord
	//left.SetText("Long text is looooooooooooooong")
	right := container.NewVSplit(
		widget.NewLabel("Label"),
		widget.NewButton("Button", func() { fmt.Println("button tapped!") }),
	)
	return container.NewHSplit(container.NewVScroll(left), right)
}
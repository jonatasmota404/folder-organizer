package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	var data = []string{"a", "string", "list"}

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			//return widget.NewLabel("template")
			return container.NewAdaptiveGrid(3, widget.NewLabel("item x"),
			widget.NewButton("+", nil),
			widget.NewButton("-", nil))
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			//o.(*widget.Label).SetText(data[i])
			btnAdd := o.(*fyne.Container).Objects[2].(*widget.Button)
			btnAdd.OnTapped = func() {
				fmt.Println(i)
			}

			btnRemove := o.(*fyne.Container).Objects[2].(*widget.Button)
			btnRemove.OnTapped = func() {
				fmt.Println(i)
			}

	
		})
	
	button := widget.NewButton("Append", func() {
		val := fmt.Sprintf("Item %d", len(data) +1)
		data = append(data, val)
		list.Refresh()
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
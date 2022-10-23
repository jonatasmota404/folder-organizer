package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choice Widgets")

	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
		log.Println("Select set to", value)
	})

	separator := widget.NewSeparator()

	addFolderToOrganize := widget.NewButton("Folder Open", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err != nil {
				fmt.Println(err)
				return
			}
			if list == nil {
				log.Println("Cancelled")
				return
			}

			fmt.Println(list)
		}, myWindow)
	})

	openWindowToAddFiles := widget.NewButton("Open window", func() {
		window := myApp.NewWindow("teste");
		window.Show()

		extensionPath := widget.NewEntry()

		path := widget.NewEntry()

		buttonFolder := widget.NewButton("open folder", func() {
			dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println(lu.Path())
				path.SetText(lu.Path())
			}, window)
		})

		window.SetContent(container.NewVBox(extensionPath, path, buttonFolder))
	})

	
	myWindow.SetContent(container.NewVBox(combo, addFolderToOrganize, separator, openWindowToAddFiles))
	myWindow.Show()

	myApp.Run()
}

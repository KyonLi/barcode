package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/pdf417"
	"github.com/boombuler/barcode/qr"
	"image"
)

var (
	t = "QR"
	canvasImageInstance *canvas.Image
	imageContainerInstance *fyne.Container
	entryInstance *widget.Entry
	radioInstance *widget.Radio
)

func canvasImage() *canvas.Image {
	if canvasImageInstance == nil {
		i := canvas.NewImageFromImage(image.NewAlpha(image.Rectangle{}))
		canvasImageInstance = i
		return i
	} else {
		return canvasImageInstance
	}
}

func imageContainer() *fyne.Container {
	if imageContainerInstance == nil {
		container := fyne.NewContainer(canvasImage())
		container.Layout = layout.NewFixedGridLayout(fyne.NewSize(450, 450))
		imageContainerInstance = container
		return container
	} else {
		return imageContainerInstance
	}
}

func entry() *widget.Entry {
	if entryInstance == nil {
		entry := widget.NewEntry()
		entry.OnChanged = func(s string) {
			if len(s) > 0 {
				switch t {
				case "QR":
					code, _ := qr.Encode(s, qr.M, qr.Auto)
					code, _ = barcode.Scale(code, 450, 450)
					canvasImage().Image = code
					canvasImage().Resize(fyne.NewSize(450, 450))
					canvasImage().Move(fyne.NewPos(0, 0))
				case "PDF417":
					code, _ := pdf417.Encode(s, 4)
					code, _ = barcode.Scale(code, 450, 120)
					canvasImage().Image = code
					canvasImage().Resize(fyne.NewSize(450, 120))
					canvasImage().Move(fyne.NewPos(0, (450-120)/2))
				}
			} else {
				canvasImage().Image = image.NewAlpha(image.Rectangle{})
			}
			canvas.Refresh(canvasImage())
		}
		entryInstance = entry
		return entry
	} else {
		return entryInstance
	}
}

func radio() *widget.Radio {
	if radioInstance == nil {
		r := widget.NewRadio([]string{"QR", "PDF417"}, func(s string) {
			t = s
			entry().OnChanged(entry().Text)
		})
		r.SetSelected(t)
		radioInstance = r
		return r
	} else {
		return radioInstance
	}
}

func main() {
	application := app.New()
	w := application.NewWindow("Barcode Generator")
	w.SetContent(widget.NewVBox(
		imageContainer(),
		radio(),
		widget.NewLabel("Input Text"),
		entry(),
		widget.NewButton("Clear", func() {
			entry().SetText("")
		}),
	))
	w.ShowAndRun()
}

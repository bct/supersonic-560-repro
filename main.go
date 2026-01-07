package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

const pollFrequency = 500 * time.Millisecond
const sliderWidth = 1000

func main() {
	fyneApp := app.New()
	mainWindow := fyneApp.NewWindow("supersonic #560 repro")

	slider := widget.Slider{
		Value:       0,
		Min:         0,
		Max:         1,
		Step:        0.002,
		Orientation: widget.Horizontal,
	}
	mainWindow.SetContent(&slider)

	//start a loop to update the play time
	progress := 0.0

	pollingTick := time.NewTicker(pollFrequency)
	go func() {
		for {
			select {
			case <-pollingTick.C:
				progress += 1
				if progress > sliderWidth {
					progress = 0
				}

				log.Printf("tick %v", progress)
				fyne.Do(func() {
					log.Printf("tick %v - start update", progress)
					slider.SetValue(progress / sliderWidth)
					log.Printf("tick %v - finish update", progress)
				})
			}
		}
	}()

	// display the window
	mainWindow.ShowAndRun()
}

package main

import "routingnav/ui"

func main() {
	a := ui.NewApp()
	a.Run()

	a.Cleanup()
}

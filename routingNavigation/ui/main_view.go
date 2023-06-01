package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func mainView(ctx *AppContext) *fyne.Container {
	return container.NewCenter(
		container.NewVBox(
			widget.NewLabel("Main"),
			container.NewHBox(
				widget.NewButton("View1",
					func() {
						ctx.Push(List)
					},
				),
				widget.NewButton("to Details",
					func() {
						ctx.NavigateToWithParam(NewGame, "from main")
					},
				),
			),
		),
	)
}

func otherView(ctx *AppContext) *fyne.Container {
	return container.NewCenter(
		container.NewVBox(
			widget.NewLabel("View1"),
			widget.NewEntry(),
			container.NewHBox(
				widget.NewButton("Back",
					func() {
						ctx.Pop()
					},
				),
				widget.NewButton("to Details",
					func() {
						ctx.PushWithParam(NewGame, "from view 1")
					},
				),
			),
		),
	)
}

func otherView2(ctx *AppContext) *fyne.Container {
	param := fmt.Sprintf("param: %s", ctx.RouteParam.(string))

	backBtn := widget.NewButton("Back",
		func() {
			ctx.Pop()
		},
	)
	if ctx.NavStack.Size() < 1 {
		backBtn.Disable()
	}

	return container.NewCenter(
		container.NewVBox(
			widget.NewLabel("DetailsView"),
			widget.NewLabel(param),
			container.NewVBox(
				backBtn,
				widget.NewButton("To Main",
					func() {
						ctx.NavigateTo(Main)
					},
				),
			),
		),
	)
}

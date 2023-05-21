package main

import (
	"todoapp/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")
	w.Resize(fyne.NewSize(300, 400))

	data := []models.Todo{
		models.NewTodo("Some stuff"),
		models.NewTodo("Some more stuff"),
		models.NewTodo("Some other things"),
	}
	todos := binding.NewUntypedList()
	for _, t := range data {
		todos.Append(t)
	}
	newtodoDescTxt := widget.NewEntry()
	newtodoDescTxt.PlaceHolder = "New Todo Description..."
	addBtn := widget.NewButton("Add", func() {
		todos.Append(models.NewTodo(newtodoDescTxt.Text))
		newtodoDescTxt.Text = ""
	})
	addBtn.Disable()
	newtodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container
			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // Left
				// RIGHT ↓
				addBtn,
				// take the rest of the space
				newtodoDescTxt,
			),
			nil, // Right
			nil, // Left
			// the rest will take all the rest of the space
			widget.NewListWithData(
				// the binding.List type
				todos,
				// func that returns the component structure of the List Item
				// exactly the same as the Simple List
				func() fyne.CanvasObject {
					return container.NewBorder(
						nil, nil, nil,
						// left of the border
						widget.NewCheck("", func(b bool) {}),
						// takes the rest of the space
						widget.NewLabel(""),
					)
				},
				// func that is called for each item in the list and allows
				// but this time we get the actual DataItem we need to cast
				func(di binding.DataItem, o fyne.CanvasObject) {
					ctr, _ := o.(*fyne.Container)
					// ideally we should check `ok` for each one of those casting
					// but we know that they are those types for sure
					l := ctr.Objects[0].(*widget.Label)
					c := ctr.Objects[1].(*widget.Check)
					/*
						diu, _ := di.(binding.Untyped).Get()
						todo := diu.(models.Todo)
					*/
					todo := models.NewTodoFromDataItem(di)
					l.SetText(todo.Description)
					c.SetChecked(todo.Done)
				}),
		),
	)
	w.ShowAndRun()
}

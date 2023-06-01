package ui_test

import (
	"routingnav/ui"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialContextState(t *testing.T) {
	ctx := ui.NewAppContext(ui.Main, WindowMock{})
	assert.Equal(t, ui.Main, ctx.CurrentRoute())

	assert.IsType(t, WindowMock{}, ctx.GetWindow())
}

func TestRouting(t *testing.T) {
	ctx := ui.NewAppContext(ui.Main, WindowMock{})
	assert.Equal(t, ui.Main, ctx.CurrentRoute())
	ctx.NavigateTo(ui.List)
	assert.Equal(t, ui.List, ctx.CurrentRoute())

	ctx.NavigateToWithParam(ui.NewGame, "someId")
	assert.Equal(t, ui.NewGame, ctx.CurrentRoute())
	assert.Equal(t, "someId", ctx.RouteParam.(string))

	ctx.NavigateTo(ui.List)
	assert.Equal(t, ui.List, ctx.CurrentRoute())
	assert.Nil(t, ctx.RouteParam)
}

//TODO: add testing for push/pop routes

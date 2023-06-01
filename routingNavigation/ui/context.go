package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type ViewMode uint8

const (
	Replace ViewMode = iota
	Push
	Pop
)

type AppContext struct {
	Version string
	//TODO: maybe those 3 props can be handled by Navstack
	Route      binding.String
	RouteParam any
	RouteMode  ViewMode
	//

	NavStack *NavStack

	w fyne.Window
}

func NewAppContext(initialRoute AppRoute, window fyne.Window) AppContext {
	route := initialRoute.String()
	return AppContext{
		Route:    binding.BindString(&route),
		NavStack: NewNavStack(),
		w:        window,
	}
}

func (c *AppContext) GetClipboard() fyne.Clipboard {
	return c.w.Clipboard()
}

func (c *AppContext) GetWindow() fyne.Window {
	return c.w
}

func (c *AppContext) OnRouteChange(callback func()) {
	c.Route.AddListener(binding.NewDataListener(callback))
}

func (c *AppContext) CurrentRoute() AppRoute {
	r, _ := c.Route.Get()
	return RouteFromString(r)
}

func (c *AppContext) NavigateTo(route AppRoute) {
	c.RouteParam = nil
	c.RouteMode = Replace
	c.Route.Set(route.String())
}

func (c *AppContext) NavigateToWithParam(route AppRoute, param any) {
	c.RouteParam = param
	c.RouteMode = Replace
	c.Route.Set(route.String())
}

func (c *AppContext) PushWithParam(route AppRoute, param any) {
	c.NavStack.Push(NewNavStackItem(c.CurrentRoute(), c.RouteParam))
	c.RouteParam = param
	c.RouteMode = Push
	c.Route.Set(route.String())
}

func (c *AppContext) Push(route AppRoute) {
	c.NavStack.Push(NewNavStackItem(c.CurrentRoute(), c.RouteParam))
	c.RouteMode = Push
	c.RouteParam = nil
	c.Route.Set(route.String())
}

func (c *AppContext) Pop() {
	nsi, ok := c.NavStack.Pop()
	if !ok {
		//TODO: instead of panic maybe should revert to a base view?
		panic("Trying to pop even tho you have no views on the stack")
	}

	c.RouteMode = Pop
	c.RouteParam = nsi.routeParam
	c.Route.Set(nsi.route.String())
}

func (c *AppContext) CacheViewOnStack(content fyne.CanvasObject) {
	//TODO: maybe return an error or something if this Peek fails
	i, ok := c.NavStack.Peek()
	if !ok {
		return
	}

	i.SetContent(content)
}

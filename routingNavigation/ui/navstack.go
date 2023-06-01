package ui

import "fyne.io/fyne/v2"

type NavStackItem struct {
	route      AppRoute
	routeParam any
	content    fyne.CanvasObject
}

func (nsi *NavStackItem) SetContent(c fyne.CanvasObject) {
	nsi.content = c

}

func NewNavStackItem(route AppRoute, routeParam any) *NavStackItem {
	return &NavStackItem{
		route:      route,
		routeParam: routeParam,
		content:    nil,
	}
}

type NavStack struct {
	stack         []*NavStackItem
	latestContent fyne.CanvasObject
}

func NewNavStack() *NavStack {
	return &NavStack{
		stack: []*NavStackItem{},
	}
}

func (ns *NavStack) Clear() {
	ns.stack = []*NavStackItem{}
}

func (ns *NavStack) Push(item *NavStackItem) {
	ns.stack = append(ns.stack, item)
}

func (ns *NavStack) Pop() (*NavStackItem, bool) {
	if ns.Size() < 1 {
		return nil, false
	}

	last := ns.lastIndex()
	nsi := ns.stack[last]
	ns.stack = ns.stack[:last]
	ns.latestContent = nsi.content
	return nsi, true
}

func (ns *NavStack) lastIndex() int {
	return len(ns.stack) - 1
}

func (ns *NavStack) Size() int {
	return len(ns.stack)
}

func (ns *NavStack) GetPopContent() fyne.CanvasObject {
	c := ns.latestContent
	ns.latestContent = nil
	//todo: bool return too
	return c
}

func (ns *NavStack) Peek() (*NavStackItem, bool) {
	if len(ns.stack) < 1 {
		return nil, false
	}
	return ns.stack[len(ns.stack)-1], true
}

package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/svg"
	"github.com/nathanhack/svg/attr"
	elm "github.com/nathanhack/svg/elem"
)

//go:generate env GOARCH=wasm GOOS=js go build -o ../static/main.wasm

type Body struct {
	vecty.Core
}

type Item struct {
	vecty.Core
}

func (i *Item) Mount() {
	vecty.Rerender(i)
}

func (i *Item) Render() vecty.ComponentOrHTML {
	return elem.Div(
		svg.Render(
			svg.SVG(
				attr.Width(300),
				attr.Height(100),
				attr.ViewBox(0, 0, 300, 100),
				attr.Stroke("red"),
				attr.Fill("grey"),
				elm.Circle(attr.Cx(50), attr.Cy(50), attr.R(40)),
				elm.Circle(attr.Cx(150), attr.Cy(50), attr.R(4)),
			),
		),
	)
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(&Item{})
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}

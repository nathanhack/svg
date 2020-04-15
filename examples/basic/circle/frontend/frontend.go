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
				attr.ViewBox(0, 0, 300, 100),
				attr.Stroke("red"),
				attr.Fill("grey"),
				elm.Circle(attr.Cx(50), attr.Cy(50), attr.R(40)),
				elm.Circle(attr.Cx(150), attr.Cy(50), attr.R(4)),

				svg.SVG(
					attr.ViewBox(0, 0, 10, 10),
					attr.Width(100),
					elm.Circle(attr.Cx(5), attr.Cy(5), attr.R(4)),
				),
			),
		),
	)
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		vecty.Tag(
			"svg",
			vecty.Markup(
				vecty.Attribute("viewbox", " 0 0 100 100"),
				vecty.Attribute("fill", "grey"),
				vecty.Attribute("stroke", "red"),
			),
			vecty.Tag(
				"circle", vecty.Markup(
					vecty.Attribute("cx", "50"),
					vecty.Attribute("cy", "50"),
					vecty.Attribute("r", "40"),
				),
			),
		),
	)
}

func main() {
	vecty.SetTitle("Simple Example")
	vecty.RenderBody(&Body{})
}

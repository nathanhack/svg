package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/svg"
	"github.com/nathanhack/svg/attr"
	"github.com/nathanhack/svg/attr/transforms"
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
				attr.Width(100),
				attr.Height(100),
				attr.ViewBox(0, 0, 20, 20),
				attr.Fill("black"),
				elm.Circle(
					attr.Cx(10),
					attr.Cy(10),
					attr.R(10),
				),
				elm.Circle(
					attr.Cx(10),
					attr.Cy(10),
					attr.R(5),
					attr.Stroke("#FF6347"),
					attr.StrokeWidth(10),
					attr.StrokeDasharray(35*31.4/100, 31.4),
					attr.Transform(transforms.Rotate(-90, 10, 10)),
				),
			),
		),
	)
}

func (b Body) Render() vecty.ComponentOrHTML {
	return elem.Body(&Item{})
}

func main() {
	vecty.SetTitle("Simple Pie chart Example")
	vecty.RenderBody(&Body{})
}

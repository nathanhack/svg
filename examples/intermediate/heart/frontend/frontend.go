package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/nathanhack/svg"
	"github.com/nathanhack/svg/attr"
	"github.com/nathanhack/svg/attr/path"
	elm "github.com/nathanhack/svg/svgelem"
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
				attr.Stroke("red"),
				attr.Fill("grey"),
				attr.ViewBox(0, 0, 100, 100),
				elm.Path(
					attr.Fill("purple"),
					attr.Stroke("red"),
					attr.D(
						//D takes multiple paths
						// Path's can be chained together and creates a "path"
						path.M(10, 30).
							A(20, 20, 0, 0, 1, 50, 30).
							A(20, 20, 0, 0, 1, 90, 30).
							Q(90, 60, 50, 90).
							Q(10, 60, 10, 30).
							Z(),
						// Or two or more paths can be added together
						path.M(50, 50)+path.H(100),
						// or each path can be added separately
						path.M(50, 50),
						path.V(100),
						//and as always literal strings work too
						"M 50 50 h -25",
					),
				),
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

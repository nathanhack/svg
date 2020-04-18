package svg

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/svg/internal"
)

var Namespace = "http://www.w3.org/2000/svg"

type Number interface{}

type Length interface{}

type LengthOrPercent interface{}

type NumberOrPercent interface{}

func Percent(number Number) interface{} {
	return internal.Stringify(number, "%")
}

type Component interface {
	svg()
}

type Element interface {
	Component
	Tag() string
	Attributes() []Attribute
	Elements() []Element
	Inner() string
}

type Root interface {
	Element
}

type Core struct {
	vecty.Core
}

func (c *Core) svg() {}

type svgComponent struct {
	Core
	elements []Element
	attrs    []Attribute
}

func (s *svgComponent) Inner() string {
	return ""
}

func (s *svgComponent) Attributes() []Attribute {
	return s.attrs
}

func (s *svgComponent) Tag() string {
	return "svg"
}

func (s *svgComponent) Elements() []Element {
	return s.elements
}

func SVG(elementsOrComponents ...Component) Root {
	attrs := make([]Attribute, 0)
	el := make([]Element, 0)

	for _, x := range elementsOrComponents {
		switch x.(type) {
		case Element:
			el = append(el, x.(Element))
		case Attribute:
			tmp := x.(Attribute)
			attrs = append(attrs, tmp)
		}
	}

	return &svgComponent{
		elements: el,
		attrs:    attrs,
	}
}

type Attribute interface {
	Component
	vecty.Applyer
}

func Render(svg Root) vecty.ComponentOrHTML {
	return render(svg)
}

func render(el Element) vecty.ComponentOrHTML {
	markups := []vecty.Applyer{vecty.Namespace(Namespace)}
	items := make([]vecty.MarkupOrChild, 0)

	for _, a := range el.Attributes() {
		markups = append(markups, a)
	}

	for _, e := range el.Elements() {
		items = append(items, render(e))
	}

	items = append(items, vecty.Markup(markups...))

	return vecty.Tag(el.Tag(), items...)
}

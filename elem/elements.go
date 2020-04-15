package elem

import "github.com/nathanhack/svg"

type circle struct {
	svg.Core
	attrs []svg.Attribute
}

func (c circle) Tag() string {
	return "circle"
}

func (c circle) Attributes() []svg.Attribute {
	return c.attrs
}

func (c circle) Elements() []svg.Element { return nil }

func Circle(attrs ...svg.Attribute) svg.Element {
	return &circle{
		attrs: attrs,
	}
}

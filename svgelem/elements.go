package svgelem

import "github.com/nathanhack/svg"

type element struct {
	svg.Core
	Name  string
	Attrs []svg.Attribute
	Elems []svg.Element
	In    string
}

func (e *element) Inner() string {
	return e.Inner()
}

func (e *element) Tag() string {
	return e.Name
}

func (e *element) Attributes() []svg.Attribute {
	return e.Attrs
}

func (e *element) Elements() []svg.Element {
	return e.Elems
}

func attrsOnly(tag string, attrs ...svg.Attribute) svg.Element {
	return &element{
		Name:  tag,
		Attrs: attrs,
	}
}

//Element is for creating generic elements not currently defined in the package
func Element(tag string, inner string, attrsOrElements ...svg.Component) svg.Element {
	attrs := make([]svg.Attribute, 0)
	elements := make([]svg.Element, 0)

	for _, x := range attrsOrElements {
		switch x.(type) {
		case svg.Element:
			elements = append(elements, x.(svg.Element))
		case svg.Attribute:
			tmp := x.(svg.Attribute)
			attrs = append(attrs, tmp)
		}
	}

	return &element{
		Name:  tag,
		Attrs: attrs,
		Elems: elements,
		In:    inner,
	}
}

///////////////////////////////////////////////////

func A(attrsOrElements ...svg.Component) svg.Element {
	return Element("a", "", attrsOrElements...)
}

func Animate(attrsOrElements ...svg.Component) svg.Element {
	return Element("animate", "", attrsOrElements...)
}

func AnimateMotion(attrsOrElements ...svg.Component) svg.Element {
	return Element("animateMotion", "", attrsOrElements...)
}
func Circle(attrsOrElements ...svg.Component) svg.Element {
	return Element("circle", "", attrsOrElements...)
}

func ClipPath(attrsOrElements ...svg.Component) svg.Element {
	return Element("clipPath", "", attrsOrElements...)
}

func Defs(elements ...svg.Element) svg.Element {
	return &element{
		Name:  "defs",
		Elems: elements,
	}
}

//Desc provides an accessible, long-text description of any SVG container element or graphics element.
func Desc(desc string) svg.Element {
	return &element{
		Name: "desc",
		In:   desc,
	}
}

//Discard allows authors to specify the time at which particular elements are to be discarded,
// thereby reducing the resources required by an SVG user agent.
func Discard(attrs ...svg.Attribute) svg.Element {
	return attrsOnly("discard", attrs...)
}

func Ellipse(attrsOrElements ...svg.Component) svg.Element {
	return Element("ellipse", "", attrsOrElements...)
}

func FeBlend(attrsOrElements ...svg.Component) svg.Element {
	return Element("feBlend", "", attrsOrElements...)
}

func FeColorMatrix(attrsOrElements ...svg.Component) svg.Element {
	return Element("feColorMatrix", "", attrsOrElements...)
}

func FeComponentTransfer(attrsOrElements ...svg.Component) svg.Element {
	return Element("feComponentTransfer", "", attrsOrElements...)
}

func FeComposite(attrsOrElements ...svg.Component) svg.Element {
	return Element("feComposite", "", attrsOrElements...)
}

func FeConvolveMatrix(attrsOrElements ...svg.Component) svg.Element {
	return Element("feConvolveMatrix", "", attrsOrElements...)
}

func FeDiffuseLighting(attrsOrElements ...svg.Component) svg.Element {
	return Element("feDiffuseLighting", "", attrsOrElements...)
}

func FeDisplacementMap(attrsOrElements ...svg.Component) svg.Element {
	return Element("feDisplacementMap", "", attrsOrElements...)
}

func FeDropShadow(attrsOrElements ...svg.Component) svg.Element {
	return Element("feDropShadow", "", attrsOrElements...)
}

func FeDistantLight(attrsOrElements ...svg.Component) svg.Element {
	return Element("feDistantLight", "", attrsOrElements...)
}

func FeFlood(attrsOrElements ...svg.Component) svg.Element {
	return Element("feFlood", "", attrsOrElements...)
}

func FeFuncA(attrsOrElements ...svg.Component) svg.Element {
	return Element("feFuncA", "", attrsOrElements...)
}

func FeFuncB(attrsOrElements ...svg.Component) svg.Element {
	return Element("feFuncB", "", attrsOrElements...)
}

func FeFuncG(attrsOrElements ...svg.Component) svg.Element {
	return Element("feFuncG", "", attrsOrElements...)
}

func FeFuncR(attrsOrElements ...svg.Component) svg.Element {
	return Element("feFuncR", "", attrsOrElements...)
}

func FeGaussianBlur(attrsOrElements ...svg.Component) svg.Element {
	return Element("feGaussianBlur", "", attrsOrElements...)
}

func FeImage(attrsOrElements ...svg.Component) svg.Element {
	return Element("feImage", "", attrsOrElements...)
}

func FeMerge(attrsOrElements ...svg.Component) svg.Element {
	return Element("feMerge", "", attrsOrElements...)
}

func FeMergeNode(attrsOrElements ...svg.Component) svg.Element {
	return Element("feMergeNode", "", attrsOrElements...)
}

func FeOffset(attrsOrElements ...svg.Component) svg.Element {
	return Element("feOffset", "", attrsOrElements...)
}

func FePointLight(attrsOrElements ...svg.Component) svg.Element {
	return Element("fePointLight", "", attrsOrElements...)
}

func FeSpecularLighting(attrsOrElements ...svg.Component) svg.Element {
	return Element("feSpecularLighting", "", attrsOrElements...)
}

func FeSpotLight(attrsOrElements ...svg.Component) svg.Element {
	return Element("feSpotLight", "", attrsOrElements...)
}

func FeTile(attrsOrElements ...svg.Component) svg.Element {
	return Element("feTile", "", attrsOrElements...)
}

func FeTurbulence(attrsOrElements ...svg.Component) svg.Element {
	return Element("feTurbulence", "", attrsOrElements...)
}

func Filter(attrsOrElements ...svg.Component) svg.Element {
	return Element("filter", "", attrsOrElements...)
}

func ForeignObject(attrsOrElements ...svg.Component) svg.Element {
	return Element("foreignObject", "", attrsOrElements...)
}

func G(attrsOrElements ...svg.Component) svg.Element {
	return Element("g", "", attrsOrElements...)
}

func Image(attrsOrElements ...svg.Component) svg.Element {
	return Element("image", "", attrsOrElements...)
}

func Line(attrsOrElements ...svg.Component) svg.Element {
	return Element("line", "", attrsOrElements...)
}

func LinearGradient(attrsOrElements ...svg.Component) svg.Element {
	return Element("linearGradient", "", attrsOrElements...)
}

func Marker(attrsOrElements ...svg.Component) svg.Element {
	return Element("marker", "", attrsOrElements...)
}

func Mask(attrsOrElements ...svg.Component) svg.Element {
	return Element("mask", "", attrsOrElements...)
}

func Metadata(data string) svg.Element {
	return &element{
		Name: "metadata",
		In:   data,
	}
}
func MPath(attrsOrElements ...svg.Component) svg.Element {
	return Element("mpath", "", attrsOrElements...)
}

func Path(attrsOrElements ...svg.Component) svg.Element {
	return Element("path", "", attrsOrElements...)
}

func Pattern(attrsOrElements ...svg.Component) svg.Element {
	return Element("pattern", "", attrsOrElements...)
}

func Polyline(attrsOrElements ...svg.Component) svg.Element {
	return Element("polyline", "", attrsOrElements...)
}

func Polygon(attrsOrElements ...svg.Component) svg.Element {
	return Element("polygon", "", attrsOrElements...)
}

func RadialGradient(attrsOrElements ...svg.Component) svg.Element {
	return Element("radialGradient", "", attrsOrElements...)
}

func Rect(attrsOrElements ...svg.Component) svg.Element {
	return Element("rect", "", attrsOrElements...)
}

func Set(attrs ...svg.Attribute) svg.Element {
	return attrsOnly("set", attrs...)
}

func Stop(attrsOrElements ...svg.Component) svg.Element {
	return Element("stop", "", attrsOrElements...)
}

func Switch(attrsOrElements ...svg.Component) svg.Element {
	return Element("switch", "", attrsOrElements...)
}

func Symbol(attrsOrElements ...svg.Component) svg.Element {
	return Element("symbol", "", attrsOrElements...)
}

func Text(attrsOrElements ...interface{}) svg.Element {
	attrs := make([]svg.Attribute, 0)
	elements := make([]svg.Element, 0)
	text := ""
	for _, x := range attrsOrElements {
		switch x.(type) {
		case svg.Element:
			elements = append(elements, x.(svg.Element))
		case svg.Attribute:
			tmp := x.(svg.Attribute)
			attrs = append(attrs, tmp)
		case string:
			text = x.(string)
		default:
			panic("Text parameters must be either an Element, Attribute or a string")
		}
	}

	return &element{
		Name:  "text",
		Attrs: attrs,
		Elems: elements,
		In:    text,
	}
}

func TextPath(attrsOrElements ...interface{}) svg.Element {
	attrs := make([]svg.Attribute, 0)
	elements := make([]svg.Element, 0)
	text := ""
	for _, x := range attrsOrElements {
		switch x.(type) {
		case svg.Element:
			elements = append(elements, x.(svg.Element))
		case svg.Attribute:
			tmp := x.(svg.Attribute)
			attrs = append(attrs, tmp)
		case string:
			text = x.(string)
		default:
			panic("TextPath parameters must be either an Element, Attribute or a string.")
		}
	}

	return &element{
		Name:  "textPath",
		Attrs: attrs,
		Elems: elements,
		In:    text,
	}
}

func Title(value string, attrs ...svg.Attribute) svg.Element {
	return &element{
		Name:  "title",
		Attrs: attrs,
		In:    value,
	}
}

func Tspan(value string, attrs ...svg.Attribute) svg.Element {
	return &element{
		Name:  "tspan",
		Attrs: attrs,
		In:    value,
	}
}

func Use(attrs ...svg.Attribute) svg.Element {
	return attrsOnly("use", attrs...)
}

func View(attrs ...svg.Attribute) svg.Element {
	return attrsOnly("view", attrs...)
}

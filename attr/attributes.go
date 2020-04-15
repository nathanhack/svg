package attr

import (
	"github.com/nathanhack/svg"
	"github.com/nathanhack/svg/internal"
)

type attribute struct {
	svg.Core
	N string
	V string
}

func (a *attribute) Name() string {
	return a.N
}

func (a *attribute) Value() string {
	return a.V
}

////////////////////////////////////////////

func Class(class string) svg.Attribute {
	return &attribute{
		N: "class",
		V: class,
	}
}

func ClipPath(clipPath string) svg.Attribute {
	return &attribute{
		N: "clip-path",
		V: clipPath,
	}
}

func ClipRule(rule string) svg.Attribute {
	return &attribute{
		N: "clip-rule",
		V: rule,
	}
}

func Color(color string) svg.Attribute {
	return &attribute{
		N: "clip-rule",
		V: color,
	}
}

//ColorInterpolation specifies the color space for gradient interpolations, color animations, and alpha compositing.
//Values: auto | sRGB | linearRGB
//Default: sRBG
func ColorInterpolation(colorInterpolation string) svg.Attribute {
	return &attribute{
		N: "color-interpolation",
		V: colorInterpolation,
	}
}

//Cursor specifies the mouse cursor displayed when the mouse pointer is over an element
//Values: auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize| text | wait | help
func Cursor(cursor string) svg.Attribute {
	return &attribute{
		N: "cursor",
		V: cursor,
	}
}

//Cx is the x-axis coordinate of the center of the circle.
func Cx(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "cx",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//Cy is the y-axis coordinate of the center of the circle.
func Cy(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "cy",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func Height(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "height",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//Display lets you control the rendering of graphical or container elements.
//Some of the Value: block | inline | run-in | flow | flow-root | table | flex |
// grid | ruby | contents | none | inline-block | inline-list-item | inline-table | inline-flex | inline-grid
func Display(display string) svg.Attribute {
	return &attribute{
		N: "display",
		V: display,
	}
}

//Fill has two different meanings. For shapes and text it's a presentation attribute that
//defines the color (or any SVG paint servers like gradients or patterns) used to paint the
//element; for animation it defines the final state of the animation.
func Fill(fill string) svg.Attribute {
	return &attribute{
		N: "fill",
		V: fill,
	}
}

//FillOpacity is a presentation attribute defining the opacity of the paint server (color, gradient, pattern, etc) applied to a shape.
//Value: [0-1] | percentage
func FillOpacity(numberOrPercent svg.NumberOrPercent) svg.Attribute {
	return &attribute{
		N: "fill-opacity",
		V: internal.Stringify(numberOrPercent, ""),
	}
}

//FillRule is a presentation attribute defining the algorithm to use to determine the inside part of a shape.
func FillRule(rule string) svg.Attribute {
	return &attribute{
		N: "fill-rule",
		V: rule,
	}
}

//Filter attribute specifies the filter effects defined by the <filter> element that shall be applied to its element.
func Filter(filter string) svg.Attribute {
	return &attribute{
		N: "filter",
		V: filter,
	}

}

//ID specifies the element's ID
func ID(id string) svg.Attribute {
	return &attribute{
		N: "id",
		V: id,
	}
}

//Mask is a presentation attribute mainly used to bind a given <mask> element with the element the attribute belongs to.
func Mask(mask string) svg.Attribute {
	return &attribute{
		N: "mask",
		V: mask,
	}
}

//Opacity specifies the transparency of an object or of a group of objects, that is, the degree to which the background behind the element is overlaid.
//Values: [0-1]
func Opacity(number svg.Number) svg.Attribute {
	return &attribute{
		N: "opacity",
		V: internal.Stringify(number, ""),
	}
}

//PathLength is the total length for the circle's circumference, in user units.
//Values: bounding-box | visiblePainted | visibleFill | visibleStroke | visible | painted | fill | stroke | all | none
func PathLength(number svg.Number) svg.Attribute {
	return &attribute{
		N: "pathLength",
		V: internal.Stringify(number, ""),
	}
}

//PointerEvents is a presentation attribute that allows defining whether or when an element may be the target of a mouse event.
func PointerEvents(event string) svg.Attribute {
	return &attribute{
		N: "pointer-events",
		V: event,
	}
}

//R is the radius of the circle. A value lower or equal to zero disables rendering of the circle.
func R(length svg.Length) svg.Attribute {
	return &attribute{
		N: "r",
		V: internal.Stringify(length, ""),
	}
}

//ShapeRendering provides hints to the renderer about what tradeoffs to make when rendering shapes like paths, circles, or rectangles.
//Values: auto | optimizeSpeed | crispEdges | geometricPrecision
func ShapeRendering(rendering string) svg.Attribute {
	return &attribute{
		N: "shape-rendering",
		V: rendering,
	}
}

//Stroke is a presentation attribute defining the color (or any SVG paint servers like gradients or patterns) used to paint the outline of the shape.
func Stroke(color string) svg.Attribute {
	return &attribute{
		N: "stroke",
		V: color,
	}
}

//StrokeDasharray is a presentation attribute defining the pattern of dashes and gaps used to paint the outline of the shape.
func StrokeDasharray(dasharray string) svg.Attribute {
	return &attribute{
		N: "stroke-dasharray",
		V: dasharray,
	}
}

//StrokeDashoffset is a presentation attribute defining an offset on the rendering of the associated dash array.
func StrokeDashoffset(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "stroke-dashoffset",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//StrokeLinecap is a presentation attribute defining the shape to be used at the end of open subpaths when they are stroked.
//Values: butt | round | square
func StrokeLinecap(style string) svg.Attribute {
	return &attribute{
		N: "stroke-linecap",
		V: style,
	}
}

//StrokeLinejoin is a presentation attribute defining the shape to be used at the corners of paths when they are stroked.
func StrokeLinejoin(style string) svg.Attribute {
	return &attribute{
		N: "stroke-linejoin",
		V: style,
	}
}

//StrokeMiterlimit is a presentation attribute defining a limit on the ratio of the miter length to the stroke-width used to draw a miter join.
func StrokeMiterlimit(number svg.Number) svg.Attribute {
	return &attribute{
		N: "stroke-miterlimit",
		V: internal.Stringify(number, ""),
	}
}

//StrokeOpacity is a presentation attribute defining the opacity of the paint server (color, gradient, pattern, etc) applied to the stroke of a shape.
//Values: [0-1] | percentage
func StrokeOpacity(numberOrPercent svg.NumberOrPercent) svg.Attribute {
	return &attribute{
		N: "stroke-opacity",
		V: internal.Stringify(numberOrPercent, ""),
	}
}

//StrokeWidth is a presentation attribute defining the width of the stroke to be applied to the shape.
func StrokeWidth(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "stroke-width",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func Style(style string) svg.Attribute {
	return &attribute{
		N: "style",
		V: style,
	}
}

func TabIndex(integer svg.Number) svg.Attribute {
	return &attribute{
		N: "tabindex",
		V: internal.Stringify(integer, ""),
	}
}

//VectorEffect pecifies the vector effect to use when drawing an object. Vector effects are applied
//before any of the other compositing operations, i.e. filters, masks and clips.
//Values: none | non-scaling-stroke | non-scaling-size | non-rotation | fixed-position
func VectorEffect(effect string) svg.Attribute {
	return &attribute{
		N: "vector-effect",
		V: effect,
	}
}

func ViewBox(minX, minY, width, height svg.Number) svg.Attribute {
	return &attribute{
		N: "viewbox",
		V: internal.Stringify(minX, " ") +
			internal.Stringify(minY, " ") +
			internal.Stringify(width, " ") +
			internal.Stringify(height, " "),
	}
}

//Visibility lets you control the visibility of graphical elements. With a value of hidden or collapse the current graphics element is invisible.
//Values: visible | hidden | collapses
func Visibility(state string) svg.Attribute {
	return &attribute{
		N: "visibility",
		V: state,
	}
}

func Width(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "width",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func X(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "x",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func XmlnsDefault() svg.Attribute {
	return &attribute{
		N: "xmlns",
		V: "http://www.w3.org/2000/svg",
	}
}

func Xmlns(url string) svg.Attribute {
	return &attribute{
		N: "xmlns",
		V: url,
	}
}

func Y(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &attribute{
		N: "y",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

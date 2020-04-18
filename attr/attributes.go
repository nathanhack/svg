package attr

import (
	"github.com/gopherjs/vecty"
	"github.com/nathanhack/svg"
	"github.com/nathanhack/svg/attr/transforms"
	"github.com/nathanhack/svg/internal"
	"strings"
)

type nameValueAttribute struct {
	svg.Core
	N string
	V string
}

func (n *nameValueAttribute) Apply(h *vecty.HTML) {
	vecty.Attribute(n.N, n.V).Apply(h)
}

//Attribute is for creating generic attributes not currently defined in the package
func Attribute(tag, value string) svg.Attribute {
	return &nameValueAttribute{
		N: tag,
		V: value,
	}
}

////////////////////////////////////////////

//Accumulate controls whether or not an animation is cumulative.
//Values: none | sum
func Accumulate(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "accumulate",
		V: mode,
	}
}

//Additive controls whether or not an animation is additive.
//Values: replace | sum
func Additive(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "additive",
		V: mode,
	}
}

//AttributeName indicates the name of the CSS property or attribute of the target element that is going to be changed during an animation.
func AttributeName(name string) svg.Attribute {
	return &nameValueAttribute{
		N: "attributeName",
		V: name,
	}
}

//BaseFrequency represents the base frequency parameter for the noise function of the <feTurbulence> filter primitive.
func BaseFrequency(freq svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "baseFrequency",
		V: internal.Stringify(freq, "s"),
	}
}

//Begin defines when an animation should begin or when an element should be discarded.
//Defaults to seconds
func Begin(time svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "begin",
		V: internal.Stringify(time, "s"),
	}
}

//Bias shifts the range of the filter. After applying the kernelMatrix of the <feConvolveMatrix> element to the input image to
//  yield a number and applied the divisor attribute, the bias attribute is added to each component.
func Bias(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "bias",
		V: internal.Stringify(number, ""),
	}
}

//By specifies a relative offset value for an attribute that will be modified during an animation.
func By(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "by",
		V: internal.Stringify(number, ""),
	}
}

//CalcMode specifies the interpolation mode for the animation.
//Values: discrete | linear | paced | spline
func CalcMode(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "calcMode",
		V: mode,
	}
}

func Class(class string) svg.Attribute {
	return &nameValueAttribute{
		N: "class",
		V: class,
	}
}

//ClipPath defines a clipping path, to be used used by the clip-path property.
func ClipPath(clipPath string) svg.Attribute {
	return &nameValueAttribute{
		N: "clip-path",
		V: clipPath,
	}
}

//ClipPathUnits indicates which coordinate system to use for the contents of the <clipPath> element.
//Values: userSpaceOnUse | objectBoundingBox
func ClipPathUnits(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "clipPathUnits",
		V: value,
	}
}

func ClipRule(rule string) svg.Attribute {
	return &nameValueAttribute{
		N: "clip-rule",
		V: rule,
	}
}

func Color(color string) svg.Attribute {
	return &nameValueAttribute{
		N: "clip-rule",
		V: color,
	}
}

//ColorInterpolation specifies the color space for gradient interpolations, color animations, and alpha compositing.
//Values: auto | sRGB | linearRGB
//Default: sRBG
func ColorInterpolation(colorInterpolation string) svg.Attribute {
	return &nameValueAttribute{
		N: "color-interpolation",
		V: colorInterpolation,
	}
}

//Cursor specifies the mouse cursor displayed when the mouse pointer is over an element
//Values: auto | crosshair | default | pointer | move | e-resize | ne-resize | nw-resize | n-resize | se-resize | sw-resize | s-resize | w-resize| text | wait | help
func Cursor(cursor string) svg.Attribute {
	return &nameValueAttribute{
		N: "cursor",
		V: cursor,
	}
}

//Cx is the x-axis coordinate of the center of the circle.
func Cx(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "cx",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//Cy is the y-axis coordinate of the center of the circle.
func Cy(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "cy",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//D defines a path to be drawn.
func D(path string) svg.Attribute {
	return &nameValueAttribute{
		N: "d",
		V: path,
	}
}

//DiffuseConstant represents the kd value in the Phong lighting model. In SVG, this can be any non-negative number.
func DiffuseConstant(kd svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "diffuseConstant",
		V: internal.Stringify(kd, ""),
	}
}

//Display lets you control the rendering of graphical or container elements.
//Some of the Value: block | inline | run-in | flow | flow-root | table | flex |
// grid | ruby | contents | none | inline-block | inline-list-item | inline-table | inline-flex | inline-grid
func Display(display string) svg.Attribute {
	return &nameValueAttribute{
		N: "display",
		V: display,
	}
}

//Divisor specifies the value by which the resulting number of applying the kernelMatrix of a <feConvolveMatrix>
//  element to the input image color value is divided to yield the destination color value.
func Divisor(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "divisor",
		V: internal.Stringify(number, "s"),
	}
}

//Dur indicates the simple duration of an animation.
//Defaults to seconds
func Dur(time svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "dur",
		V: internal.Stringify(time, "s"),
	}
}

//Dx indicates a shift along the x-axis on the position of an element or its content.
func Dx(value svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "dx",
		V: internal.Stringify(value, ""),
	}
}

//Dy indicates a shift along the y-axis on the position of an element or its content.
func Dy(value svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "dy",
		V: internal.Stringify(value, ""),
	}
}

//EdgeMode determines how to extend the input image as necessary with color values so that the matrix
// operations can be applied when the kernel is positioned at or near the edge of the input image.
//Values: duplicate | wrap | none
func EdgeMode(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "edgeMode",
		V: mode,
	}
}

//End defines an end value for the animation that can constrain the active duration.
//Defaults to second, if not time pass in a string
func End(value svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "dur",
		V: internal.Stringify(value, "s"),
	}
}

//Fill has two different meanings. For shapes and text it's a presentation attribute that
//defines the color (or any SVG paint servers like gradients or patterns) used to paint the
//element; for animation it defines the final state of the animation.
func Fill(fill string) svg.Attribute {
	return &nameValueAttribute{
		N: "fill",
		V: fill,
	}
}

//FillOpacity is a presentation attribute defining the opacity of the paint server (color, gradient, pattern, etc) applied to a shape.
//Value: [0-1] | percentage
func FillOpacity(numberOrPercent svg.NumberOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "fill-opacity",
		V: internal.Stringify(numberOrPercent, ""),
	}
}

//FillRule is a presentation attribute defining the algorithm to use to determine the inside part of a shape.
func FillRule(rule string) svg.Attribute {
	return &nameValueAttribute{
		N: "fill-rule",
		V: rule,
	}
}

//Filter attribute specifies the filter effects defined by the <filter> element that shall be applied to its element.
func Filter(filter string) svg.Attribute {
	return &nameValueAttribute{
		N: "filter",
		V: filter,
	}
}

//FilterUnits defines the coordinate system for the attributes x, y, width and height.
//Values: userSpaceOnUse | objectBoundingBox
func FilterUnits(units string) svg.Attribute {
	return &nameValueAttribute{
		N: "filter",
		V: units,
	}
}

//Fr defines the radius of the start circle of the radial gradient.
func Fr(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "fr",
		V: internal.Stringify(length, ""),
	}
}

//From indicates the initial value of the attribute that will be modified during the animation.
func From(value svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "from",
		V: internal.Stringify(value, ""),
	}
}

//Fx defines the x coordinate of the start circle of the radial gradient.
func Fx(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "fx",
		V: internal.Stringify(length, ""),
	}
}

//Fy defines the y coordinate of the start circle of the radial gradient.
func Fy(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "fx",
		V: internal.Stringify(length, ""),
	}
}

//GradientTransform contains the definition of an optional additional transformation from the gradient coordinate system onto the target coordinate system.
func GradientTransform(transforms ...transforms.Transform) svg.Attribute {
	sb := strings.Builder{}
	for _, t := range transforms {
		sb.WriteString(string(t))
	}

	return &nameValueAttribute{
		N: "gradientTransform",
		V: sb.String(),
	}
}

//GradientUnits defines the coordinate system used for attributes specified on the gradient elements.
//Values: userSpaceOnUse | objectBoundingBox
func GradientUnits(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "gradientUnits",
		V: mode,
	}
}

func Height(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "height",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func Href(url string) svg.Attribute {
	return &nameValueAttribute{
		N: "href",
		V: url,
	}
}

//ID specifies the element's ID
func ID(id string) svg.Attribute {
	return &nameValueAttribute{
		N: "id",
		V: id,
	}
}

//In identifies the second input for the given filter primitive. It works exactly like the in attribute.
//Values: SourceGraphic | SourceAlpha | BackgroundImage | BackgroundAlpha | FillPaint | StrokePaint | <filter-primitive-reference>
func In(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "in",
		V: value,
	}
}

//In2 identifies input for the given filter primitive.
//Values: SourceGraphic | SourceAlpha | BackgroundImage | BackgroundAlpha | FillPaint | StrokePaint | <filter-primitive-reference>
func In2(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "in2",
		V: value,
	}
}

//K1 defines one of the values to be used within the the arithmetic operation of the <feComposite> filter primitive.
func K1(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "k1",
		V: internal.Stringify(number, ""),
	}
}

//K2 defines one of the values to be used within the the arithmetic operation of the <feComposite> filter primitive.
func K2(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "k2",
		V: internal.Stringify(number, ""),
	}
}

//K3 defines one of the values to be used within the the arithmetic operation of the <feComposite> filter primitive.
func K3(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "k3",
		V: internal.Stringify(number, ""),
	}
}

//K4 defines one of the values to be used within the the arithmetic operation of the <feComposite> filter primitive.
func K4(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "k4",
		V: internal.Stringify(number, ""),
	}
}

//KernelMatrix defines the list of numbers that make up the kernel matrix for the <feConvolveMatrix> element.
func KernelMatrix(cellNumbers ...svg.Number) svg.Number {
	sb := strings.Builder{}
	for _, t := range cellNumbers {
		sb.WriteString(internal.Stringify(t, " "))
	}

	return &nameValueAttribute{
		N: "kernelMatrix",
		V: sb.String(),
	}
}

//KeySplines defines a set of Bézier curve control points associated with the keyTimes list, defining a
// cubic Bézier function that controls interval pacing.
func KeySplines(curves ...string) svg.Attribute {
	sb := strings.Builder{}
	for i, t := range curves {
		sb.WriteString(t)
		if i < len(curves)-1 {
			sb.WriteString(" ; ")
		}
	}
	return &nameValueAttribute{
		N: "keySplines",
		V: sb.String(),
	}
}

//KeyTimes represents a list of time values used to control the pacing of the animation.
func KeyTimes(times ...svg.Number) svg.Attribute {
	sb := strings.Builder{}
	for i, t := range times {
		sb.WriteString(internal.Stringify(t, ""))
		if i < len(times)-1 {
			sb.WriteString(" ; ")
		}
	}

	return &nameValueAttribute{
		N: "keyTimes",
		V: sb.String(),
	}
}

//LengthAdjust controls how the text is stretched into the length defined by the textLength attribute.
//Values: spacing | spacingAndGlyphs
func LengthAdjust(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "lengthAdjust",
		V: mode,
	}
}

//LimitingConeAngle represents the angle in degrees between the spot light axis (i.e. the axis between the light
//  source and the point to which it is pointing at) and the spot light cone.
func LimitingConeAngle(angle svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "limitingConeAngle",
		V: internal.Stringify(angle, ""),
	}
}

//MarkerHeight represents the height of the viewport into which the <marker> is to be fitted when it is rendered
//  according to the viewBox and preserveAspectRatio attributes.
func MarkerHeight(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "markerHeight",
		V: internal.Stringify(length, ""),
	}
}

//MarkerUnits defines the coordinate system for the markerWidth and markerUnits attributes and the contents of the <marker>.
func MarkerUnits(units string) svg.Attribute {
	return &nameValueAttribute{
		N: "markerUnits",
		V: units,
	}
}

//MarkerWidth represents the width of the viewport into which the <marker> is to be fitted when it is rendered according to
//  the viewBox and preserveAspectRatio attributes.
//Default: 3
func MarkerWidth(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "markerWidth",
		V: internal.Stringify(length, ""),
	}
}

//Mask is a presentation attribute mainly used to bind a given <mask> element with the element the attribute belongs to.
func Mask(mask string) svg.Attribute {
	return &nameValueAttribute{
		N: "mask",
		V: mask,
	}
}

//MaskUnit indicates which coordinate system to use for the contents of the <mask> element.
//Values: userSpaceOnUse | objectBoundingBox
func MaskContentUnits(units string) svg.Attribute {
	return &nameValueAttribute{
		N: "maskContentUnits",
		V: units,
	}
}

//MaskUnit indicates which coordinate system to use for the geometry properties of the <mask> element.
//Values: userSpaceOnUse | objectBoundingBox
func MaskUnit(units string) svg.Attribute {
	return &nameValueAttribute{
		N: "maskUnits",
		V: units,
	}
}

//Max specifies the maximum value of the active animation duration.
func Max(time svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "max",
		V: internal.Stringify(time, "s"),
	}
}

//Method indicates the method by which text should be rendered along the path of a <textPath> element.
//Values: align | stretch
func Method(method string) svg.Attribute {
	return &nameValueAttribute{
		N: "method",
		V: method,
	}
}

//Min specifies the minimum value of the active animation duration.
func Min(time svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "min",
		V: internal.Stringify(time, "s"),
	}
}

//Mode defines the blending mode on the <feBlend> filter primitive.
//Values: normal | multiply | screen | overlay | darken | lighten | color-dodge | color-burn | hard-light | soft-light |
//        difference | exclusion | hue | saturation | color | luminosity
func Mode(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "mode",
		V: value,
	}
}

//NumOctaves defines the number of octaves for the noise function of the <feTurbulence> primitive.
func NumOctaves(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "numOctaves",
		V: internal.Stringify(number, ""),
	}
}

//Offset defines where the gradient stop is placed along the gradient vector.
func Offset(numberOrPercent svg.NumberOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "offset",
		V: internal.Stringify(numberOrPercent, ""),
	}
}

//Opacity specifies the transparency of an object or of a group of objects, that is, the degree to which the background behind the element is overlaid.
//Values: [0-1]
func Opacity(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "opacity",
		V: internal.Stringify(number, ""),
	}
}

//Operator has two meanings based on the context itʼs used in. Either it defines the compositing or morphing operation to be performed.
//Values: over | in | out | atop | xor | lighter | arithmetic
func Operator(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "mode",
		V: value,
	}
}

//Order indicates the size of the matrix to be used by a <feConvolveMatrix> element.
func Order(cellNumbers ...svg.Number) svg.Number {
	sb := strings.Builder{}
	for _, t := range cellNumbers {
		sb.WriteString(internal.Stringify(t, " "))
	}

	return &nameValueAttribute{
		N: "order",
		V: sb.String(),
	}
}

//Orient indicates how a marker is rotated when it is placed at its position on the shape.
//Values: auto | auto-start-reverse | <angle> | <number>
func Orient(orient svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "orient",
		V: internal.Stringify(orient, ""),
	}
}

//PathLength is the total length for the circle's circumference, in user units.
//Values: bounding-box | visiblePainted | visibleFill | visibleStroke | visible | painted | fill | stroke | all | none
func PathLength(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "pathLength",
		V: internal.Stringify(number, ""),
	}
}

//PatternContentUnits indicates which coordinate system to use for the contents of the <pattern> element.
//Values: userSpaceOnUse | objectBoundingBox
func PatternContentUnits(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "patternContentUnits",
		V: mode,
	}
}

//PatternTransform defines a list of transform definitions that are applied to a pattern tile.
func PatternTransform(transforms ...transforms.Transform) svg.Attribute {
	sb := strings.Builder{}
	for _, t := range transforms {
		sb.WriteString(string(t))
	}

	return &nameValueAttribute{
		N: "patternTransform",
		V: sb.String(),
	}
}

//PatternUnits indicates which coordinate system to use for the geometry properties of the <pattern> element.
//Values: userSpaceOnUse | objectBoundingBox
func PatternUnits(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "patternUnits",
		V: mode,
	}
}

//Points defines a list of points.
func Points(points string) svg.Attribute {
	return &nameValueAttribute{
		N: "points",
		V: points,
	}
}

//PointsAtX represents the x location in the coordinate system established by attribute primitiveUnits on the <filter> element of the point at which the light source is pointing.
func PointsAtX(coordinate svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "pointsAtX",
		V: internal.Stringify(coordinate, ""),
	}
}

//PointsAtY represents the y location in the coordinate system established by attribute primitiveUnits on the <filter> element of the point at which the light source is pointing.
func PointsAtY(coordinate svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "pointsAtY",
		V: internal.Stringify(coordinate, ""),
	}
}

//PointsAtZ represents the z location in the coordinate system established by attribute primitiveUnits on the <filter> element of the point at which the light source is pointing, assuming that, in the initial local coordinate system, the positive z-axis comes out towards the person viewing the content and assuming that one unit along the z-axis equals one unit in x and y.
func PointsAtZ(coordinate svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "pointsAtZ",
		V: internal.Stringify(coordinate, ""),
	}
}

//PointerEvents is a presentation attribute that allows defining whether or when an element may be the target of a mouse event.
func PointerEvents(event string) svg.Attribute {
	return &nameValueAttribute{
		N: "pointer-events",
		V: event,
	}
}

//PreserveAspectRatio indicates how an element with a viewBox providing a given aspect ratio must fit into a viewport with a different aspect ratio.
func PreserveAspectRatio(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "preserveAspectRatio",
		V: value,
	}
}

//PreserveAlpha indicates how a <feConvolveMatrix> element handled alpha transparency.
func PreserveAlpha(preserve bool) svg.Attribute {
	s := "false"
	if preserve {
		s = "true"
	}
	return &nameValueAttribute{
		N: "preserveAlpha",
		V: s,
	}
}

//PrimitiveUnits specifies the coordinate system for the various length values within the filter primitives and for the attributes that define the filter primitive subregion.
//Values: userSpaceOnUse | objectBoundingBox
func PrimitiveUnits(units string) svg.Attribute {
	return &nameValueAttribute{
		N: "primitiveUnits",
		V: units,
	}
}

//R is the radius of the circle. A value lower or equal to zero disables rendering of the circle.
func R(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "r",
		V: internal.Stringify(length, ""),
	}
}

//RefX defines the x coordinate of an element’s reference point.
// Values: <number> | left | center | right
func RefX(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "refX",
		V: internal.Stringify(length, ""),
	}
}

//RefY defines the y coordinate of an element’s reference point.
// Values: <number> | top | center | bottom
func RefY(length svg.Length) svg.Attribute {
	return &nameValueAttribute{
		N: "refY",
		V: internal.Stringify(length, ""),
	}
}

//RepeatCount indicates the number of times an animation will take place.
//Values: a number or "indefinite"
func RepeatCount(count svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "repeatCount",
		V: internal.Stringify(count, ""),
	}
}

//RepeatCount indicates the number of times an animation will take place.
//Values: a time or "indefinite"
//Defaults to seconds
func RepeatDur(duration svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "repeatDur",
		V: internal.Stringify(duration, "s"),
	}
}

//Restart specifies whether or not an animation can restart.
//Values: always | whenNotActive | never
func Restart(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "restart",
		V: value,
	}
}

//Rotate specifies how the animated element rotates as it travels along a path specified in an <animateMotion> element.
//Values: auto | auto-reverse | <number>
func Rotate(value svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "rotate",
		V: internal.Stringify(value, ""),
	}
}

//Scale defines the displacement scale factor to be used on a <feDisplacementMap> filter primitive.
func Scale(scale svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "scale",
		V: internal.Stringify(scale, ""),
	}
}

//Seed represents the starting number for the pseudo random number generator of the <feTurbulence> filter primitive.
func Seed(seed svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "seed",
		V: internal.Stringify(seed, ""),
	}
}

//ShapeRendering provides hints to the renderer about what tradeoffs to make when rendering shapes like paths, circles, or rectangles.
//Values: auto | optimizeSpeed | crispEdges | geometricPrecision
func ShapeRendering(rendering string) svg.Attribute {
	return &nameValueAttribute{
		N: "shape-rendering",
		V: rendering,
	}
}

//Side determines the side of a path the text is placed on (relative to the path direction).
//Values left | right
func Side(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "side",
		V: mode,
	}
}

//Spacing indicates how the user agent should determine the spacing between typographic characters that are to be rendered along a path.
//Values: auto | exact
func Spacing(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "spacing",
		V: mode,
	}
}

//SpecularExponent controls the focus for the light source. The bigger the value the brighter the light.
func SpecularExponent(exponent svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "specularExponent",
		V: internal.Stringify(exponent, ""),
	}
}

//SpreadMethod determines how a shape is filled beyond the defined edges of a gradient.
//Values: pad | reflect | repeat
func SpreadMethod(mode string) svg.Attribute {
	return &nameValueAttribute{
		N: "spreadMethod",
		V: mode,
	}
}

//StartOffset defines an offset from the start of the path for the initial current text position along the path after
// converting the path to the <textPath> element's coordinate system.
func StartOffset(offset svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "startOffset",
		V: internal.Stringify(offset, ""),
	}
}

//StitchTiles defines how the Perlin Noise tiles behave at the border.
//Values: noStitch | stitch
func StitchTiles(tiles string) svg.Attribute {
	return &nameValueAttribute{
		N: "stitchTiles",
		V: tiles,
	}
}

//StopColor defines the color of the gradient stop.
func StopColor(color string) svg.Attribute {
	return &nameValueAttribute{
		N: "stop-color",
		V: color,
	}
}

//StopOpacity defines the opacity of a given color gradient stop.
//Value: [0-1]
func StopOpacity(opacity svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "stop-opacity",
		V: internal.Stringify(opacity, ""),
	}
}

//Stroke is a presentation attribute defining the color (or any SVG paint servers like gradients or patterns) used to paint the outline of the shape.
func Stroke(color string) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke",
		V: color,
	}
}

//StrokeDasharray is a presentation attribute defining the pattern of dashes and gaps used to paint the outline of the shape.
func StrokeDasharray(dasharray string) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke-dasharray",
		V: dasharray,
	}
}

//StrokeDashoffset is a presentation attribute defining an offset on the rendering of the associated dash array.
func StrokeDashoffset(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke-dashoffset",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//StrokeLinecap is a presentation attribute defining the shape to be used at the end of open subpaths when they are stroked.
//Values: butt | round | square
func StrokeLinecap(style string) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke-linecap",
		V: style,
	}
}

//StrokeLinejoin is a presentation attribute defining the shape to be used at the corners of paths when they are stroked.
func StrokeLinejoin(style string) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke-linejoin",
		V: style,
	}
}

//StrokeMiterlimit is a presentation attribute defining a limit on the ratio of the miter length to the stroke-width used to draw a miter join.
func StrokeMiterlimit(number svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke-miterlimit",
		V: internal.Stringify(number, ""),
	}
}

//StrokeOpacity is a presentation attribute defining the opacity of the paint server (color, gradient, pattern, etc) applied to the stroke of a shape.
//Values: [0-1] | percentage
func StrokeOpacity(numberOrPercent svg.NumberOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke-opacity",
		V: internal.Stringify(numberOrPercent, ""),
	}
}

//StrokeWidth is a presentation attribute defining the width of the stroke to be applied to the shape.
func StrokeWidth(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "stroke-width",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func Style(style string) svg.Attribute {
	return &nameValueAttribute{
		N: "style",
		V: style,
	}
}

//SurfaceScale represents the height of the surface for a light filter primitive.
func SurfaceScale(height svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "surfaceScale",
		V: internal.Stringify(height, ""),
	}
}

//SystemLanguage represents a list of supported language tags.
func SystemLanguage(lang string) svg.Attribute {
	return &nameValueAttribute{
		N: "systemLanguage",
		V: lang,
	}
}

func TabIndex(integer svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "tabindex",
		V: internal.Stringify(integer, ""),
	}
}

//TargetX determines the positioning in horizontal direction of the convolution matrix relative to a given target pixel in the input image.
func TargetX(integer svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "targetX",
		V: internal.Stringify(integer, ""),
	}
}

//TargetY determines the positioning in vertical direction of the convolution matrix relative to a given target pixel in the input image.
func TargetY(integer svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "targetY",
		V: internal.Stringify(integer, ""),
	}
}

//TO indicates the final value of the attribute that will be modified during the animation.
func To(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "to",
		V: value,
	}
}

//TextLength available on SVG <text> and <tspan> elements, lets you specify the width of the space into which the text will draw.
func TextLength(value svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "textLength",
		V: internal.Stringify(value, ""),
	}
}

func Transform(transforms ...transforms.Transform) svg.Attribute {
	sb := strings.Builder{}
	for _, t := range transforms {
		sb.WriteString(string(t))
	}

	return &nameValueAttribute{
		N: "transform",
		V: sb.String(),
	}
}

//Type is a generic attribute and it has different meaning based on the context in which it's used.
func Type(value string) svg.Attribute {
	return &nameValueAttribute{
		N: "type",
		V: value,
	}
}

//Values has different meanings, depending upon the context where itʼs used, either it defines
// a sequence of values used over the course of an animation, or itʼs a list of numbers for a color matrix, which is interpreted differently depending on the type of color change to be performed.
func Values(numbers ...svg.Number) svg.Attribute {
	sb := strings.Builder{}
	for i, t := range numbers {
		sb.WriteString(internal.Stringify(t, ""))
		if i < len(numbers)-1 {
			sb.WriteString(" ; ")
		}
	}

	return &nameValueAttribute{
		N: "values",
		V: sb.String(),
	}
}

//VectorEffect pecifies the vector effect to use when drawing an object. Vector effects are applied
//before any of the other compositing operations, i.e. filters, masks and clips.
//Values: none | non-scaling-stroke | non-scaling-size | non-rotation | fixed-position
func VectorEffect(effect string) svg.Attribute {
	return &nameValueAttribute{
		N: "vector-effect",
		V: effect,
	}
}

func ViewBox(minX, minY, width, height svg.Number) svg.Attribute {
	return &nameValueAttribute{
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
	return &nameValueAttribute{
		N: "visibility",
		V: state,
	}
}

func Width(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "width",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func X(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "x",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//X2 is used to specify the first x-coordinate for drawing an SVG element that requires more than one coordinate.
func X1(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "x1",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

func X2(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "x2",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//XChannelSelector indicates which color channel from in2 to use to displace the pixels in in along the x-axis.
//Values: R | G | B | A
func XChannelSelector(channel string) svg.Attribute {
	return &nameValueAttribute{
		N: "xChannelSelector",
		V: channel,
	}
}

func Y(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "y",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//Y1 is used to specify the first y-coordinate for drawing an SVG element that requires more than one coordinate.
func Y1(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "y1",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//Y2 is used to specify the first y-coordinate for drawing an SVG element that requires more than one coordinate.
func Y2(lengthOrPercent svg.LengthOrPercent) svg.Attribute {
	return &nameValueAttribute{
		N: "y2",
		V: internal.Stringify(lengthOrPercent, ""),
	}
}

//YChannelSelector indicates which color channel from in2 to use to displace the pixels in in along the y-axis.
//Values: R | G | B | A
func YChannelSelector(channel string) svg.Attribute {
	return &nameValueAttribute{
		N: "yChannelSelector",
		V: channel,
	}
}

//Z defines the location along the z-axis for a light source in the coordinate system established by the primitiveUnits attribute on the <filter> element, assuming that, in the initial coordinate system, the positive z-axis comes out towards the person viewing the content and assuming that one unit along the z-axis equals one unit in x and y.
func Z(coordinate svg.Number) svg.Attribute {
	return &nameValueAttribute{
		N: "z",
		V: internal.Stringify(coordinate, ""),
	}
}

package transforms

import (
	"github.com/nathanhack/svg"
	"github.com/nathanhack/svg/internal"
	"strings"
)

type Transform string

func Matrix(a, b, c, d, e, f svg.Number) Transform {
	sb := strings.Builder{}
	sb.WriteString("matrix(")
	sb.WriteString(internal.Stringify(a, " "))
	sb.WriteString(internal.Stringify(b, " "))
	sb.WriteString(internal.Stringify(c, " "))
	sb.WriteString(internal.Stringify(d, " "))
	sb.WriteString(internal.Stringify(e, " "))
	sb.WriteString(internal.Stringify(f, " "))

	sb.WriteString(")")

	return Transform(sb.String())
}

func Rotate(angle svg.Number, xy ...svg.Number) Transform {
	if len(xy) != 0 && len(xy) != 2 {
		panic("Rotate expects zero or two coordinates, x and y.")
	}
	sb := strings.Builder{}
	sb.WriteString("rotate(")
	sb.WriteString(internal.Stringify(angle, " "))
	for _, c := range xy {
		sb.WriteString(internal.Stringify(c, " "))
	}
	sb.WriteString(")")

	return Transform(sb.String())
}

func Scale(xy ...svg.Number) Transform {
	if len(xy) == 0 {
		panic("Scale expects at least one scale.")
	}
	if len(xy) > 2 {
		panic("Scale expects at most two scaling factors.")
	}
	sb := strings.Builder{}
	sb.WriteString("scale(")
	for _, c := range xy {
		sb.WriteString(internal.Stringify(c, " "))
	}
	sb.WriteString(")")

	return Transform(sb.String())
}

func SkewX(skew svg.Number) Transform {
	return Transform("skewX(" + internal.Stringify(skew, " ") + ")")
}

func SkewY(skew svg.Number) Transform {
	return Transform("skewY(" + internal.Stringify(skew, " ") + ")")
}

func Translate(xy ...svg.Number) Transform {
	if len(xy) == 0 {
		panic("Translate expects at least one translation.")
	}
	if len(xy) > 2 {
		panic("Translate expects at most two translation value.")
	}
	sb := strings.Builder{}
	sb.WriteString("translate(")
	for _, c := range xy {
		sb.WriteString(internal.Stringify(c, " "))
	}
	sb.WriteString(")")

	return Transform(sb.String())
}

package color

import (
	"math"
)

//http://www.rapidtables.com/convert/color/index.htm

func RGB2HSL(r, g, b float64) (h, s, l float64) {
	r = r / 255
	g = g / 255
	b = b / 255

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	d := max - min
	if d == 0 {
		h = 0
		s = 0
	} else if r == max {
		h = math.Mod((g-b)/d, 6)
	} else if g == max {
		h = (b-r)/d + 2
	} else {
		h = (r-g)/d + 4
	}

	h = h * 60
	if h < 0 {
		h = h + 360
	}

	l = (max + min) / 2
	if d != 0 {
		s = d / (1 - math.Abs(2*l-1))
	}
	return
}

func HSL2RGB(h, s, l float64) (r, g, b float64) {
	h = h / 60
	c := (1 - math.Abs(2*l-1)) * s
	x := (1 - math.Abs(math.Mod(h, 2)-1)) * c

	if h < 1 {
		r = c
		g = x
	} else if h < 2 {
		r = x
		g = c
	} else if h < 3 {
		g = c
		b = x
	} else if h < 4 {
		g = x
		b = c
	} else if h < 5 {
		r = x
		b = c
	} else {
		r = c
		b = x
	}
	m := l - c/2
	r = r + m
	g = g + m
	b = b + m
	return math.Floor(r * 255), math.Floor(g * 255), math.Floor(b * 255)
}

func RGB2HSV(r, g, b float64) (h, s, v float64) {
	r = r / 255
	g = g / 255
	b = b / 255

	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	d := max - min

	if d == 0 {
		h = 0
		s = 0
	} else if r == max {
		h = math.Mod((g-b)/d, 6)
	} else if g == max {
		h = (b-r)/d + 2
	} else {
		h = (r-g)/d + 4
	}

	h = h * 60
	if h < 0 {
		h = h + 360
	}

	v = max
	if d != 0 {
		s = d / v
	}
	return
}

func HSV2RGB(h, s, v float64) (r, g, b float64) {
	h = h / 60
	c := v * s
	x := (1 - math.Abs(math.Mod(h, 2)-1)) * c

	if h < 1 {
		r = c
		g = x
	} else if h < 2 {
		r = x
		g = c
	} else if h < 3 {
		g = c
		b = x
	} else if h < 4 {
		g = x
		b = c
	} else if h < 5 {
		r = x
		b = c
	} else {
		r = c
		b = x
	}

	m := v - c
	r = r + m
	g = g + m
	b = b + m
	return math.Floor(r * 255), math.Floor(g * 255), math.Floor(b * 255)
}

func CMYK2RGB(c, m, y, k float64) (r, g, b float64) {
	var c_, m_, y_, k_ float64 = 1 - c, 1 - m, 1 - y, 1 - k
	r = 255 * c_ * k_
	g = 255 * m_ * k_
	b = 255 * y_ * k_
	return
}

func RGB2CMYK(r, g, b float64) (c, m, y, k float64) {
	var r_, g_, b_ float64 = r / 255, g / 255, b / 255
	k = 1 - math.Max(math.Max(r_, g_), b_)
	c = (1 - r_ - k) / (1 - k)
	m = (1 - g_ - k) / (1 - k)
	y = (1 - b_ - k) / (1 - k)
	return
}

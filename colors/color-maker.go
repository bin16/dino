package colors

import (
	"image/color"
	"log"
)

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347, 349, 353, 359}

type DinoPalette struct {
	primary    color.RGBA
	secondary  color.RGBA
	border     color.RGBA
	background color.RGBA
}

func MakePal(keys []byte) pixPalette {
	t, _, _ := calcSeeds(keys)

	c0 := selectHSV(t)
	h, s, v := c0.HSVasInt()
	dur := t % 180

	log.Println("Keys", keys)
	log.Println("T", t)
	log.Println("DSP", dur)
	log.Println("Color 0", c0)

	if dur >= 150 {
		return mk150p(h, s, v, t)
	} else if dur >= 120 {
		return mk120p(h, s, v, t)
	} else if dur >= 90 {
		return mk90p(h, s, v, t)
	} else if dur >= 60 {
		return mk60p(h, s, v, t)
	} else if dur >= 30 {
		return mk30p(h, s, v, t)
	} else if dur < 30 {
		return mk0p(h, s, v, t)
	}

	return pixPalette{
		p:  c0.toRGBA(),
		s:  c0.toRGBA(),
		b:  c0.toRGBA(),
		bg: c0.toRGBA(),
	}
}

func mk150p(h, s, v, t int) pixPalette {
	order := 1
	dur := t % 180

	ph := h
	ps := s
	pv := v

	sh := ph
	ss := ps
	sv := pv

	bh := ph
	bs := 100 - ps
	bv := ifOrNot(pv > 30, pv/2, 5)

	gh := ph + order*dur
	gs := ps
	gv := pv

	cP := NewHSV(ph, ps, pv)
	cS := NewHSV(sh, ss, sv)
	cB := NewHSV(bh, bs, bv)
	cBg := NewHSV(gh, gs, gv)

	return pixPalette{
		p:  cP.toRGBA(),
		s:  cS.toRGBA(),
		b:  cB.toRGBA(),
		bg: cBg.toRGBA(),
	}
}

func mk120p(h, s, v, t int) pixPalette {
	order := 1
	dur := t % 180

	ph := h
	ps := s
	pv := v

	sh := ph + order*60
	ss := ps
	sv := pv

	gh := ph - order*dur
	gs := ps
	gv := pv

	bh := 0
	bs := maxInt(5, ps/5)
	bv := maxInt(pv-50, pv/2)

	cP := NewHSV(ph, ps, pv)
	cS := NewHSV(sh, ss, sv)
	cB := NewHSV(bh, bs, bv)
	cBg := NewHSV(gh, gs, gv)

	return pixPalette{
		p:  cP.toRGBA(),
		s:  cS.toRGBA(),
		b:  cB.toRGBA(),
		bg: cBg.toRGBA(),
	}
}

func mk90p(h, s, v, t int) pixPalette {
	order := 1
	dur := t % 180

	h1 := h + order*dur
	h2 := h - order*dur
	h3 := h + 180

	gh := h
	ph := h1
	sh := h2
	bh := h3

	bs, ps, gs, ss := s, s, s, s
	pv, gv, sv := v, v, v
	bv := maxInt(v/2, v-50)

	cP := NewHSV(ph, ps, pv)
	cS := NewHSV(sh, ss, sv)
	cB := NewHSV(bh, bs, bv)
	cBg := NewHSV(gh, gs, gv)

	log.Println("90p", cP, cS, cB, cBg)

	return pixPalette{
		p:  cP.toRGBA(),
		s:  cS.toRGBA(),
		b:  cB.toRGBA(),
		bg: cBg.toRGBA(),
	}
}

func mk60p(h, s, v, t int) pixPalette {
	order := 1
	dur := t % 180

	ph := h
	ps := s
	pv := v

	sh := ph + order*dur
	ss := ps
	sv := pv

	gh := ph - order*120
	gs := ps
	gv := pv

	bh := 0
	bs := maxInt(5, ps/5)
	bv := maxInt(pv-50, pv/2)

	cP := NewHSV(ph, ps, pv)
	cS := NewHSV(sh, ss, sv)
	cB := NewHSV(bh, bs, bv)
	cBg := NewHSV(gh, gs, gv)

	return pixPalette{
		p:  cP.toRGBA(),
		s:  cS.toRGBA(),
		b:  cB.toRGBA(),
		bg: cBg.toRGBA(),
	}
}

func mk30p(h, s, v, t int) pixPalette {
	order := 1
	dur := t % 180

	h1 := h + order*dur
	h2 := h - order*dur
	h3 := h + 180

	gh := h
	ph := h1
	sh := h2
	bh := h3

	bs, ps, gs, ss := s, s, s, s
	pv, gv, sv := v, v, v
	bv := maxInt(v/2, v-50)

	cP := NewHSV(ph, ps, pv)
	cS := NewHSV(sh, ss, sv)
	cB := NewHSV(bh, bs, bv)
	cBg := NewHSV(gh, gs, gv)

	return pixPalette{
		p:  cP.toRGBA(),
		s:  cS.toRGBA(),
		b:  cB.toRGBA(),
		bg: cBg.toRGBA(),
	}
}

func mk0p(h, s, v, t int) pixPalette {
	order := 1
	dur := t % 180

	ph := h
	ps := s
	pv := v

	sh := ph + order*dur
	sv := pv + order*t%15
	ss := ps

	bh := ph - t%59
	bs := ps - t%15
	bv := ifOrNot(pv > 50, pv-50, pv/2)

	dgb := 79 + order*t%101

	gh := bh + dgb
	gs := ifOrNot(dgb < 60, bs, 100-bs)
	gv := 100 - bv

	cP := NewHSV(ph, ps, pv)
	cS := NewHSV(sh, ss, sv)
	cB := NewHSV(bh, bs, bv)
	cBg := NewHSV(gh, gs, gv)

	return pixPalette{
		p:  cP.toRGBA(),
		s:  cS.toRGBA(),
		b:  cB.toRGBA(),
		bg: cBg.toRGBA(),
	}
}

func selectHSV(t int) HSV {
	h := t%359 + 1
	s := t%61 + 13
	v := t%71 + 29

	return NewHSV(h, s, v)
}

func ifOrNot(c bool, a, b int) int {
	if c {
		return a
	}

	return b
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func calcSeeds(keys []byte) (t1, t2, t3 int) {
	var seed1, seed2, seed3 int = 0, 0, 0
	for i := 0; i < len(keys); i++ {
		item := int(keys[i])
		seed1 += item * primes[i]
		seed2 += (i + 1) * item
		seed3 += item
	}

	return seed1, seed2, seed3
}

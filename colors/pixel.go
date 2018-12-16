package colors

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"path"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
)

type PixelPaint struct {
	tpl         string // tpl name
	tag         string // tag name
	tplDir      string // tpl dir
	tplImg      image.Image
	colorsIndex []int
	pal         pixPalette
	target      string
}

func (pix PixelPaint) tplPath() string {
	if len(pix.tag) > 0 {
		return path.Join(pix.tplDir, pix.tpl+"-"+pix.tag)
	}

	return path.Join(pix.tplDir, pix.tpl)
}

func (pix *PixelPaint) SetTpl(tpl string) {
	pix.tpl = tpl
}

func (pix *PixelPaint) SetTag(tag string) {
	pix.tag = tag
}

func (pix *PixelPaint) LoadTplImg() error {
	tplPath := pix.tplPath()

	tplFile, err := os.Open(tplPath)
	if err != nil {
		return err
	}
	defer tplFile.Close()

	tplImg, err := png.Decode(tplFile)
	if err != nil {
		return err
	}
	pix.tplImg = tplImg

	return nil
}

func (pix *PixelPaint) findIndex() (p, s, b, bg int) {
	pal := pix.tplImg.ColorModel().(color.Palette)
	iP := pal.Index(red)
	iS := pal.Index(green)
	iB := pal.Index(black)
	iBG := pal.Index(blue)

	pix.colorsIndex = []int{iP, iS, iB, iBG}
	return iP, iS, iB, iBG
}

func (pix *PixelPaint) SetPalette(pal pixPalette) {
	pix.pal = pal
}

func (pix *PixelPaint) draw() {
	pal := pix.tplImg.ColorModel().(color.Palette)
	p, s, b, bg := pix.findIndex()

	pal[p] = pix.pal.P()
	pal[s] = pix.pal.S()
	pal[b] = pix.pal.B()
	pal[bg] = pix.pal.BG()
}

func (pix *PixelPaint) Export() error {
	pix.draw()

	imgFile, err := os.Create(pix.target)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	err1 := png.Encode(imgFile, pix.tplImg)
	if err1 != nil {
		return err1
	}

	return nil
}

func (pix *PixelPaint) ExportAs(targetPath string) error {
	pix.target = targetPath
	return pix.Export()
}

func NewPixel(tplPath, target string) PixelPaint {
	pix := PixelPaint{
		tpl:    path.Base(tplPath),
		tplDir: path.Dir(tplPath),
		tag:    "",
		target: target,
	}

	return pix
}

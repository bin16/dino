package dino

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	dinoTplPal = []color.Color{
		color.RGBA{255, 0, 0, 255}, // R
		color.RGBA{0, 255, 0, 255}, // G
		color.RGBA{0, 0, 255, 255}, // B
	}
	dinoDefault = [3]color.Color{
		color.RGBA{80, 108, 100, 255},
		color.RGBA{208, 211, 213, 255},
		color.RGBA{0, 0, 0, 0},
	}
)

type Dino struct {
	tplPath    string
	targetPath string
	scale      int
	workDir    string
	// images
	tpl image.Image
	img image.Image
	// colors
	colors      [3]color.Color
	colorsIndex [3]int
	whiteIndex  int
	blackIndex  int
}

func NewDino(tplPath string) (Dino, error) {
	tplFile, err := os.Open(tplPath)
	if err != nil {
		return Dino{}, err
	}
	defer tplFile.Close()

	tplImg, err := png.Decode(tplFile)
	if err != nil {
		return Dino{}, err
	}

	pal := tplImg.ColorModel().(color.Palette)
	colorsIndex := [3]int{
		pal.Index(dinoTplPal[0]),
		pal.Index(dinoTplPal[1]),
		pal.Index(dinoTplPal[2]),
	}

	return Dino{
		tplPath:     tplPath,
		scale:       1,
		tpl:         tplImg,
		img:         tplImg,
		colors:      dinoDefault,
		colorsIndex: colorsIndex,
		whiteIndex:  pal.Index(color.RGBA{255, 255, 255, 255}),
		blackIndex:  pal.Index(color.RGBA{0, 0, 0, 255}),
	}, nil
}

func (dino *Dino) SetWorkDir(dir string) {
	dino.workDir = dir
}

func (dino *Dino) UseRandomColors() {
	rand.Seed(time.Now().UnixNano() + 255)

	colorGroup := dinoPalette[rand.Intn(len(dinoPalette))]
	groupSize := len(colorGroup)
	dino.colors = [3]color.Color{
		colorGroup[rand.Intn(groupSize)],
		colorGroup[rand.Intn(groupSize)],
		colorGroup[rand.Intn(groupSize)],
	}
}

func (dino Dino) Scale(scale int) {
	dino.scale = scale
}

func (dino Dino) Draw() {
	pal := dino.img.ColorModel().(color.Palette)

	for ix := 0; ix < len(dino.colorsIndex); ix++ {
		pal[dino.colorsIndex[ix]] = dino.colors[ix]
	}
}

func (dino Dino) Export() error {
	targetFile, err := os.Create(dino.targetPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	err1 := png.Encode(targetFile, dino.img)
	if err1 != nil {
		return err1
	}

	return nil
}

func (dino Dino) AutoExport() (string, error) {
	filename := dino.ColorsToString() + ".png"
	targetPath := filepath.Join(dino.workDir, filename)

	return targetPath, dino.ExportAs(targetPath)
}

func (dino Dino) ExportAs(targetPath string) error {
	dino.targetPath = targetPath
	return dino.Export()
}

func (dino *Dino) SetColors(colors [3]color.Color) {
	dino.colors = colors
}

func (dino *Dino) SetPrimaryColor(color color.Color) {
	dino.colors[0] = color
}

func (dino *Dino) SetSecondaryColor(color color.Color) {
	dino.colors[1] = color
}

func (dino *Dino) SetBackgroundColor(color color.Color) {
	dino.colors[2] = color
}

func (dino Dino) Colors() [3]color.Color {
	return dino.colors
}

func (dino Dino) ColorsToString() string {
	result := "dino"
	for i := 0; i < len(dino.colors); i++ {
		result += "-" + colorToHex(dino.colors[i])
	}

	return result
}

func colorToHex(co color.Color) string {
	r, g, b, _ := co.RGBA()

	return uint32ToHex(r>>8) + uint32ToHex(g>>8) + uint32ToHex(b>>8)
}

func uint32ToHex(num uint32) string {
	c := strconv.FormatUint(uint64(num), 16)
	if len(c) < 2 {
		return "0" + c
	}

	return c
}

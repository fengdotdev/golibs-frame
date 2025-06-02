package goframe

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func (g *Goframe8) DrawImg() (image.Image, error) {
	var zero image.Image

	if !g.initialized {
		return zero, fmt.Errorf("frame not initialized")
	}

	img := image.NewRGBA(image.Rect(0, 0, int(g.x), int(g.y)))

	g.ForEach(func(x uint, y uint, dot RGBADot8) error {

		dotcolor := color.RGBA{dot.R, dot.G, dot.B, dot.A}

		img.Set(int(x), int(y), dotcolor)
		return nil
	})

	return img, nil

}

func (g *Goframe8) DrawImgToData() ([]byte, error) {
	img, err := g.DrawImg()
	if err != nil {
		return nil, fmt.Errorf("error drawing image: %w", err)
	}

	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil, fmt.Errorf("error encoding image to PNG: %w", err)
	}

	return buf.Bytes(), nil
}

func (g *Goframe8) DrawImgToPNG(filename string) error {
	img, err := g.DrawImg()
	if err != nil {
		return fmt.Errorf("error drawing image: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("error encoding image to PNG: %w", err)
	}

	return nil
}

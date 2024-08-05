package main

import (
	"image"
	"image/color"

	"github.com/makeworld-the-better-one/dither/v2"
)

func ditherImage(img image.Image) []uint8 {
	palette := []color.Color{
		color.Black,
		color.White,
	}

	// Create ditherer
	d := dither.NewDitherer(palette)
	d.Matrix = dither.FloydSteinberg

	img = d.Dither(img)

	buffer := make([]uint8, img.Bounds().Max.Y*img.Bounds().Max.X/8)

	bbb := 0
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < (img.Bounds().Max.Y / 8); y++ {
			var buf uint8
			for i := uint8(0); i < 8; i++ {
				r, g, b, _ := img.At(img.Bounds().Max.X-x-1, (y*8)+int(i)).RGBA()
				if r == 0 && g == 0 && b == 0 {
					buf |= 1 << (7 - i)
				}
			}
			buffer[bbb] = buf
			bbb++
		}
	}

	return buffer
}

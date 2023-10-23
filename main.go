package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func initImage(width, height int, bgColor color.RGBA) *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	rect := image.Rectangle{upLeft, lowRight}
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	return img
}

func main() {
	width := 1700
	height := 1200
	bgColor := color.RGBA{255, 255, 255, 0xff}
	img := initImage(width, height, bgColor)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// dx := (float64(j-(width/2)) / 100000000000) - 0.7465
			// dy := (float64(i-(height/2)) / 100000000000) - 0.094749
			dx := float64(j-(width/2))/500 - 0.75
			dy := float64(i-(height/2)) / 500
			a := dx
			b := dy

			for t := 0; t < 1000; t++ {
				distance := (a * a) - (b * b) + dx
				b = (2 * (a * b)) + dy
				a = distance

				if distance > 10 {
					particleColor := color.RGBA{
						uint8(t / 2),
						uint8(t * 3),
						uint8(t * 3),
						0xff,
					}

					img.Set(j, i, particleColor)

					break
				}
			}
		}
	}

	f, _ := os.Create("mandelbrot.png")
	png.Encode(f, img)
}

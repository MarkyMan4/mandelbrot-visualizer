package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
)

type cliArgs struct {
	zoom    float64
	outfile string
	x       float64
	y       float64
}

func initImage(width, height int, bgColor color.RGBA) *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	rect := image.Rectangle{upLeft, lowRight}
	img := image.NewRGBA(rect)
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.ZP, draw.Src)

	return img
}

func parseArgs() (*cliArgs, error) {
	// set defaults
	args := &cliArgs{
		zoom:    400,
		outfile: "mandelbrot.png",
		x:       0.746499,
		y:       0.094748999999,
	}

	var err error = nil

	if len(os.Args) <= 1 {
		return args, nil
	}

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-z" || os.Args[i] == "--zoom" {
			if i == len(os.Args)-1 {
				err = fmt.Errorf("provide an numeric value for zoom")
				break
			}

			zoom, parseErr := strconv.ParseFloat(os.Args[i+1], 64)

			if parseErr != nil {
				return args, parseErr
			}

			args.zoom = zoom
		} else if os.Args[i] == "-o" || os.Args[i] == "--outfile" {
			if i == len(os.Args)-1 {
				err = fmt.Errorf("provide an string value for outfile")
				break
			}

			outfile := os.Args[i+1]
			args.outfile = outfile
		} else if os.Args[i] == "-x" {
			if i == len(os.Args)-1 {
				err = fmt.Errorf("provide an numeric value for x")
				break
			}

			x, parseErr := strconv.ParseFloat(os.Args[i+1], 64)

			if parseErr != nil {
				return args, parseErr
			}

			args.x = x
		} else if os.Args[i] == "-y" {
			if i == len(os.Args)-1 {
				err = fmt.Errorf("provide an numeric value for y")
				break
			}

			y, parseErr := strconv.ParseFloat(os.Args[i+1], 64)

			if parseErr != nil {
				return args, parseErr
			}

			args.y = y
		}
	}

	return args, err
}

func main() {
	args, err := parseArgs()

	if err != nil {
		panic(err)
	}

	// try 10000000000000000 for zoon
	zoom := args.zoom
	outfile := args.outfile
	x := args.x
	y := args.y

	width := 1920
	height := 1080
	bgColor := color.RGBA{200, 200, 200, 0xff}
	img := initImage(width, height, bgColor)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {

			dx := (float64(j-(width/2)) / zoom) - x
			dy := (float64(i-(height/2)) / zoom) - y
			a := dx
			b := dy

			for t := 0; t < 1200; t++ {
				distance := (a * a) - (b * b) + dx
				b = (2 * (a * b)) + dy
				a = distance

				if distance > 10 {
					particleColor := color.RGBA{
						uint8(t / 7),
						uint8(t * 2),
						uint8(t * 3),
						0xff,
					}

					img.Set(j, i, particleColor)

					break
				}
			}
		}
	}

	f, _ := os.Create(outfile)
	png.Encode(f, img)
}

package coimage

import (
	"log"
	"os"

	"image"
	"image/draw"
	"image/jpeg"
	"path/filepath"
)

type Direction int

type imageData struct {
	image  *image.Image
	width  int
	height int
}

const (
	Top Direction = iota
	Left
	Bottom
	Right
)

func Co(pattern string, destination string, direction Direction) {
	totalWidth, totalHeight, maxWidth, maxHeight := 0, 0, 0, 0

	imageList := []*imageData{}
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	if len(matches) == 0 {
		log.Fatal("No matched files found!")
	}
	for _, match := range matches {
		file, err := os.Open(match)
		if err != nil {
			log.Fatal(err)
		}
		img, _, err := image.Decode(file)
		if err != nil {
			log.Fatal(err)
		}

		bound := img.Bounds()
		width := bound.Max.X
		height := bound.Max.Y
		switch direction {
		case Top:
		case Bottom:
			totalHeight += height
			if width > maxWidth {
				maxWidth = width
			}
		case Left:
		case Right:
			totalWidth += width
			if height > maxHeight {
				maxHeight = height
			}
		}

		imageList = append(imageList, &imageData{
			image:  &img,
			width:  width,
			height: height,
		})
	}

	var rgba *image.RGBA
	switch direction {
	case Top:
	case Bottom:
		rgba = image.NewRGBA(image.Rect(0, 0, maxWidth, totalHeight))
	case Left:
	case Right:
		rgba = image.NewRGBA(image.Rect(0, 0, totalWidth, maxHeight))
	}

	var x, y int
	if direction == Top {
		y = totalHeight
	} else if direction == Left {
		x = totalWidth
	}
	for _, img := range imageList {
		switch direction {
		case Top:
			y -= img.height
			rect := image.Rect(x, y, x+img.width, y+img.height)
			draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
		case Bottom:
			rect := image.Rect(x, y, x+img.width, y+img.height)
			draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
			y += img.height
		case Left:
			x -= img.height
			rect := image.Rect(x, y, x+img.width, y+img.height)
			draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
		case Right:
			rect := image.Rect(x, y, x+img.width, y+img.height)
			draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
			x += img.height
		}
	}

	out, err := os.Create(destination)
	if err != nil {
		log.Fatal(err)
		return
	}
	jpeg.Encode(out, rgba, &jpeg.Options{Quality: 90})
}

func CoTop(pattern string, destination string) {
	Co(pattern, destination, Top)
}

func CoLeft(pattern string, destination string) {
	Co(pattern, destination, Left)
}

func CoBottom(pattern string, destination string) {
	Co(pattern, destination, Bottom)
}

func CoRight(pattern string, destination string) {
	Co(pattern, destination, Right)
}

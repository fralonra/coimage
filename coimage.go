package coimage

import (
	"log"
	"os"
	"strconv"
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

type outputData struct {
	totalWidth int
	totalHeight int
	imageList []*imageData
}

const (
	Top Direction = iota
	Left
	Bottom
	Right
)

func Co(pattern string, destination string, direction Direction) {
	output := &outputData{
		imageList: []*imageData{},
	}
	outputList := []*outputData{output}
	
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
		case Bottom, Top:
			tmpHeight := output.totalHeight + height
			if tmpHeight >= 1 << 16 {
				output = &outputData{
					imageList: []*imageData{},
				}
				outputList = append(outputList, output)
			}
			output.totalHeight += height
			if width > output.totalWidth {
				output.totalWidth = width
			}
		case Right, Left:
			tmpWidth := output.totalWidth + width
			if tmpWidth >= 1 << 16 {
				output = &outputData{
					imageList: []*imageData{},
				}
				outputList = append(outputList, output)
			}
			output.totalWidth += width
			if height > output.totalHeight {
				output.totalHeight = height
			}
		}

		output.imageList = append(output.imageList, &imageData{
			image:  &img,
			width:  width,
			height: height,
		})
	}

	for index, output := range outputList {
		var rgba *image.RGBA
		switch direction {
		case Bottom, Top:
			rgba = image.NewRGBA(image.Rect(0, 0, output.totalWidth, output.totalHeight))
		case Right, Left:
			rgba = image.NewRGBA(image.Rect(0, 0, output.totalWidth, output.totalHeight))
		}
		var x, y int
		if direction == Top {
			y = output.totalHeight
		} else if direction == Left {
			x = output.totalWidth
		}

		for _, img := range output.imageList {
			switch direction {
			case Bottom:
				rect := image.Rect(x, y, img.width, y+img.height)
				draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
				y += img.height
			case Right:
				rect := image.Rect(x, y, x+img.width, img.height)
				draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
				x += img.height
			case Top:
				y -= img.height
				rect := image.Rect(x, y, img.width, y+img.height)
				draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
			case Left:
				x -= img.height
				rect := image.Rect(x, y, x+img.width, img.height)
				draw.Draw(rgba, rect, *img.image, image.Point{0, 0}, draw.Over)
			}
		}

		outfile := destination
		if len(outputList) > 0 {
			outfile += "." + strconv.Itoa(index+1)
		}
		out, err := os.Create(outfile)
		if err != nil {
			log.Fatal(err)
		}
		err = jpeg.Encode(out, rgba, &jpeg.Options{Quality: 90})
		if err != nil {
			log.Fatal(err)
		}
	}
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

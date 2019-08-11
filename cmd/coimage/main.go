package main

import (
	"log"
	"os"
	"sort"

	"github.com/fralonra/coimage"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "coimage"
	app.Usage = "concat images"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "direction, d",
			Value: "bottom",
			Usage: "direction",
		},
		cli.StringFlag{
			Name:  "pattern, p",
			Value: "",
			Usage: "file pattern",
		},
		cli.StringFlag{
			Name:  "out, o",
			Value: "out.jpg",
			Usage: "output file",
		},
	}

	app.Action = func(c *cli.Context) error {
		var direction coimage.Direction
		switch c.String("direction") {
		case "t", "top":
			direction = coimage.Top
		case "l", "left":
			direction = coimage.Left
		case "b", "bottom":
			direction = coimage.Bottom
		case "r", "right":
			direction = coimage.Right
		}

		coimage.Co(c.String("pattern"), c.String("out"), direction)
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

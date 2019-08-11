# coimage

[![GoDoc](https://godoc.org/github.com/fralonra/coimage?status.svg)](https://godoc.org/github.com/fralonra/coimage)
![License](https://img.shields.io/badge/License-MIT-blue.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/fralonra/coimage)](https://goreportcard.com/report/github.com/fralonra/coimage)

- Combines images together.
- Supports appending to any side.
- Splits output image automatically if it's too large.

### Installation

```
$ go get -u github.com/fralonra/coimage
```

### Running in command line

Install the `coimage` binary first. In your `$GOPATH/src/github.com/fralonra/coimage` folder, run the following command:

```
$ go install ./...
```

Then run `coimage -h` to see the help.

```
NAME:
   coimage - concat images

USAGE:
   coimage [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --direction value, -d value  direction (default: "bottom")
   --out value, -o value        output file (default: "out.jpg")
   --pattern value, -p value    file pattern
   --help, -h                   show help
   --version, -v                print the version
```

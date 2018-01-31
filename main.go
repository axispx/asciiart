package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"math/rand"
	"os"

	"github.com/fogleman/gg"
)

const (
	// Set of visible ASCII characters.
	asciiset = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"
)

func main() {
	file := flag.String("file", "", "The path of the file to be used.")
	char := flag.String("char", "", "The character to be used. Random characters will be used for every pixel if not provided.")
	charsize := flag.Int("charsize", 10, "The size of the character")
	out := flag.String("out", "out.png", "The name of the output image(PNG).")
	set := flag.String("set", asciiset, "The set of characters to be used.")
	serial := flag.Bool("serial", false, "Use the characters from set serially rather than randomly.")
	flag.Parse()

	// Check if the file is not provided.
	if *file == "" {
		fmt.Println("No file provided.")
		os.Exit(1)
	}

	// Check if character os more than one characters are provided.
	if len(*char) > 1 {
		fmt.Println("Provide a single character")
		os.Exit(1)
	}

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	dc := gg.NewContext(img.Bounds().Max.X, img.Bounds().Max.Y)
	dc.Clear()
	err = dc.LoadFontFace("/Library/Fonts/RobotoMono-Regular.ttf", float64(*charsize))
	if err != nil {
		log.Fatal(err)
	}

	generateArt(img, dc, *char, *charsize, *serial, *set)

	// Save the image
	err = dc.SavePNG(*out)
	if err != nil {
		log.Fatal(err)
	}
}

func generateArt(img image.Image, dc *gg.Context, char string, charsize int, serial bool, set string) {
	bounds := img.Bounds()
	serialIndex := 0
	// The increments are equal to the size of the character.
	// The X-axis is filled first for horizontal text.
	for y := bounds.Min.Y; y < bounds.Max.Y; y += charsize {
		for x := bounds.Min.X; x < bounds.Max.X; x += charsize {
			r, g, b, a := img.At(x, y).RGBA()
			red := float64(r) / 65535.0
			green := float64(g) / 65535.0
			blue := float64(b) / 65535.0
			aa := float64(a) / 65535.0
			dc.SetRGBA(red, green, blue, aa)

			// If the character is not provided then a random
			// character from the set will be taken and drawn.
			if char == "" {
				// Draw the characters serially if the serial flag is provided.
				if serial {
					dc.DrawString(string(set[serialIndex]), float64(x), float64(y))
					serialIndex++

					// Reset serialIndex if it reaches it final position.
					if serialIndex == len(set) {
						serialIndex = 0
					}
				} else {
					dc.DrawString(string(set[rand.Intn(len(set))]), float64(x), float64(y))
				}
			} else {
				dc.DrawString(char, float64(x), float64(y))
			}
		}

	}
}

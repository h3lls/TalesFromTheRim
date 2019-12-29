package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"strconv"
	"image"
	"image/color"
	"image/png"
	"github.com/disintegration/imaging"
)

var CharMap = map[rune]color.RGBA{
	'*': color.RGBA{238, 44, 44, 255},
	'?': color.RGBA{238, 229, 222, 255},
	
	// banners/bright colors	
	'0': color.RGBA{255, 255, 255, 255},	// white banner, rice, cotton
	'1': color.RGBA{238, 0, 0, 255},	// red banner
	'2': color.RGBA{0, 205, 0, 255},	// green banner, peas
	'3': color.RGBA{205, 205, 0, 255},	// yellow banner, corn
	'4': color.RGBA{28, 134, 238, 255},	// blue banner
	'5': color.RGBA{255, 0, 255, 255},	// magenta banner
	'6': color.RGBA{0, 238, 238, 255},	// cyan banner

	// terrain colors
	'a':color.RGBA{139, 26, 26, 255},	// cherries
	'b':color.RGBA{144, 238, 144, 255},	// plains, grove
	'c':color.RGBA{127, 255, 0, 255},	// apples
	'd':color.RGBA{102, 205, 170, 255},	// jungle
	'e':color.RGBA{42, 175, 110, 255},	// swamp
	'f':color.RGBA{0, 100, 0, 255},	// forest
	'g':color.RGBA{108, 140, 0, 255},	// olives, hops
	'h':color.RGBA{153, 255, 255, 255},	// tundra
	'i':color.RGBA{0, 191, 255, 255},	// river
	'j':color.RGBA{30, 144, 255, 255},	// oasis
	'k':color.RGBA{79, 148, 205, 255},	// ocean
	'l':color.RGBA{238, 233, 191, 255},	// wheat, barley
	'm':color.RGBA{255, 236, 139, 255},	// desert
	'n':color.RGBA{255, 165, 79, 255},	// peaches
	'o':color.RGBA{238, 154, 0, 255},	// oranges, gourds
	'p':color.RGBA{142, 142, 56, 255},	// trench
	'q':color.RGBA{139, 117, 0, 255},	// mountain
	'r':color.RGBA{193, 193, 193, 255},	// road
	's':color.RGBA{81, 81, 81, 255},	// building
	't':color.RGBA{0, 0, 102, 255},	// Dark Blue
	'u':color.RGBA{0, 76, 153, 255},	// Dark Azure Blue
	'v':color.RGBA{153, 0, 76, 255},	// Dark Magenta
	'w':color.RGBA{0, 153, 153, 255},	// Dark Cyan
	'x':color.RGBA{102, 255, 102, 255},	// Lime Green
	'y':color.RGBA{0, 153, 0, 255},	// Dark Lime Green
	'z':color.RGBA{204, 102, 0, 255},	// Dark Orange
	'A':color.RGBA{255, 153, 204, 255},	// Pink
	'B':color.RGBA{255, 51, 153, 255},	// Dark Pink
	'C':color.RGBA{210, 180, 140, 255},	// Tan
	'D':color.RGBA{127, 0, 255, 255},	// Violet
	'E':color.RGBA{76, 0, 153, 255},	// Deep Violet
}



func main() {

	// Create standard map
	createMap("/TalesFromTheRim/data/map.txt", "/var/www/muddingonline/map.png")

	// Create terrain only
	createMap("/TalesFromTheRim/lib/world/wld/map.txt", "/var/www/muddingonline/map-terrain.png")

	// Create political map
	createMap("/TalesFromTheRim/data/map-political.txt", "/var/www/muddingonline/map-political.png")

}

func createMap(inputFile string, outputImage string) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var validResolution = regexp.MustCompile(`(\d+)x(\d+)`)

	width := 0
	height := 0
	for scanner.Scan() {
		txt := scanner.Text()
		fmt.Println(txt)
		if validResolution.MatchString(txt) {
			// Found resolution
			str := strings.Split(txt, "x")
			width, _ = strconv.Atoi(str[0])
			height, _ = strconv.Atoi(str[1])
			fmt.Printf("Map Resolution: %d, %d\r\n", width, height)
			break
		} else {
			log.Fatal("Could not find image resolution")
		}
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	x := 0;
	for scanner.Scan() {
		x += 1;
		txt := scanner.Text()
		for idx, char := range(txt) {
			img.Set(idx, x, CharMap[char]);
		}
	}

	// Encode as PNG.
	f, _ := os.Create(outputImage)
	png.Encode(f, imaging.FlipV(img))
}
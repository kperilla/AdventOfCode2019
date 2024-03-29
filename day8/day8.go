package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type layer [][]int

func buildLayer(width int, height int, data []int) layer {
	if len(data) < width*height {
		panic("UH OH")
	}
	outLayer := make([][]int, height)
	i := 0
	for y := 0; y < height; y++ {
		outLayer[y] = data[i : i+width]
		i += width
	}
	return outLayer
}

func stringToSlice(digStr string) []int {
	var intArr = []int{}
	strArr := strings.Split(strings.TrimSpace(digStr), "")
	for i := 0; i < len(strArr); i++ {
		val, err := strconv.Atoi(strArr[i])
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, val)
	}
	return intArr
}

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(dat)
}

func imageLayers(width int, height int, data []int) []layer {
	i := 0
	layers := []layer{}
	layerLen := width * height
	for i+layerLen <= len(data) {
		layers = append(layers, buildLayer(width, height, data[i:i+layerLen]))
		i += layerLen
	}
	return layers
}

func numDigitsInLayer(l layer, dig int) int {
	count := 0
	for y := 0; y < len(l); y++ {
		for x := 0; x < len(l[y]); x++ {
			if l[y][x] == dig {
				count++
			}
		}
	}
	return count
}

func findFewestZero(layers []layer) layer {
	min := -1
	var minLayer layer
	for _, l := range layers {
		count := numDigitsInLayer(l, 0)
		if min == -1 || count < min {
			min = count
			minLayer = l
		}
	}
	return minLayer
}

func findTopPixel(layers []layer, y int, x int) int {
	top := 2
	for i := 0; i < len(layers); i++ {
		if layers[i][y][x] != 2 {
			top = layers[i][y][x]
			break
		}
	}
	return top
}

func findTopImage(layers []layer) layer {
	topArr := []int{}
	height := len(layers[0])
	width := len(layers[0][0])
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			topArr = append(topArr, findTopPixel(layers, y, x))
		}
	}
	return buildLayer(width, height, topArr)
}

func intArrToStr(arr []int) string {
	b := make([]byte, 0, len(arr))
	for _, n := range arr {
		b = strconv.AppendInt(b, int64(n), 10)
	}
	return string(b)
}

func main() {
	arr := stringToSlice(readFile("day8/input"))
	layers := imageLayers(25, 6, arr)
	minLayer := findFewestZero(layers)
	fmt.Println("Part 1:")
	fmt.Println(numDigitsInLayer(minLayer, 1) * numDigitsInLayer(minLayer, 2))
	fmt.Println()

	fmt.Println("Part 2:")
	topImg := findTopImage(layers)
	height := len(layers[0])
	for y := 0; y < height; y++ {
		fmt.Println(intArrToStr(topImg[y]))
	}

}

package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Point2D struct {
	x int
	y int
}

const numberOfThread int = 20

var (
	r         = regexp.MustCompile(`\((\d*),(\d*)\)`)
	waitGroup sync.WaitGroup
)

func findArea(inputChannel chan string) {
	for pointsStr := range inputChannel {
		var points []Point2D
		for _, p := range r.FindAllStringSubmatch(pointsStr, -1) {
			x, _ := strconv.Atoi(p[1])
			y, _ := strconv.Atoi(p[2])
			points = append(points, Point2D{x, y})
		}
		area := 0.0
		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]

			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		fmt.Println(math.Abs(area) / 2.0)
	}
	waitGroup.Done()
}

func main() {
	absPath, _ := filepath.Abs("./")
	dat, _ := os.ReadFile(filepath.Join(absPath, "polygons.txt"))
	text := string(dat)
	inputChannel := make(chan string, 1000)
	for i := 0; i < numberOfThread; i++ {
		fmt.Println("impit channel >>>>>  ")
		waitGroup.Add(1)
		go findArea(inputChannel)
	}

	start := time.Now()
	for _, line := range strings.Split(text, "\n") {
		inputChannel <- line
	}
	close(inputChannel)
	waitGroup.Wait()
	eleapsed := time.Since(start)
	fmt.Println("completed in >>>>>>>>> ")
	fmt.Println(eleapsed)

}

package year2021

import (
	"io/ioutil"
	"math"
	"regexp"

	"github.com/lanphiergm/adventofcodego/internal/utils"
)

// Trick Shot Part 1 computes the highest Y position the probe reaches
func TrickShotPart1(filename string) interface{} {
	xMin, xMax, yMin, yMax := parseTargetArea(filename)

	// Use a finite sum to determine Vx^2+Vx-2x = 0 then use the quadratic
	// formula, ignoring the negative root
	vXMin := int((-1.0 + math.Sqrt(1.0+8.0*float64(xMin))) / 2.0)
	vXMax := int((-1.0+math.Sqrt(1.0+8.0*float64(xMax)))/2.0) + 2

	// I'm sure there is a math solution to finding this, but I spent an hour
	// thinking about it and the brute force solution was just so fast
	vYMin := 0
	vYMax := 200

	maxHeight := 0

	for vX0 := vXMin; vX0 <= vXMax; vX0++ {
		for vY0 := vYMin; vY0 <= vYMax; vY0++ {
			currMaxHeight := 0
			didHit := false
			isPast := false
			x := 0
			y := 0
			vX := vX0
			vY := vY0
			for !isPast {
				x += vX
				y += vY
				if y > currMaxHeight {
					currMaxHeight = y
				}
				if vX > 0 {
					vX--
				}
				vY--
				if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
					didHit = true
					break
				}
				if x > xMax || y < yMin {
					isPast = true
				}
			}
			if didHit && currMaxHeight > maxHeight {
				maxHeight = currMaxHeight
			}
		}
	}

	return maxHeight
}

// Trick Shot Part 2 computes the number of distinct initial velocities that
// will hit the target area
func TrickShotPart2(filename string) interface{} {
	xMin, xMax, yMin, yMax := parseTargetArea(filename)

	vXMin := 0
	vXMax := 300
	vYMin := -300
	vYMax := 300

	hitCount := 0

	for vX0 := vXMin; vX0 <= vXMax; vX0++ {
		for vY0 := vYMin; vY0 <= vYMax; vY0++ {
			isPast := false
			x := 0
			y := 0
			vX := vX0
			vY := vY0
			for !isPast {
				x += vX
				y += vY
				if vX > 0 {
					vX--
				}
				vY--
				if x >= xMin && x <= xMax && y >= yMin && y <= yMax {
					hitCount++
					break
				}
				if x > xMax || y < yMin {
					isPast = true
				}
			}
		}
	}
	return hitCount
}

func parseTargetArea(filename string) (int, int, int, int) {
	contents, _ := ioutil.ReadFile(filename)
	rs := `target area: x=([^\.]*)\.\.([^,]*), y=([^\.]*)\.\.(.*)`
	r := regexp.MustCompile(rs)
	result := r.FindAllStringSubmatch(string(contents), -1)
	return utils.Atoi(result[0][1]), utils.Atoi(result[0][2]),
		utils.Atoi(result[0][3]), utils.Atoi(result[0][4])
}

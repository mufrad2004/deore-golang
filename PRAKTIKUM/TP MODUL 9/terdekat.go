package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func jarak_1301223029(p1, p2 point) float64 {
	return math.Sqrt(((p1.x - p2.x) * (p1.x - p2.x)) + ((p1.y - p2.y) * (p1.y - p2.y)))
}

func main() {
	var (
		a, b point
		c, d point
	)

	fmt.Scan(&a.x, &a.y, &b.x, &b.y, &c.x, &c.y, &d.x, &d.y)

	if jarak_1301223029(a, b) > jarak_1301223029(c, d) {
		fmt.Print("Titik terdekat adalah titik c", c, "dengan d", d, "dengan jarak ", jarak_1301223029(c, d))
	} else {
		fmt.Print("Titik terdekat adalah titik a", a, "dengan b", b, "dengan jarak ", jarak_1301223029(a, b))

	}

}

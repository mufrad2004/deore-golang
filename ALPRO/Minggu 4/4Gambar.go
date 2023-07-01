package main

import "fmt"

func zoomIn_zoomOut(w, h *int, s string, p *int) {
	if s == "+" {
		*w = *w * *p
		*h = *h * *p
	} else if s == "-" {
		*w = *w / *p
		*h = *h / *p
	}
}

func main() {
	var (
		w, h int
		s    string
		p    int
	)
	fmt.Scan(&w, &h)
	fmt.Scan(&s, &p)
	zoomIn_zoomOut(&w, &h, s, &p)

	fmt.Print(w, h)
}

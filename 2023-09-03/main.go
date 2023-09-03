package main

import "fmt"

/**
 * dva * dva = hetyre
 * find dva, hetyre
 *
 * два * два = четыре
 * решаем математический ребус, где разная буква - разная цифра
 */
func main() {
	var dva = 122
	var hetyre int
	for {
		dva++
		if dva > 999 {
			break
		}
		hetyre = dva * dva
		if hetyre < 99999 {
			continue
		}
		a := dva % 10
		v := ((dva - a) % 100) / 10
		d := ((dva - a - v) % 1000) / 100
		if a == v || a == d || v == d {
			continue
		}
		e := hetyre % 10
		r := ((hetyre - e) % 100) / 10
		y := ((hetyre - e - r) % 1000) / 100
		t := ((hetyre - e - r - y) % 10000) / 1000
		xe := ((hetyre - e - r - y - t) % 100000) / 10000
		if xe != e {
			continue
		}
		h := ((hetyre - e - r - y - t - xe) % 1000000) / 100000
		if e == r || e == y || e == t || e == h || r == y || r == t || r == h || y == t || y == h ||
			d == e || d == r || d == y || d == t || d == h ||
			v == e || v == r || v == y || v == t || v == h ||
			a == e || a == r || a == y || a == t || a == h {
			continue
		}
		fmt.Printf("%d * %d = %d\n", dva, dva, hetyre)
	}
}

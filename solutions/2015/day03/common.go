package day03

var houses = make(map[int]map[int]int, 0)

func addHouse(x, y int, houses map[int]map[int]int) {
	if _, ok := houses[x]; !ok {
		houses[x] = make(map[int]int, 0)
	}
	houses[x][y] = 1
}

func getTotalHouses() int {
	totalHouses := 0
	for x := range houses {
		for y := range houses[x] {
			y++
			totalHouses += 1
		}
	}
	return totalHouses
}
